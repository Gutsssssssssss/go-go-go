package api

import "time"

type ServerMsgType string

const (
	ServerChat ServerMsgType = "chat"
	ServerGame ServerMsgType = "game"
)

// ServerMessage is a message from server
type ServerMessage struct {
	Time    time.Time     `json:"time"`
	Type    ServerMsgType `json:"type"`
	Message string        `json:"message"`
}
