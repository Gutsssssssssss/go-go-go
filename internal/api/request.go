package api

import "encoding/json"

const (
	GameEventRequest RequestType = "game_event"
)

type RequestType string

type Request struct {
	ID   string          `json:"id"`
	Type RequestType     `json:"type"`
	Data json.RawMessage `json:"data"`
}
