package game

import (
	"fmt"
)

func (c *GameClient) Chat(content string) error {
	if c.conn == nil {
		return fmt.Errorf("no connection")
	}
	err := sendChatMessage(c.conn, c.gameData.Player.StoneType.String(), content)
	if err != nil {
		return fmt.Errorf("failed to send chat: %w", err)
	}
	return nil
}
