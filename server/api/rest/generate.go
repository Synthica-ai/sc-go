package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/stablecog/sc-go/database/ent"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/shared"
	"github.com/stablecog/sc-go/utils"
	"k8s.io/klog/v2"
)

// POST generate endpoint
// Adds generate to queue, if authenticated, returns the ID of the generation
func (c *RestAPI) HandleCreateGeneration(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	userID := c.GetUserIDIfAuthenticated(w, r)
	if userID == nil {
		return
	}
	fmt.Printf("--- GetUserIDIFAuthenticated took: %s\n", time.Now().Sub(start))

	// Parse request body
	start = time.Now()
	reqBody, _ := io.ReadAll(r.Body)
	var generateReq requests.GenerateRequestBody
	err := json.Unmarshal(reqBody, &generateReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}
	fmt.Printf("--- ParseRequestBody took: %s\n", time.Now().Sub(start))

	// Make sure the stream ID is valid
	start = time.Now()
	if !utils.IsSha256Hash(generateReq.StreamID) {
		responses.ErrInvalidStreamID(w, r)
		return
	}
	fmt.Printf("--- Validate stream ID took: %s\n", time.Now().Sub(start))

	// Validate request body
	if generateReq.Height > shared.MAX_GENERATE_HEIGHT {
		responses.ErrBadRequest(w, r, fmt.Sprintf("Height is too large, max is: %d", shared.MAX_GENERATE_HEIGHT))
		return
	}

	if generateReq.Width > shared.MAX_GENERATE_WIDTH {
		responses.ErrBadRequest(w, r, fmt.Sprintf("Width is too large, max is: %d", shared.MAX_GENERATE_WIDTH))
		return
	}

	if generateReq.Width*generateReq.Height*generateReq.InferenceSteps >= shared.MAX_PRO_PIXEL_STEPS {
		klog.Infof(
			"Pick fewer inference steps or smaller dimensions: %d - %d - %d",
			generateReq.Width,
			generateReq.Height,
			generateReq.InferenceSteps,
		)
		responses.ErrBadRequest(w, r, "Pick fewer inference steps or smaller dimensions")
		return
	}

	if generateReq.NumOutputs < 0 {
		generateReq.NumOutputs = shared.DEFAULT_GENERATE_NUM_OUTPUTS
	}
	if generateReq.NumOutputs > shared.MAX_GENERATE_NUM_OUTPUTS {
		klog.Infof("Number of outputs can't be more than %d", shared.MAX_GENERATE_NUM_OUTPUTS)
		responses.ErrBadRequest(w, r, fmt.Sprintf("Number of outputs can't be more than %d", shared.MAX_GENERATE_NUM_OUTPUTS))
		return
	}

	// Validate model and scheduler IDs in request are valid
	start = time.Now()
	if !shared.GetCache().IsValidGenerationModelID(generateReq.ModelId) {
		klog.Infof("invalid_model_id: %s", generateReq.ModelId)
		responses.ErrBadRequest(w, r, "invalid_model_id")
		return
	}

	if !shared.GetCache().IsValidShedulerID(generateReq.SchedulerId) {
		klog.Infof("invalid_scheduler_id: %s", generateReq.SchedulerId)
		responses.ErrBadRequest(w, r, "invalid_scheduler_id")
		return
	}
	fmt.Printf("--- Checking model and scheduler IDs took: %s\n", time.Now().Sub(start))

	// Generate seed if not provided
	if generateReq.Seed < 0 {
		rand.Seed(time.Now().Unix())
		generateReq.Seed = rand.Intn(math.MaxInt32)
	}

	// Parse request headers
	start = time.Now()
	countryCode := utils.GetCountryCode(r)
	deviceInfo := utils.GetClientDeviceInfo(r)
	fmt.Printf("--- Parse request headers took: %s\n", time.Now().Sub(start))

	start = time.Now()

	// ! TODO - parallel generation toggle

	// Get model and scheduler name for cog
	modelName := shared.GetCache().GetGenerationModelNameFromID(generateReq.ModelId)
	schedulerName := shared.GetCache().GetSchedulerNameFromID(generateReq.SchedulerId)
	if modelName == "" || schedulerName == "" {
		klog.Errorf("Error getting model or scheduler name: %s - %s", modelName, schedulerName)
		responses.ErrInternalServerError(w, r, "An unknown error has occured")
		return
	}

	// Format prompts
	start = time.Now()
	generateReq.Prompt = utils.FormatPrompt(generateReq.Prompt)
	generateReq.NegativePrompt = utils.FormatPrompt(generateReq.NegativePrompt)
	fmt.Printf("--- Format prompts took: %s\n", time.Now().Sub(start))

	// For live page update
	var livePageMsg shared.LivePageMessage
	// For keeping track of this request as it gets sent to the worker
	var requestId string
	// Cog request
	var cogReqBody requests.CogQueueRequest

	// Wrap everything in a DB transaction
	// We do this since we want our credit deduction to be atomic with the whole process
	if err := c.Repo.WithTx(func(tx *ent.Tx) error {
		// Bind a client to the transaction
		DB := tx.Client()
		// Deduct credits from user
		deducted, err := c.Repo.DeductCreditsFromUser(*userID, int32(generateReq.NumOutputs), DB)
		if err != nil {
			klog.Errorf("Error deducting credits: %v", err)
			responses.ErrInternalServerError(w, r, "Error deducting credits from user")
			return err
		} else if !deducted {
			responses.ErrInsufficientCredits(w, r)
			return responses.InsufficientCreditsErr
		}
		fmt.Printf("--- Deduct credits took took: %s\n", time.Now().Sub(start))

		// Create generation
		start = time.Now()
		g, err := c.Repo.CreateGeneration(
			*userID,
			string(deviceInfo.DeviceType),
			deviceInfo.DeviceOs,
			deviceInfo.DeviceBrowser,
			countryCode,
			generateReq,
			DB)
		if err != nil {
			klog.Errorf("Error creating generation: %v", err)
			responses.ErrInternalServerError(w, r, "Error creating generation")
			return err
		}
		fmt.Printf("--- Create generation took: %s\n", time.Now().Sub(start))

		// Request Id matches generation ID
		requestId = g.ID.String()

		// For live page update
		livePageMsg = shared.LivePageMessage{
			ProcessType: generateReq.ProcessType,
			ID:          utils.Sha256(requestId),
			CountryCode: countryCode,
			Status:      shared.LivePageQueued,
			Width:       generateReq.Width,
			Height:      generateReq.Height,
			CreatedAt:   g.CreatedAt,
		}

		cogReqBody = requests.CogQueueRequest{
			WebhookEventsFilter: []requests.WebhookEventFilterOption{requests.WebhookEventFilterStart, requests.WebhookEventFilterStart},
			RedisPubsubKey:      shared.COG_REDIS_EVENT_CHANNEL,
			Input: requests.BaseCogRequest{
				ID:                   requestId,
				LivePageData:         livePageMsg,
				Prompt:               generateReq.Prompt,
				NegativePrompt:       generateReq.NegativePrompt,
				Width:                fmt.Sprint(generateReq.Width),
				Height:               fmt.Sprint(generateReq.Height),
				NumInferenceSteps:    fmt.Sprint(generateReq.InferenceSteps),
				GuidanceScale:        fmt.Sprint(generateReq.GuidanceScale),
				Model:                modelName,
				Scheduler:            schedulerName,
				Seed:                 fmt.Sprint(generateReq.Seed),
				NumOutputs:           fmt.Sprint(generateReq.NumOutputs),
				OutputImageExtension: string(shared.DEFAULT_GENERATE_OUTPUT_EXTENSION),
				OutputImageQuality:   fmt.Sprint(shared.DEFAULT_GENERATE_OUTPUT_QUALITY),
				ProcessType:          generateReq.ProcessType,
			},
		}

		start = time.Now()
		err = c.Redis.EnqueueCogRequest(r.Context(), cogReqBody)
		if err != nil {
			klog.Errorf("Failed to write request %s to queue: %v", requestId, err)
			responses.ErrInternalServerError(w, r, "Failed to queue generate request")
			return err
		}
		fmt.Printf("--- Enqueue cog request took: %s\n", time.Now().Sub(start))
		return nil
	}); err != nil {
		klog.Errorf("Error in transaction: %v", err)
		return
	}

	// Track the request in our internal map
	start = time.Now()
	c.Redis.SetCogRequestStreamID(r.Context(), requestId, generateReq.StreamID)
	fmt.Printf("--- Put request in map took: %s\n", time.Now().Sub(start))

	// Send live page update
	go c.Hub.BroadcastLivePageMessage(livePageMsg)

	// Start the timeout timer
	go func() {
		// sleep
		time.Sleep(shared.REQUEST_COG_TIMEOUT)
		// this will trigger timeout if it hasnt been finished
		c.Repo.FailCogMessageDueToTimeoutIfTimedOut(responses.CogStatusUpdate{
			Input:  cogReqBody.Input,
			Error:  "TIMEOUT",
			Status: responses.CogFailed,
		})
	}()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &responses.QueuedResponse{
		ID: requestId,
	})
}

// HTTP POST submit a generation to gallery
func (c *RestAPI) HandleSubmitGenerationToGallery(w http.ResponseWriter, r *http.Request) {
	userID := c.GetUserIDIfAuthenticated(w, r)
	if userID == nil {
		return
	}

	// Parse request body
	reqBody, _ := io.ReadAll(r.Body)
	var submitToGalleryReq requests.GenerateSubmitToGalleryRequestBody
	err := json.Unmarshal(reqBody, &submitToGalleryReq)
	if err != nil {
		responses.ErrUnableToParseJson(w, r)
		return
	}

	submitted, err := c.Repo.SubmitGenerationOutputsToGalleryForUser(submitToGalleryReq.GenerationOutputIDs, *userID)
	if err != nil {
		responses.ErrInternalServerError(w, r, "Error submitting generation outputs to gallery")
		return
	}

	res := responses.GenerateSubmitToGalleryResponse{
		Submitted: submitted,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
