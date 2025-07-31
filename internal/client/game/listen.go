package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

const (
	responseTimeout = time.Second * 3
)

// start ping pong and message handling
func (c *GameClient) StartListenConn(done chan struct{}) error {
	if c.conn == nil {
		return fmt.Errorf("no connection")
	}
	c.done = done
	go func() {
		select {
		case <-c.done:
			_ = c.conn.Close()
			return
		default:
			for {
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
						slog.Error("read message", "err", err)
						return
					}
					return
				}
				decoder := json.NewDecoder(bytes.NewReader(message))
				for decoder.More() {
					var m api.Message
					err = decoder.Decode(&m)
					if err != nil {
						slog.Error("failed to unmarshal message", "err", err, "message", string(message))
						continue
					}
					switch m.Type {
					case api.GameEventMessage:
						err = c.handleGameEvent(m.Data.(game.Event))
						if err != nil {
							slog.Error("handleGameEvent", "err", err)
						}
					case api.ResponseMessage:
						c.responseCh <- m.Data.(api.Response)
					default:
						slog.Error("unknown message type", "type", m.Type)
					}
				}
			}
		}
	}()
	return nil
}

func (c *GameClient) handleGameEvent(evt game.Event) error {
	slog.Debug("handleGameEvent", "type", evt.Type)
	switch evt.Type {
	case game.PlayerStartGame:
		data := evt.Data.(game.PlayerStartGameData)
		c.gameData = &GameData{
			Turn:   data.Turn,
			Player: data.Player,
			Stones: data.Stones,
			Size:   data.Size,
		}
	case game.StoneAnimations:
		data := evt.Data.(game.StoneAnimationsData)
		c.gameData.Stones = data.FinalStones
		c.AnimationCh <- &data
	case game.TurnStart:
		data := evt.Data.(game.TurnStartData)
		c.gameData.Turn = data.Turn
	default:
		slog.Error("unknown event type", "type", evt.Type)
	}
	return nil
}

func (c *GameClient) waitResponse() (api.Response, error) {
	select {
	case resp := <-c.responseCh:
		return resp, nil
	case <-time.After(responseTimeout):
		return api.Response{}, fmt.Errorf("no response received")
	}
}
