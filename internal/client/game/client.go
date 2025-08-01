package game

import (
	"log/slog"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
)

type GameClient struct {
	conn           *websocket.Conn
	done           chan struct{}
	gameData       *GameData
	serverMessages []api.ServerMessage

	// channels
	UIUpdateCh chan UIUpdate
	responseCh chan api.Response
}

type Reason string

const (
	GameStarted Reason = "game_started"
	Animation   Reason = "animation"
	ServerMsg   Reason = "server_msg"
	GameOver    Reason = "game_over"
)

type UIUpdate struct {
	Reason Reason
	Data   any
}

func NewGameClient(conn *websocket.Conn) *GameClient {
	return &GameClient{
		conn:           conn,
		responseCh:     make(chan api.Response),
		UIUpdateCh:     make(chan UIUpdate),
		gameData:       &GameData{},
		serverMessages: []api.ServerMessage{},
	}
}

func (c *GameClient) IsPlayerTurn() bool {
	return c.gameData.Turn == int(c.gameData.Player.ID)
}

func (c *GameClient) GetGameData() *GameData {
	return c.gameData
}

func (c *GameClient) GetServerMessages() []api.ServerMessage {
	return c.serverMessages
}

func (c *GameClient) Close() {
	if c.done != nil {
		close(c.done)
	}
	if c.responseCh != nil {
		close(c.responseCh)
	}
	if c.UIUpdateCh != nil {
		close(c.UIUpdateCh)
	}
	slog.Info("game client closed")
}
