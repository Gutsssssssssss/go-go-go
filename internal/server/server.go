package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader     websocket.Upgrader
	waitingQueue chan waiting
	sessions     map[string]*Session
}

func NewServer() *Server {
	return &Server{
		upgrader:     websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }},
		waitingQueue: make(chan waiting),
	}
}
