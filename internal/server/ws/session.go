package ws

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

const (
	maxBufferSize = 256
)

// Session represents a game session
type Session struct {
	game         *game.Game
	clients      map[uuid.UUID]*Client
	registerCh   chan *Client
	unregisterCh chan *Client
	messageCh    chan message
}

func NewGameSession() *Session {
	game := game.NewGame()
	return &Session{
		game:         game,
		clients:      make(map[uuid.UUID]*Client),
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
				s.clients[client.id] = client
				slog.Info("Session: Registered", "clientID", client.id)
				needStart, err := s.game.AddPlayer(client.id.String())
				if err != nil {
					slog.Error("failed to add player", "err", err)
				}
				if needStart {
					s.game.StartGame()
					for id := range s.clients {
						evt := s.game.GetPlayerStartGameEvent(id.String())
						s.sendGameEvent(id, evt)
					}
				}
			case client := <-s.unregisterCh:
				delete(s.clients, client.id)
				slog.Info("Session: Unregistered", "clientID", client.id)
			case msg := <-s.messageCh:
				err := s.handleMessage(msg)
				if err != nil {
					slog.Error("handleMessage", "err", err)
				}
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

func (s *Session) broadcastGameEvent(evt game.Event) {
	for id := range s.clients {
		s.sendClientWithJSON(id, api.Message{
			Type: api.GameEventMessage,
			Data: evt,
		})
	}
}

func (s *Session) sendClientWithJSON(clientID uuid.UUID, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal data", "err", err)
		return
	}
	client := s.clients[clientID]
	client.messageCh <- jsonData
}

func (s *Session) sendGameEvent(clientID uuid.UUID, evt game.Event) {
	s.sendClientWithJSON(clientID, api.Message{
		Type: api.GameEventMessage,
		Data: evt,
	})
}

func (s *Session) sendResponse(clientID uuid.UUID, id string, status api.ResponseStatus, message string) {
	s.sendClientWithJSON(clientID,
		api.Message{
			Type: api.ResponseMessage,
			Data: api.Response{
				ID:      id,
				Status:  status,
				Message: message,
			},
		},
	)
}

// handleMessage handles a message from a client
func (s *Session) handleMessage(msg message) error {
	var m api.Message
	err := json.Unmarshal(msg.payload, &m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	switch m.Type {
	case api.RequestMessage:
		err = s.handleRequest(msg.clientID, m.Data.(api.Request))
		if err != nil {
			return fmt.Errorf("handleRequest: %w", err)
		}
	case api.ChatMessage:
		slog.Info("ChatEvent", "message", string(msg.payload))
	}
	return nil
}

func (s *Session) handleRequest(clientID uuid.UUID, req api.Request) error {
	switch req.Type {
	case api.GameEventRequest:
		var evt game.Event
		if err := json.Unmarshal(req.Data, &evt); err != nil {
			return fmt.Errorf("failed to unmarshal event: %w", err)
		}
		nxtEvt, err := s.handleGameEvent(evt)
		if err != nil {
			s.sendResponse(clientID, req.ID, api.ResponseFailed, "failed to handle game event")
			return err
		}
		s.sendResponse(clientID, req.ID, api.ResponseSuccess, "game event successfully handled")
		s.broadcastGameEvent(nxtEvt)
	default:
		return fmt.Errorf("unknown request type: %s", req.Type)
	}
	return nil
}

// handleGameEvent handles a game event, and returns the next event and whether it needs to be broadcasted
func (s *Session) handleGameEvent(evt game.Event) (nxt game.Event, err error) {
	switch evt.Type {
	case game.Shoot:
		nxt, err = s.game.ShootStone(evt.Data.(game.ShootData))
		if err != nil {
			return game.Event{}, err
		}
	default:
		return game.Event{}, fmt.Errorf("unknown event type: %s", evt.Type)
	}
	return nxt, nil
}
