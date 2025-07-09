package server

import (
	"github.com/yanmoyy/go-go-go/internal/game"
)

// type connections map[string]*websocket.Conn

// Session is a temporary game session
type Session struct {
	game *game.Game
	// conns map[string]*websocket.Conn // key is playerID
}

func NewSession() *Session {
	g := game.NewGame()
	return &Session{
		game: g,
	}
}
