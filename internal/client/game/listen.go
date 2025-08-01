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

const responseTimeout = time.Second * 3

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
					case api.GameEventMsg:
						err = c.handleGameEvent(m.Data.(game.Event))
						if err != nil {
							slog.Error("handleGameEvent", "err", err)
						}
					case api.ResponseMsg:
						c.responseCh <- m.Data.(api.Response)
					case api.ServerMsg:
						serverMsg := m.Data.(api.ServerMessage)
						slog.Info("server message received", "message", serverMsg)
						c.serverMessages = append(c.serverMessages, serverMsg)
						c.UIUpdateCh <- UIUpdate{Reason: ServerMsg}
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
		c.UIUpdateCh <- UIUpdate{Reason: GameStarted}
	case game.ShootResult:
		data := evt.Data.(game.ShootResultData)
		c.UIUpdateCh <- UIUpdate{Reason: Animation, Data: &data.Animation}
		c.gameData.Stones = data.Stones
		c.gameData.Turn = data.Turn
	case game.GameOver:
		c.UIUpdateCh <- UIUpdate{Reason: GameOver}
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
