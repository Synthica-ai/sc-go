package requests

import "github.com/google/uuid"

type DeactiveApiTokenRequest struct {
	ID uuid.UUID `json:"id"`
}

type NewTokenRequest struct {
	Name   string `json:"name"`
	Public *bool  `json:"public"`
}

type UpdateTokenRequest struct {
	ID     string `json:"id"`
	Public *bool  `json:"public"`
}
