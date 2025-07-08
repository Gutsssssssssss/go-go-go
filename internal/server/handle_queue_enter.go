package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn     *websocket.Conn
	playerID string
}

type matchQueue struct {
	player *Player
	reply  chan *Player
}

type Server struct {
	upgrader websocket.Upgrader
	mq       chan matchQueue
}

func NewServer() *Server {
	return &Server{
		upgrader: websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }},
		mq:       make(chan matchQueue),
	}
}

func NewPlayer(conn *websocket.Conn, playerID string) *Player {
	return &Player{
		conn:     conn,
		playerID: playerID,
	}
}

func (s *Server) RunMatcher() {
	var waiting *matchQueue
	for m := range s.mq {
		if waiting == nil {
			waiting = &m
			continue
		}
		waiting.reply <- m.player
		m.reply <- waiting.player
		waiting = nil
	}
}

func startGame(player *Player, peer *Player) {

}

func (s *Server) HandleQueueEnter(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed: %s\n", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}
	defer conn.Close()

	playerID := strings.TrimPrefix(r.URL.Path, "/api/queue/enter/")
	player := NewPlayer(conn, playerID)

	reply := make(chan *Player)
	s.mq <- matchQueue{player: player, reply: reply}

	select {
	case peer := <-reply:
		// matched
		fmt.Println("Matched!")
		startGame(player, peer)
	case <-time.After(5 * time.Second):
		// matching failed and send signal to play with cpu
		fmt.Println("Match failed!")

		return
	}
}
