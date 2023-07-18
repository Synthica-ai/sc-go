package requests

import "github.com/google/uuid"

type DeleteGenerationRequest struct {
	GenerationOutputIDs []uuid.UUID `json:"generation_output_ids"`
}

type DeleteVoiceRequest struct {
	VoiceID uuid.UUID `json:"voice_id"`
}
