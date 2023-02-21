package shared

import (
	"time"
)

type LivePageStatus string

const (
	LivePageQueued     LivePageStatus = "queued"
	LivePageProcessing LivePageStatus = "processing"
	LivePageSucceeded  LivePageStatus = "succeeded"
	LivePageFailed     LivePageStatus = "failed"
	LivePageNSFW       LivePageStatus = "nsfw"
)

type LivePageMessage struct {
	ProcessType      ProcessType    `json:"process_type"`
	ID               string         `json:"id"`
	CountryCode      string         `json:"country_code"`
	Status           LivePageStatus `json:"status"`
	Width            int32          `json:"width"`
	Height           int32          `json:"height"`
	TargetNumOutputs int32          `json:"target_num_outputs"`
	ActualNumOutputs int            `json:"actual_num_outputs"`
	CreatedAt        time.Time      `json:"created_at"`
	StartedAt        *time.Time     `json:"started_at,omitempty"`
	CompletedAt      *time.Time     `json:"completed_at,omitempty"`
}
