package api

import (
	"encoding/json"
	"fmt"

	"github.com/yanmoyy/go-go-go/internal/game"
)

type MsgType string

const (
	MatchMsg     MsgType = "match"
	GameEventMsg MsgType = "game_event"
	ChatMsg      MsgType = "chat"   // chat message from client
	ServerMsg    MsgType = "server" // log message from server
	RequestMsg   MsgType = "request"
	ResponseMsg  MsgType = "response"
)

type Message struct {
	Type MsgType `json:"type"`
	Data any     `json:"data"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type MsgType         `json:"type"`
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

func unmarshalData(t MsgType, data []byte) (any, error) {
	switch t {
	case MatchMsg:
		var d MatchData
		err := json.Unmarshal(data, &d)
		return d, err
	case GameEventMsg:
		var d game.Event
		err := json.Unmarshal(data, &d)
		return d, err
	case ChatMsg:
		var d ChatData
		err := json.Unmarshal(data, &d)
		return d, err
	case ServerMsg:
		var d ServerMessage
		err := json.Unmarshal(data, &d)
		return d, err
	case RequestMsg:
		var d Request
		err := json.Unmarshal(data, &d)
		return d, err
	case ResponseMsg:
		var d Response
		err := json.Unmarshal(data, &d)
		return d, err
	}
	return nil, fmt.Errorf("unknown message type: %s", t)
}
