package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type Session struct {
	game         *game.Game
	players      map[*Player]bool
	registerCh   chan *Player
	unregisterCh chan *Player
	broadcastCh  chan []byte
}

type Player struct {
	id        uuid.UUID
	conn      *websocket.Conn
	session   *Session
	messageCh chan []byte
}

func NewSession() *Session {
	g := game.NewGame()
	return &Session{
		game:         g,
		players:      make(map[*Player]bool),
		registerCh:   make(chan *Player),
		unregisterCh: make(chan *Player),
		broadcastCh:  make(chan []byte),
	}
}

func (s *Session) ListenSession() {
	for {
		select {
		case player := <-s.registerCh:
			s.players[player] = true
		case player := <-s.unregisterCh:
			delete(s.players, player)
		case message := <-s.broadcastCh:
			for player := range s.players {
				select {
				case player.messageCh <- message:
				default:
					close(player.messageCh)
					delete(s.players, player)
				}
			}
		}
	}
}
