package game

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

func sendChatMessage(conn *websocket.Conn, name, content string) error {
	return sendJSON(conn, api.Message{
		Type: api.ChatMsg,
		Data: api.ChatData{
			Name:    name,
			Content: content,
		},
	})
}

func sendGameEventRequest(conn *websocket.Conn, id string, evt game.Event) error {
	jsonData, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	err = sendJSON(conn, api.Message{
		Type: api.RequestMsg,
		Data: api.Request{
			ID:   id,
			Type: api.GameEventRequest,
			Data: jsonData,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	return nil
}

func sendJSON(conn *websocket.Conn, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		return err
	}
	return nil
}
