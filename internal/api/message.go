package api

import (
	"encoding/json"
	"fmt"

	"github.com/yanmoyy/go-go-go/internal/game"
)

type MessageType string

const (
	MatchMessage     MessageType = "match"
	GameEventMessage MessageType = "game_event"
	ChatMessage      MessageType = "chat"
	RequestMessage   MessageType = "request"
	ResponseMessage  MessageType = "response"
)

type Message struct {
	Type MessageType `json:"type"`
	Data any         `json:"data"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type MessageType     `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	d, err := unmarshalData(temp.Type, temp.Data)
	if err != nil {
		return err
	}
	m.Type = temp.Type
	m.Data = d
	return nil
}

func unmarshalData(t MessageType, data []byte) (any, error) {
	switch t {
	case MatchMessage:
		var d MatchData
		err := json.Unmarshal(data, &d)
		return d, err
	case GameEventMessage:
		var d game.Event
		err := json.Unmarshal(data, &d)
		return d, err
	case ChatMessage:
		var d string
		err := json.Unmarshal(data, &d)
		return d, err
	case RequestMessage:
		var d Request
		err := json.Unmarshal(data, &d)
		return d, err
	case ResponseMessage:
		var d Response
		err := json.Unmarshal(data, &d)
		return d, err
	}
	return nil, fmt.Errorf("unknown message type: %s", t)
}
