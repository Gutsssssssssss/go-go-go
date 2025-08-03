package server

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/server/ws"
)

type Server struct {
	upgrader     websocket.Upgrader
	waitingQueue chan waiting
	sessions     map[uuid.UUID]*ws.Session
	removeQueue  chan uuid.UUID
	cfg          Config
}

type Config struct {
	UseDB bool
	DB    *sql.DB
}

func NewServer(cfg Config) *Server {
	return &Server{
		upgrader:     websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }},
		waitingQueue: make(chan waiting),
		removeQueue:  make(chan uuid.UUID),
		sessions:     make(map[uuid.UUID]*ws.Session),
		cfg:          cfg,
	}
}
