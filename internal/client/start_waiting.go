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
	matchCh := make(chan api.MatchData)

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
			var msg api.Message
			if err := json.Unmarshal(message, &msg); err != nil {
				errCh <- fmt.Errorf("unmarshal: %w", err)
				return
			}
			matchCh <- msg.Data.(api.MatchData)
		}
	}()
	for {
		select {
		case err := <-errCh:
			return uuid.Nil, fmt.Errorf("read: %w", err)
		case match := <-matchCh:
			switch match.Status {
			case api.MatchSuccess:
				c.game = game.NewGameClient(conn)
				return match.Opponent, nil
			case api.MatchFailed:
				return uuid.Nil, nil
			default:
				return uuid.Nil, fmt.Errorf("unknown status: %s", match.Status)
			}
		case <-ctx.Done():
			return uuid.Nil, nil
		case <-done:
			return uuid.Nil, nil
		}
	}
}
