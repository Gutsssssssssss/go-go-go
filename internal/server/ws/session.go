package ws

import (
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/game"
)

const (
	maxBufferSize = 256
)

// Session represents a game session
type Session struct {
	game         *game.Game
	clients      map[*Client]bool
	registerCh   chan *Client
	unregisterCh chan *Client
	messageCh    chan message
}

func NewGameSession() *Session {
	game := game.NewGame()
	return &Session{
		game:         game,
		clients:      make(map[*Client]bool),
		registerCh:   make(chan *Client),
		unregisterCh: make(chan *Client),
		messageCh:    make(chan message, maxBufferSize),
	}
}

type message struct {
	clientID uuid.UUID
	payload  []byte
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
				needStart, err := s.game.AddPlayer(client.id.String())
				if err != nil {
					slog.Error("failed to add player", "err", err)
				}
				if needStart {
					evt := s.game.StartGame()
					s.broadcastWithJSON(evt)
				}
			case client := <-s.unregisterCh:
				delete(s.clients, client)
				slog.Info("Session: Unregistered", "clientID", client.id)
			case msg := <-s.messageCh:
				s.handleMessage(msg)
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

// Send sends a message to session from a client
func (s *Session) Send(msg message) {
	s.messageCh <- msg
}

func (s *Session) broadcastWithJSON(data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal data", "err", err)
		return
	}
	for client := range s.clients {
		client.messageCh <- jsonData
	}
}

// handleMessage handles a message from a client
func (s *Session) handleMessage(msg message) {
	t, n, err := getMessageType(msg.payload)
	if err != nil {
		slog.Error("failed to get message type", "err", err)
		return
	}
	switch t {
	case GameEvent:
		slog.Info("GameEvent", "message", string(msg.payload[n:]))
		// TODO: handle game events
	case ChatEvent:
		slog.Info("ChatEvent", "message", string(msg.payload[n:]))
		// TODO: handle chat events
	}
}
