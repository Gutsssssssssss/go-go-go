package game

import (
	"log/slog"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type GameClient struct {
	conn     *websocket.Conn
	done     chan struct{}
	gameData *GameData

	// channels
	StartGameCh chan struct{} // check if the game is started
	AnimationCh chan *game.AnimationData
	responseCh  chan api.Response
}

func NewGameClient(conn *websocket.Conn) *GameClient {
	return &GameClient{
		conn:        conn,
		responseCh:  make(chan api.Response),
		AnimationCh: make(chan *game.AnimationData),
		StartGameCh: make(chan struct{}),
		gameData:    &GameData{},
	}
}

func (c *GameClient) IsPlayerTurn() bool {
	return c.gameData.Turn == int(c.gameData.Player.ID)
}

func (c *GameClient) GetGameData() *GameData {
	return c.gameData
}

func (c *GameClient) Close() {
	if c.done != nil {
		close(c.done)
	}
	if c.responseCh != nil {
		close(c.responseCh)
	}
	if c.AnimationCh != nil {
		close(c.AnimationCh)
	}
	slog.Info("game client closed")
}
