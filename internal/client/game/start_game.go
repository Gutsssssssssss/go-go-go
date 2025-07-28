package game

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/yanmoyy/go-go-go/internal/game"
)

// start ping pong and message handling
func (c *GameClient) StartGame(done chan struct{}) error {
	if c.conn == nil {
		return fmt.Errorf("no connection")
	}
	go func () {
		select {
		case <-done:
			return
		default:
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				slog.Error("read message", "err", err)
				return
			}
			evt, err := parseMessage(message)
			if err != nil {
				slog.Error("failed to parse message", "err", err)
				return
			}
			switch evt.Type {
			case game.PlayerStartGame:
				// TODO: handle player start game
				data := evt.Data.(game.PlayerStartGameData)
				c.data.Player = data.Player
				c.data.Turn = int(data.Turn)
			// case game.TurnStart:
				// TODO: handle turn start
			case game.StoneAnimations:
				// TODO: handle stone animations
			default:
				slog.Error("unknown event type", "type", evt.Type)
			}		
		}
	}
	}()
	return nil
}

// parseMessage parses a message to game event
func parseMessage(message []byte) (game.Event, error) {
	var evt game.Event
	err := json.Unmarshal(message, &evt)
	if err != nil {
		return evt, fmt.Errorf("failed to unmarshal message: %w", err)	
	}
	return evt, nil
}