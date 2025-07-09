package server

import (
	"log"
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/api"
)

type waiting struct {
	userID  string
	replyCh chan matchInfo
}

type matchInfo struct {
	player1 string
	player2 string
	gameID  string
}

func newMatchInfo(player1, player2 string) matchInfo {
	gameID := createGameID()
	return matchInfo{
		player1: player1,
		player2: player2,
		gameID:  gameID,
	}
}

// gotta be uuid or unique something
func createGameID() string {
	return ""
}

func (s *Server) ListenMatchWaiting() {
	var buf []waiting
	for w := range s.waitingQueue {
		buf = append(buf, w)
		if len(buf) >= 2 {
			// two waiting players
			w1, w2 := buf[0], buf[1]
			info := newMatchInfo(w1.userID, w2.userID)
			w1.replyCh <- info
			w2.replyCh <- info
			// clear 2 waiting players
			buf = buf[2:]
			log.Printf("Matched! %s and %s\n", w1.userID, w2.userID)
			s.createGameSession(info.gameID)
		}
	}
}

// create game session
func (s *Server) createGameSession(gameID string) {
	s.sessions[gameID] = NewSession()
}

func (s *Server) HandleWaiting(w http.ResponseWriter, r *http.Request) {
	// TODO: check if user's waiting is exist
	// if user stop waiting, remove user from waiting queue
	const timeout = 5 * time.Second

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade failed: %s\n", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}

	defer conn.Close()

	idString := r.PathValue("id")
	// TODO: change idString to more secure id (ex: uuid)
	replyCh := make(chan matchInfo)
	s.waitingQueue <- waiting{userID: idString, replyCh: replyCh}

	select {
	case info := <-replyCh:
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "match_success", GameID: info.gameID,
		}); err != nil {
			log.Printf("Error sending JSON: %s\n", err)
			sendCloseMessage(conn)
		}
		return
	case <-time.After(timeout):
		log.Println("Can not find other player")
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "match_failed",
		}); err != nil {
			log.Printf("Error sending JSON: %s\n", err)
		}
		sendCloseMessage(conn)
		return
	}
}
