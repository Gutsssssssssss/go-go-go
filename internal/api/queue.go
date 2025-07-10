package api

import "github.com/google/uuid"

type QueueMessage struct {
	Message string    `json:"message"`
	GameID  uuid.UUID `json:"game_id"`
}
