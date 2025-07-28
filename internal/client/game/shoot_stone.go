package game

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/game"
)


func (c *GameClient) ShootStone(stoneID int, degrees int, power int) error {
	if c.conn == nil {
		return fmt.Errorf("no connection")
	}
	data := game.ShootData{
		PlayerID: c.data.Player.ID,
		StoneID:  stoneID,
		Velocity: game.ConvertToVelocity(float64(degrees), float64(power)),
	}
	evt := game.Event{Type: game.Shoot, Data: data}
	msg, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	err = c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}	


