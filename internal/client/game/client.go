package game

import (
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type GameClient struct {
	conn        *websocket.Conn
	gameData    *GameData
	AnimationCh chan *game.StoneAnimationsData
	responseCh  chan api.Response
}

func NewGameClient(conn *websocket.Conn) *GameClient {
	return &GameClient{
		conn:        conn,
		responseCh:  make(chan api.Response),
		AnimationCh: make(chan *game.StoneAnimationsData),
	}
}

func (c *GameClient) IsPlayerTurn() bool {
	return c.gameData.Turn == int(c.gameData.Player.ID)
}

func (c *GameClient) GetGameData() *GameData {
	return c.gameData
}
