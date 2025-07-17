package ws

import (
	"log/slog"

	"github.com/yanmoyy/go-go-go/internal/game"
)

// Session represents a game session
type Session struct {
	game         *game.Game
	clients      map[*Client]bool
	registerCh   chan *Client
	unregisterCh chan *Client
	messageCh    chan []byte
}

func NewGameSession(game *game.Game) *Session {
	return &Session{
		game:         game,
		clients:      make(map[*Client]bool),
		registerCh:   make(chan *Client),
		unregisterCh: make(chan *Client),
		messageCh:    make(chan []byte),
	}
}

// ListenSession listens for new clients and broadcasts messages to all clients
// by using go routines
func (s *Session) Listen() {
	go func() {
		for {
			select {
			case client := <-s.registerCh:
				s.clients[client] = true
				slog.Info("Session: Registered", "clientID", client.id)
			case client := <-s.unregisterCh:
				delete(s.clients, client)
				slog.Info("Session: Unregistered", "clientID", client.id)
			case message := <-s.messageCh:
				s.handleMessage(message)
			}
		}
	}()
}

// Register registers a client to the session
func (s *Session) Register(client *Client) {
	s.registerCh <- client
}

// Unregister unregisters a client from the session
func (s *Session) Unregister(client *Client) {
	s.unregisterCh <- client
}

// Broadcast broadcasts a message to all clients
func (s *Session) Send(message []byte) {
	s.messageCh <- message
}

// handleMessage handles a message from a client
func (s *Session) handleMessage(message []byte) {
	_, err := parseMessage(message)
	if err != nil {
		slog.Error("failed to parse message", "err", err)
		return
	}
}
