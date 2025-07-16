package api

import "github.com/google/uuid"

type GetIDResponse struct {
	ID uuid.UUID `json:"id"`
}
