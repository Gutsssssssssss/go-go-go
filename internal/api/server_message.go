package api

import (
	"fmt"
	"time"
)

type ServerMsgType string

const (
	ServerChat ServerMsgType = "chat"
	ServerGame ServerMsgType = "game"
)

// ServerMessage is a message from server
type ServerMessage struct {
	Time    time.Time     `json:"time"`
	From    string        `json:"from"`
	Type    ServerMsgType `json:"type"`
	Message string        `json:"message"`
}

func (m ServerMessage) String() string {
	return fmt.Sprintf("ServerMessage{Time: %s, From: %s, Type: %s, Message: %s}", m.Time, m.From, m.Type, m.Message)
}
