package ws

import (
	"log/slog"

	"github.com/yanmoyy/go-go-go/internal/game"
)

type Session struct {
	game         *game.Game
	clients      map[*Client]bool
	registerCh   chan *Client
	unregisterCh chan *Client
	broadcastCh  chan []byte
}

func NewGameSession(game *game.Game) *Session {
	return &Session{
		game:         game,
		clients:      make(map[*Client]bool),
		registerCh:   make(chan *Client),
		unregisterCh: make(chan *Client),
		broadcastCh:  make(chan []byte),
	}
}
func (s *Session) ListenSession() {
	for {
		select {
		case client := <-s.registerCh:
			s.clients[client] = true
			slog.Info("Session: Registered", "clientID", client.id)
		case client := <-s.unregisterCh:
			delete(s.clients, client)
			slog.Info("Session: Unregistered", "clientID", client.id)
		case message := <-s.broadcastCh:
			for client := range s.clients {
				select {
				case client.messageCh <- message:
				default: // message channel is full
					close(client.messageCh)
					delete(s.clients, client)
				}
			}
		}
	}
}

func (s *Session) Register(client *Client) {
	s.clients[client] = true
}
