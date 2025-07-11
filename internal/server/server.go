package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/database"
)

type Server struct {
	upgrader     websocket.Upgrader
	waitingQueue chan waiting
	sessions     map[uuid.UUID]*Session
	removeQueue  chan uuid.UUID
	db           database.Database
}

func NewServer(db database.Database) *Server {
	return &Server{
		upgrader:     websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }},
		waitingQueue: make(chan waiting),
		removeQueue:  make(chan uuid.UUID),
		sessions:     make(map[uuid.UUID]*Session),
		db:           db,
	}
}
