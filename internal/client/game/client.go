package game

import (
	"github.com/gorilla/websocket"
)


type GameClient struct {
	conn *websocket.Conn
	data GameData 
}

func NewGameClient(conn *websocket.Conn) *GameClient {
	return &GameClient{
		conn: conn,
	}
}

func (c *GameClient) IsPlayerTurn() bool {
	return c.data.Turn== int(c.data.Player.ID)
}

func (c *GameClient) GetGameData() GameData {
	return c.data
}

