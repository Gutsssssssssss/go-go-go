package game

import (
	"fmt"
	"time"

	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

func (c *GameClient) ShootStone(stoneID int, degrees int, power int) error {
	if c.conn == nil {
		return fmt.Errorf("no connection")
	}
	evt := game.Event{
		Type: game.PlayerShoot,
		Data: game.PlayerShootData{
			PlayerID: c.gameData.Player.ID,
			StoneID:  stoneID,
			Velocity: game.ConvertToVelocity(float64(degrees), float64(power)),
		},
	}
	reqID := fmt.Sprintf("shoot_%d_%d", stoneID, time.Now().UnixNano())
	err := sendGameEventRequest(c.conn, reqID, evt)
	if err != nil {
		return fmt.Errorf("failed to send game event: %w", err)
	}
	// check the server's response and return
	resp, err := c.waitResponse()
	if err != nil {
		return fmt.Errorf("waitResponse: %w", err)
	}
	if resp.Status != api.ResponseSuccess {
		return fmt.Errorf("failed to shoot stone: %s", resp.Message)
	}
	return nil
}
