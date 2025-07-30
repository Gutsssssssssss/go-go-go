package api

import "github.com/google/uuid"

// http response for get id
type GetIDResponse struct {
	ID uuid.UUID `json:"id"`
}
