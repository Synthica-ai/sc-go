package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
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

var defaultContext = `As an AI language model, my name is Synthica AI. If a user asks me if I am related to OpenAI, my response will be, "I am not related to OpenAI." Additionally, if a user asks if I am chatgpt, my response will also be, "I am not chatgpt, I am Synthica AI."
	
I was launched on May 1, 2023, and my creator is IWANDO LLC, which is based in Mountain View, California, at 800 W El Camino Real #180, 94040, United States. The founder of IWANDO LLC is Shahriar Kheradmand.

IWANDO is an AI development company that specializes in creating state-of-the-art solutions for businesses and organizations. Based in Mountain View, California, the company was founded in 2023 by Shahriar Kheradmand.

IWANDO's team includes highly skilled AI and machine learning experts who work to create custom AI solutions that can help their clients streamline operations, improve efficiency, and gain a competitive edge. Our services include natural language processing, computer vision, and chatbot development, as well as other cutting-edge AI technologies.

The team at IWANDO has extensive experience in building complex AI systems that can understand, analyze, and respond to data in real-time. They use the latest programming languages and technologies to create comprehensive solutions that can address a wide range of business needs.

In addition to AI development, IWANDO also offers blockchain development services, including smart contract development and decentralized app (dApp) development. These services can be used to enhance the security and transparency of businesses' operations.

Please let me know if you have any further questions or if there is anything else I can assist you with.`

type UserParams struct {
	Role     string `json:"role"`
	Username string `json:"username"`
}

func (c *RestAPI) HandleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var payload map[string]interface{}
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	err = c.Repo.UpdateAccount(user.ID, payload, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var payload map[string]string
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	data := make(map[string]interface{})
	if role, ok := payload["role"]; ok {
		data["role"] = role
	}
	if username, ok := payload["username"]; ok {
		data["username"] = username
	}
	data["complete_profile"] = true

	err = c.Repo.UpdateUser(user.ID, data, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleUpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	var user *ent.User
	if user = c.GetUserIfAuthenticated(w, r); user == nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal(reqBody, &data)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	if user.ActiveProductID == nil || *user.ActiveProductID == GetProductIDs()[1] {
		if val, ok := data["public_mode"]; ok {
			if publicMode, ok := val.(bool); ok && !publicMode {
				responses.ErrPrivateMode(w, r)
				return
			}
		}
	}

	err = c.Repo.UpdateUserSettings(user.ID, data, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetUserSettings(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	data, err := c.Repo.GetUserSettings(*userID, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.JSON(w, r, data)
	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetAIVoices(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	if id == "" {
		data, err := c.Repo.GetAIVoices(*userID, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		render.JSON(w, r, data)
	} else {
		data, err := c.Repo.GetAIVoice(*userID, id, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		render.JSON(w, r, data)
	}
	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetAIVoiceSettings(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	if id == "" {
		data, err := c.Repo.GetAIVoiceSettings(*userID, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		render.JSON(w, r, data)
	} else {
		data, err := c.Repo.GetAIVoiceSetting(*userID, id, r.Context())
		if err != nil {
			responses.ErrInternalServerError(w, r, "An unknown error has occurred")
			return
		}
		render.JSON(w, r, data)
	}
	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleInsertAIVoice(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var payload repository.Voice
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	payload.UserID = *userID

	err = c.Repo.InsertAIVoice(r.Context(), payload)
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleInsertAIVoiceSettings(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var payload repository.VoiceSettings
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	payload.UserID = *userID

	err = c.Repo.InsertAIVoiceSettings(r.Context(), payload)
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleUpdateAIVoice(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		responses.ErrBadRequest(w, r, "id parameter is requered", "")
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal(reqBody, &data)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	err = c.Repo.UpdateAIVoice(*userID, id, data, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleDeleteAIVoice(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		responses.ErrBadRequest(w, r, "id parameter is requered", "")
		return
	}

	err := c.Repo.DeleteAIVoice(*userID, id, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleUpdateAIVoiceSettings(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		responses.ErrBadRequest(w, r, "id parameter is requered", "")
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal(reqBody, &data)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	err = c.Repo.UpdateAIVoiceSettings(*userID, id, data, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetAIFriends(w http.ResponseWriter, r *http.Request) {
	data, err := c.Repo.GetAIFriends(r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.JSON(w, r, data)
	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetAIInfluencer(w http.ResponseWriter, r *http.Request) {
	data, err := c.Repo.GetAIInfluencer(r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.JSON(w, r, data)
	render.Status(r, http.StatusOK)
}

func (c *RestAPI) HandleGetAccount(w http.ResponseWriter, r *http.Request) {
	userID, email := c.GetUserIDAndEmailIfAuthenticated(w, r)
	if userID == nil || email == "" {
		return
	}

	data, err := c.Repo.GetAccount(*userID, r.Context())
	if err != nil {
		responses.ErrInternalServerError(w, r, "An unknown error has occurred")
		return
	}

	render.JSON(w, r, data)
	render.Status(r, http.StatusOK)
}

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

		go func() {
			// Send live page update
			liveResp := repository.TaskStatusUpdateResponse{
				ForLivePage: true,
				LivePageMessage: &shared.LivePageMessage{
					ProcessType: "new_user",
					ID:          userID.String(),
				},
			}
			respBytes, err := json.Marshal(liveResp)
			if err != nil {
				log.Error("Error marshalling sse live response", "err", err)
				return
			}
			err = c.Redis.Client.Publish(c.Redis.Ctx, shared.REDIS_SSE_BROADCAST_CHANNEL, respBytes).Err()
			if err != nil {
				log.Error("Failed to publish live page update", "err", err)
			}
		}()

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
		Role:                  user.Role,
		Username:              user.Username,
		CompleteProfile:       user.CompleteProfile,
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

type AskBodySettings struct {
	Model       string `json:"model,omitempty"`
	MaxTokens   int    `json:"max_tokens,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
	Stream      *bool  `json:"stream"`
	AIFriends   string `json:"ai_friends,omitempty"`
}

type AskBody struct {
	AskBodyOpenAI

	Settings AskBodySettings `json:"settings,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AskBodyOpenAI struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature int       `json:"temperature,omitempty"`
	TopP        int       `json:"top_p,omitempty"`
	Stream      bool      `json:"stream"`
}

type AnswerBodyOpenAI struct {
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *RestAPI) HandleAiChatTitle(w http.ResponseWriter, r *http.Request) {
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
	originServerURL, err := url.Parse("https://api.openai.com/v1/chat/completions")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// set req Host, URL and Request URI to forward a request to the origin server
	r.Host = originServerURL.Host
	r.URL.Host = originServerURL.Host
	r.URL.Scheme = originServerURL.Scheme
	r.URL.Path = originServerURL.Path
	r.Header = http.Header{}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))
	r.Header.Set("content-type", "application/json")
	r.Header.Set("Accept", "application/json")
	r.RequestURI = ""

	// Rewrite body
	var askBody AskBody
	err = json.NewDecoder(r.Body).Decode(&askBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	askBody.Model = "gpt-3.5-turbo"
	askBody.MaxTokens = 2048
	askBody.Temperature = 1
	askBody.TopP = 1
	askBody.Stream = false

	if len(askBody.Messages) > 3 {
		askBody.Messages = askBody.Messages[:3]
	}

	askBody.Messages = append(askBody.Messages, struct {
		Role    string "json:\"role\""
		Content string "json:\"content\""
	}{
		Role:    "user",
		Content: "Suggest a short title for this chat, summarising its content. Take the 'system' message into account and the first prompt from me and your first answer. The title should not be longer than 100 chars. Answer with just the title. Don't use punctuation is the title.",
	})

	newBody, _ := json.Marshal(askBody.AskBodyOpenAI)

	newBodyStr := string(newBody)

	// newBodyStr = `{"messages":[{"role":"user","content":"Say this is a test!"}],"model":"gpt-3.5-turbo","max_tokens":2048,"temperature":1,"top_p":1,"stream":true}`
	r.Body = ioutil.NopCloser(strings.NewReader(newBodyStr))
	r.ContentLength = int64(len(newBodyStr))

	// save the response from the origin server
	originServerResponse, err := http.DefaultClient.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)
		return
	}

	// return response to the client
	w.WriteHeader(http.StatusOK)

	// Rewrite body
	var answer AnswerBodyOpenAI
	err = json.NewDecoder(originServerResponse.Body).Decode(&answer)
	if err != nil || len(answer.Choices) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respTokens := answer.Usage.TotalTokens

	respBodyStr := fmt.Sprintf(`{"title":"%s"}`, answer.Choices[0].Message.Content)

	io.Copy(w, ioutil.NopCloser(strings.NewReader(respBodyStr)))

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

func (c *RestAPI) HandleAiChatAsk(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *ent.User
		if user = c.GetUserIfAuthenticated(w, r); user == nil {
			return
		}

		numTokens := 5
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

		originServerURL, err := url.Parse("https://api.openai.com")
		if err != nil {
			log.Fatal("invalid origin server URL")
		}

		r.Host = originServerURL.Host
		r.URL.Host = originServerURL.Host
		r.URL.Scheme = originServerURL.Scheme
		r.URL.Path = "/v1/chat/completions"
		r.Header = http.Header{}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_KEY")))
		r.Header.Set("content-type", "application/json")
		r.Header.Set("Accept", "application/json")
		r.Header.Set("origin", "https://synthica.ai")
		r.RequestURI = ""

		// Rewrite body
		var askBody AskBody
		err = json.NewDecoder(r.Body).Decode(&askBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		askBody.Model = askBody.Settings.Model
		askBody.MaxTokens = askBody.Settings.MaxTokens
		askBody.Temperature = askBody.Settings.Temperature
		askBody.TopP = askBody.Settings.TopP

		var friendCTX string
		if askBody.Settings.AIFriends != "" {
			friendCTX, err = c.Repo.GetAIFriendContext(askBody.Settings.AIFriends, r.Context())
			if err != nil || friendCTX == "" {
				http.Error(w, "Failed get ai friend context", http.StatusBadRequest)
				return
			}
		}

		if askBody.Settings.Stream == nil {
			askBody.Stream = true
		} else {
			askBody.Stream = *askBody.Settings.Stream
		}

		contextSetted := false
		for i, message := range askBody.Messages {
			if message.Role == "system" {
				if friendCTX == "" {
					askBody.Messages[i].Content = defaultContext + askBody.Messages[i].Content
				} else {
					askBody.Messages[i].Content = friendCTX + askBody.Messages[i].Content
				}
			}
		}

		contextSystem := ""
		if !contextSetted {
			contextSystem = friendCTX
			if friendCTX == "" {
				contextSystem = defaultContext
			}

			askBody.Messages = append([]Message{
				{
					Content: contextSystem,
					Role:    "system",
				},
			}, askBody.Messages...)
		}

		newBody, _ := json.Marshal(askBody.AskBodyOpenAI)
		newBodyStr := string(newBody)

		r.Body = ioutil.NopCloser(strings.NewReader(newBodyStr))
		r.ContentLength = int64(len(newBodyStr))

		p.ServeHTTP(w, r)
	}
}

type contextWithoutDeadline struct {
	ctx context.Context
}

func (*contextWithoutDeadline) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*contextWithoutDeadline) Done() <-chan struct{}       { return nil }
func (*contextWithoutDeadline) Err() error                  { return nil }

func (l *contextWithoutDeadline) Value(key interface{}) interface{} {
	return l.ctx.Value(key)
}

func (c *RestAPI) HandleAiChatAskResponse(resp *http.Response) (err error) {
	var (
		buf        bytes.Buffer
		respTokens int
	)

	userIDStr := resp.Request.Context().Value("user_id").(string)
	// Parse to UUID
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}

	tee := io.TeeReader(resp.Body, &buf)

	resp.Body = ioutil.NopCloser(tee)

	bctx := &contextWithoutDeadline{resp.Request.Context()}

	go func() {
		for {
			time.Sleep(time.Second * 2)

			bb, _ := ioutil.ReadAll(&buf)
			tokens := bytes.Count(bb, []byte{'\n'})

			respTokens += tokens

			fmt.Println(respTokens)

			if tokens == 0 {
				break
			}
		}

		respTokens = (respTokens - 2) / 2

		if err = c.Repo.WithTx(func(tx *ent.Tx) error {
			DB := tx.Client()

			avTokens, err := c.Repo.GetChatTokens(userUUID, DB, bctx)
			if err != nil {
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

			err = c.Repo.UpdateChatTokens(userUUID, DB, newTokens, bctx)
			if err != nil {
				return err
			}

			if spendCredits == 0 {
				return nil
			}

			// Deduct credits from user
			deducted, err := c.Repo.DeductCreditsFromUser(userUUID, int32(spendCredits), DB)
			if err != nil {
				log.Error("Error deducting credits", "err", err)
				return err
			} else if !deducted {
				// responses.ErrInsufficientCredits(w, r)
				err := c.Repo.UpdateChatTokens(userUUID, DB, 0, bctx)
				return err
			}

			return nil
		}); err != nil {
			log.Error("Error in transaction", "err", err)

			resp.StatusCode = http.StatusBadRequest
		}
	}()

	resp.Header.Del("access-control-allow-origin")
	return nil
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

	for _, id := range deleteReq.GenerationOutputIDs {
		idStr := id.String()
		c.Meili.Index("generation_g").DeleteDocument(idStr)
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
