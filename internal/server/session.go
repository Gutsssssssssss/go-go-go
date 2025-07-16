package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type Session struct {
	game         *game.Game
	clients      map[*Client]bool
	registerCh   chan *Client
	unregisterCh chan *Client
	broadcastCh  chan []byte
}

type Client struct {
	id        uuid.UUID
	conn      *websocket.Conn
	session   *Session
	messageCh chan []byte
}

func newClient(id uuid.UUID, conn *websocket.Conn, session *Session) *Client {
	return &Client{
		id:        id,
		conn:      conn,
		session:   session,
		messageCh: make(chan []byte, 256),
	}
}

func newGameSession(game *game.Game) *Session {
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
		case client := <-s.unregisterCh:
			delete(s.clients, client)
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
