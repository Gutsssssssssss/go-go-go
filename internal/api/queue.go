package api

import "github.com/google/uuid"

const (
	QueueMessageMatchSuccess = "match_success"
	QueueMessageMatchFailed  = "match_failed"
)

type QueueMessage struct {
	Message string           `json:"message"`
	Data    QueueMessageData `json:"data"`
}

type QueueMessageData struct {
	Opponent uuid.UUID `json:"opponent,omitempty"`
}
