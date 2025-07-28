package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	game "github.com/yanmoyy/go-go-go/internal/client/game"
)

// StartWaiting starts waiting for a match and returns the connection of the game session
func StartWaiting(id uuid.UUID, ctx context.Context) (uuid.UUID, error) {
	c := getClient()
	idString := id.String()
	// web socket
	u := url.URL{Scheme: "ws", Host: c.wsHost, Path: "/ws/waiting/" + idString}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("dial: %w", err)
	}

	done := make(chan struct{})
	errCh := make(chan error)
	msgCh := make(chan api.QueueMessage)

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := conn.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				close(done)
				return
			}
			if err != nil {
				errCh <- err
				return
			}
			var msg api.QueueMessage
			if err := json.Unmarshal(message, &msg); err != nil {
				errCh <- fmt.Errorf("unmarshal: %w", err)
				return
			}
			msgCh <- msg
		}
	}()
	for {
		select {
		case err := <-errCh:
			return uuid.Nil, fmt.Errorf("read: %w", err)
		case msg := <-msgCh:
			switch msg.Message {
			case api.QueueMessageMatchSuccess:
				c.game = game.NewGameClient(conn)
				return msg.Data.Opponent, nil
			case api.QueueMessageMatchFailed:
				return uuid.Nil, nil
			default:
				return uuid.Nil, fmt.Errorf("unknown message: %s", msg.Message)
			}
		case <-ctx.Done():
			return uuid.Nil, nil
		case <-done:
			return uuid.Nil, nil
		}
	}
}
