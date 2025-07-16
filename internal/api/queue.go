package api

import "github.com/google/uuid"

type QueueMessage struct {
	Message   string    `json:"message"`
	SessionID uuid.UUID `json:"session_id"`
}
