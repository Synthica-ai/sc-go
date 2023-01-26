package requests

import "github.com/google/uuid"

type GenerateRequestBody struct {
	Prompt            string    `json:"prompt"`
	NegativePrompt    string    `json:"negative_prompt,omitempty"`
	Width             int32     `json:"width"`
	Height            int32     `json:"height"`
	NumInferenceSteps int32     `json:"num_inference_steps"`
	GuidanceScale     float32   `json:"guidance_scale"`
	ModelId           uuid.UUID `json:"model_id"`
	SchedulerId       uuid.UUID `json:"scheduler_id"`
	Seed              int       `json:"seed"`
	WebsocketId       string    `json:"websocket_id"`
	// ShouldSubmitToGallery bool    `json:"should_submit_to_gallery"`
}
