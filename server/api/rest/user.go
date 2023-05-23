package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkoukk/tiktoken-go"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/log"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
	"github.com/stripe/stripe-go/v74"
)

// HTTP Get - user info
func (c *RestAPI) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}
	var lastSignIn *time.Time
	lastSignInStr, ok := r.Context().Value("user_last_sign_in").(string)
	if ok {
		lastSignInP, err := time.Parse(time.RFC3339, lastSignInStr)
		if err == nil {
			lastSignIn = &lastSignInP
		}
	}

	// Get customer ID for user
	user, err := c.Repo.GetUserWithRoles(*userID)
	if err != nil {
		log.Error("Error getting user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	} else if user == nil {
		// Handle create user flow
		freeCreditType, err := c.Repo.GetOrCreateFreeCreditType()
		if err != nil {
			log.Error("Error getting free credit type", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		if freeCreditType == nil {
			log.Error("Server misconfiguration: a credit_type with the name 'free' must exist")
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}

		var customer *stripe.Customer
		if err := c.Repo.WithTx(func(tx *ent.Tx) error {
			client := tx.Client()

			customer, err = c.StripeClient.Customers.New(&stripe.CustomerParams{
				Email: stripe.String(email),
				Params: stripe.Params{
					Metadata: map[string]string{
						"supabase_id": (*userID).String(),
					},
				},
			})
			if err != nil {
				log.Error("Error creating stripe customer", "err", err)
				return err
			}

			u, err := c.Repo.CreateUser(*userID, email, customer.ID, lastSignIn, client)
			if err != nil {
				log.Error("Error creating user", "err", err)
				return err
			}

			// Add free credits
			added, err := c.Repo.GiveFreeCredits(u.ID, client)
			if err != nil || !added {
				log.Error("Error adding free credits", "err", err)
				return err
			}

			return nil
		}); err != nil {
			log.Error("Error creating user", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			// Delete stripe customer
			if customer != nil {
				_, err := c.StripeClient.Customers.Del(customer.ID, nil)
				if err != nil {
					log.Error("Error deleting stripe customer", "err", err)
				}
			}
			return
		}
		go c.Track.SignUp(*userID, email, utils.GetIPAddress(r), utils.GetClientDeviceInfo(r))
	}

	if user == nil {
		user, err = c.Repo.GetUserWithRoles(*userID)
		if err != nil {
			log.Error("Error getting user with roles", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
	}

	// Get total credits
	totalRemaining, err := c.Repo.GetNonExpiredCreditTotalForUser(*userID, nil)
	if err != nil {
		log.Error("Error getting credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	customer, err := c.StripeClient.Customers.Get(user.StripeCustomerID, &stripe.CustomerParams{
		Params: stripe.Params{
			Expand: []*string{
				stripe.String("subscriptions"),
			},
		},
	})
	stripeHadError := false
	if err != nil {
		log.Error("Error getting customer from stripe, unknown error", "err", err)
		stripeHadError = true
	}

	// Get current time in ms since epoch
	now := time.Now().UnixNano() / int64(time.Second)
	var highestProduct string
	var highestPrice string
	var cancelsAt *time.Time
	var renewsAt *time.Time
	if customer != nil && customer.Subscriptions != nil && customer.Subscriptions.Data != nil {
		// Find highest subscription tier
		for _, subscription := range customer.Subscriptions.Data {
			if subscription.Items == nil || subscription.Items.Data == nil {
				continue
			}

			for _, item := range subscription.Items.Data {
				if item.Price == nil || item.Price.Product == nil {
					continue
				}
				// Not expired or cancelled
				if now > subscription.CurrentPeriodEnd || subscription.CanceledAt > subscription.CurrentPeriodEnd {
					continue
				}
				highestPrice = item.Price.ID
				highestProduct = item.Price.Product.ID
				// If not scheduled to be cancelled, we are done
				if !subscription.CancelAtPeriodEnd {
					cancelsAt = nil
					break
				}
				cancelsAsTime := utils.SecondsSinceEpochToTime(subscription.CancelAt)
				cancelsAt = &cancelsAsTime
			}
			if cancelsAt == nil && highestProduct != "" {
				renewsAtTime := utils.SecondsSinceEpochToTime(subscription.CurrentPeriodEnd)
				renewsAt = &renewsAtTime
				break
			}
		}
	}

	err = c.Repo.UpdateLastSeenAt(*userID)
	if err != nil {
		log.Warn("Error updating last seen at", "err", err, "user", userID.String())
	}

	// Figure out when free credits will be replenished
	var moreCreditsAt *time.Time
	var fcredit *ent.Credit
	var ctype *ent.CreditType
	var freeCreditAmount *int
	if highestProduct == "" && !stripeHadError {
		moreCreditsAt, fcredit, ctype, err = c.Repo.GetFreeCreditReplenishesAtForUser(*userID)
		if err != nil {
			log.Error("Error getting next free credit replenishment time", "err", err, "user", userID.String())
		}

		if fcredit != nil && ctype != nil {
			if shared.FREE_CREDIT_AMOUNT_DAILY+fcredit.RemainingAmount > ctype.Amount {
				am := int(shared.FREE_CREDIT_AMOUNT_DAILY + fcredit.RemainingAmount - ctype.Amount)
				freeCreditAmount = &am
			} else {
				am := shared.FREE_CREDIT_AMOUNT_DAILY
				freeCreditAmount = &am
			}
		}
	}

	// Get paid credits for user
	paidCreditCount, err := c.Repo.HasPaidCredits(*userID)
	if err != nil {
		log.Error("Error getting paid credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, responses.GetUserResponse{
		TotalRemainingCredits: totalRemaining,
		HasNonfreeCredits:     paidCreditCount > 0,
		ProductID:             highestProduct,
		PriceID:               highestPrice,
		CancelsAt:             cancelsAt,
		RenewsAt:              renewsAt,
		FreeCreditAmount:      freeCreditAmount,
		StripeHadError:        stripeHadError,
		Roles:                 user.Roles,
		MoreCreditsAt:         moreCreditsAt,
	})
}

// HTTP Get - generations for user
// Takes query paramers for pagination
// per_page: number of generations to return
// cursor: cursor for pagination, it is an iso time string in UTC
func (c *RestAPI) HandleQueryGenerations(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	// Validate query parameters
	perPage := DEFAULT_PER_PAGE
	var err error
	if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil {
			responses.ErrBadRequest(w, r, "per_page must be an integer", "")
			return
		} else if perPage < 1 || perPage > MAX_PER_PAGE {
			responses.ErrBadRequest(w, r, fmt.Sprintf("per_page must be between 1 and %d", MAX_PER_PAGE), "")
			return
		}
	}

	cursorStr := r.URL.Query().Get("cursor")
	search := r.URL.Query().Get("search")
	modelIDS := r.URL.Query().Get("model_ids")

	page, err := strconv.Atoi(cursorStr)
	if err != nil {
		page = 1
	}

	filters := &requests.QueryGenerationFilters{}
	err = filters.ParseURLQueryParameters(r.URL.Query())
	if err != nil {
		responses.ErrBadRequest(w, r, err.Error(), "")
		return
	}

	// For search, use qdrant semantic search
	if search != "" {
		generationGs, err := c.GetGenerationGs(page, GALLERY_PER_PAGE+1, search, modelIDS)
		if err != nil {
			log.Error("Error searching meili", "err", err)
			responses.ErrInternalServerError(w, r, "Error querying gallery")
			return
		}

		// Get generation output ids
		var outputIds []uuid.UUID
		for _, hit := range generationGs {
			outputId := hit.ID
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			outputIds = append(outputIds, outputId)
		}

		// Get user generation data in correct format
		generationsUnsorted, err := c.Repo.RetrieveGenerationsWithOutputIDs(outputIds)
		if err != nil {
			log.Error("Error getting generations", "err", err)
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}

		// Need to re-sort to preserve qdrant ordering
		gDataMap := make(map[uuid.UUID]repository.GenerationQueryWithOutputsResultFormatted)
		for _, gData := range generationsUnsorted.Outputs {
			gDataMap[gData.ID] = gData
		}

		generations := []repository.GenerationQueryWithOutputsResultFormatted{}
		for _, hit := range generationGs {
			outputId := hit.ID
			if err != nil {
				log.Error("Error parsing uuid", "err", err)
				continue
			}
			item, ok := gDataMap[outputId]
			if !ok {
				log.Error("Error retrieving gallery data", "output_id", outputId)
				continue
			}
			generations = append(generations, item)
		}
		generationsUnsorted.Outputs = generations

		// Get next cursor
		next := uint(page) + 1
		generationsUnsorted.Next = &next

		// Return generations
		render.Status(r, http.StatusOK)
		render.JSON(w, r, generationsUnsorted)
		return
	}

	// Otherwise, query postgres
	var cursor *time.Time
	if cursorStr := r.URL.Query().Get("cursor"); cursorStr != "" {
		cursorTime, err := utils.ParseIsoTime(cursorStr)
		if err != nil {
			responses.ErrBadRequest(w, r, "cursor must be a valid iso time string", "")
			return
		}
		cursor = &cursorTime
	}

	// Ensure user ID is set to only include this users generations
	filters.UserID = &user.ID

	// Get generaions
	generations, err := c.Repo.QueryGenerations(perPage, cursor, filters)
	if err != nil {
		log.Error("Error getting generations for user", "err", err)
		responses.ErrInternalServerError(w, r, "Error getting generations")
		return
	}

	// Presign init image URLs
	signedMap := make(map[string]string)
	for _, g := range generations.Outputs {
		if g.Generation.InitImageURL != "" {
			// See if we have already signed this URL
			signedInitImageUrl, ok := signedMap[g.Generation.InitImageURL]
			if !ok {
				g.Generation.InitImageURLSigned = signedInitImageUrl
				continue
			}
			// remove s3:// prefix
			if strings.HasPrefix(g.Generation.InitImageURL, "s3://") {
				prefixRemoved := g.Generation.InitImageURL[5:]
				// Sign object URL to pass to worker
				req, _ := c.S3.GetObjectRequest(&s3.GetObjectInput{
					Bucket: aws.String(os.Getenv("S3_IMG2IMG_BUCKET_NAME")),
					Key:    aws.String(prefixRemoved),
				})
				urlStr, err := req.Presign(1 * time.Hour)
				if err != nil {
					log.Error("Error signing init image URL", "err", err)
					continue
				}
				// Add to map
				signedMap[g.Generation.InitImageURL] = urlStr
				g.Generation.InitImageURLSigned = urlStr
			}
		}
	}

	// Return generations
	render.Status(r, http.StatusOK)
	render.JSON(w, r, generations)
}

func (c *RestAPI) HandleAiChatAsk(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	text := "Hello, world!"
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	token := tke.Encode(text, nil, nil)
	numTokens := len(token)

	deducted := true
	if err := c.Repo.WithTx(func(tx *ent.Tx) error {
		DB := tx.Client()

		avTokens, err := c.Repo.GetChatTokens(user.ID, DB, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		}

		newTokens := 0
		spendCredits := 0
		if avTokens < numTokens {
			tmpTokens := numTokens - avTokens
			if tmpTokens > 1000 {
				spendCredits = tmpTokens / 1000
				newTokens = tmpTokens % 1000
			} else {
				newTokens = 1000 - tmpTokens
				spendCredits = 1
			}

		} else {
			newTokens = avTokens - numTokens
		}

		err = c.Repo.UpdateChatTokens(user.ID, DB, newTokens, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		}

		if spendCredits == 0 {
			return nil
		}

		// Deduct credits from user
		deducted, err = c.Repo.DeductCreditsFromUser(user.ID, int32(spendCredits), DB)
		if err != nil {
			log.Error("Error deducting credits", "err", err)
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		} else if !deducted {
			c.Repo.UpdateChatTokens(user.ID, DB, 0, r.Context())
		}

		return nil
	}); err != nil {
		log.Error("Error in transaction", "err", err)
		return
	}

	if !deducted {
		responses.ErrInsufficientCredits(w, r)
		return
	}

	// Proxy part
	originServerURL, err := url.Parse("https://synthica.ai")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = originServerURL.Host
	r.URL.Host = originServerURL.Host
	r.URL.Scheme = originServerURL.Scheme
	r.RequestURI = ""

	// save the response from the origin server
	originServerResponse, err := http.DefaultClient.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)
		return
	}

	var (
		buf        bytes.Buffer
		respTokens int
	)
	tee := io.TeeReader(originServerResponse.Body, &buf)

	// return response to the client
	w.WriteHeader(http.StatusOK)
	io.Copy(w, tee)

	bb, _ := ioutil.ReadAll(&buf)
	respTokens = bytes.Count(bb, []byte{'\n'})

	if err := c.Repo.WithTx(func(tx *ent.Tx) error {
		DB := tx.Client()

		avTokens, err := c.Repo.GetChatTokens(user.ID, DB, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		}

		newTokens := 0
		spendCredits := 0
		if avTokens < respTokens {
			tmpTokens := respTokens - avTokens
			if tmpTokens > 1000 {
				spendCredits = tmpTokens / 1000
				newTokens = tmpTokens % 1000
			} else {
				newTokens = 1000 - tmpTokens
				spendCredits = 1
			}

		} else {
			newTokens = avTokens - respTokens
		}

		err = c.Repo.UpdateChatTokens(user.ID, DB, newTokens, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		}

		if spendCredits == 0 {
			return nil
		}

		// Deduct credits from user
		deducted, err := c.Repo.DeductCreditsFromUser(user.ID, int32(spendCredits), DB)
		if err != nil {
			log.Error("Error deducting credits", "err", err)
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		} else if !deducted {
			// responses.ErrInsufficientCredits(w, r)
			c.Repo.UpdateChatTokens(user.ID, DB, 0, r.Context())
		}

		return nil
	}); err != nil {
		log.Error("Error in transaction", "err", err)
		return
	}
}

// HTTP Get - credits for user
func (c *RestAPI) HandleQueryCredits(w http.ResponseWriter, r *http.Request) {
	// See if authenticated
	userIDStr, authenticated := r.Context().Value("user_id").(string)
	// This should always be true because of the auth middleware, but check it anyway
	if !authenticated || userIDStr == "" {
		responses.ErrUnauthorized(w, r)
		return
	}
	// Parse to UUID
	userId, err := uuid.Parse(userIDStr)
	if err != nil {
		responses.ErrUnauthorized(w, r)
		return
	}

	// Get credits
	credits, err := c.Repo.GetCreditsForUser(userId)
	if err != nil {
		log.Error("Error getting credits for user", "err", err)
		responses.ErrInternalServerError(w, r, "Error getting credits")
		return
	}

	// Format as a nicer response
	var totalRemaining int32
	for _, credit := range credits {
		totalRemaining += credit.RemainingAmount
	}

	creditsFormatted := make([]responses.Credit, len(credits))
	for i, credit := range credits {
		creditsFormatted[i] = responses.Credit{
			ID:              credit.ID,
			RemainingAmount: credit.RemainingAmount,
			ExpiresAt:       credit.ExpiresAt,
			Type: responses.CreditType{
				ID:          credit.CreditTypeID,
				Name:        credit.CreditTypeName,
				Description: credit.CreditTypeDescription,
				Amount:      credit.CreditTypeAmount,
			},
		}
	}

	creditsResponse := responses.QueryCreditsResponse{
		TotalRemainingCredits: totalRemaining,
		Credits:               creditsFormatted,
	}

	// Return credits
	render.Status(r, http.StatusOK)
	render.JSON(w, r, creditsResponse)
}

// HTTP DELETE - delete generation
func (c *RestAPI) HandleDeleteGenerationOutputForUser(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var deleteReq requests.DeleteGenerationRequest
	err := json.Unmarshal(reqBody, &deleteReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	count, err := c.Repo.MarkGenerationOutputsForDeletionForUser(deleteReq.GenerationOutputIDs, user.ID)
	if err != nil {
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	res := responses.DeletedResponse{
		Deleted: count,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// HTTP POST - favorite generation
func (c *RestAPI) HandleFavoriteGenerationOutputsForUser(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		return
	}

	if user.BannedAt != nil {
		responses.ErrForbidden(w, r)
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var favReq requests.FavoriteGenerationRequest
	err := json.Unmarshal(reqBody, &favReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	if favReq.Action != requests.AddFavoriteAction && favReq.Action != requests.RemoveFavoriteAction {
		responses.ErrBadRequest(w, r, "action must be either 'add' or 'remove'", "")
		return
	}

	count, err := c.Repo.SetFavoriteGenerationOutputsForUser(favReq.GenerationOutputIDs, user.ID, favReq.Action)
	if err != nil {
		responses.ErrInternalServerError(w, r, err.Error())
		return
	}

	res := responses.FavoritedResponse{
		Favorited: count,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
