package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
)

func (c *Client) StartWaiting(id uuid.UUID, ctx context.Context) (uuid.UUID, error) {
	idString := id.String()
	// web socket
	u := url.URL{Scheme: "ws", Host: c.wsHost, Path: "/ws/waiting/" + idString}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	errCh := make(chan error)
	msgCh := make(chan api.QueueMessage)

	go func() {
		for {
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
				slog.Debug("recv", "message", string(message))
				var msg api.QueueMessage
				if err := json.Unmarshal(message, &msg); err != nil {
					errCh <- fmt.Errorf("unmarshal: %w", err)
					return
				}
				msgCh <- msg
			}
		}
	}()
	for {
		select {
		case err := <-errCh:
			return uuid.Nil, fmt.Errorf("read: %w", err)
		case msg := <-msgCh:
			if msg.Message == "match_success" {
				return msg.GameID, nil
			}
			if msg.Message == "match_fail" {
				return uuid.Nil, nil
			}
		case <-ctx.Done():
			return uuid.Nil, nil
		case <-done:
			return uuid.Nil, nil
		}
	}
}
