package server

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/api"
)

type waiting struct {
	userID  uuid.UUID
	replyCh chan matchInfo
}

type matchInfo struct {
	player1 uuid.UUID
	player2 uuid.UUID
	gameID  uuid.UUID
}

func newMatchInfo(player1, player2 uuid.UUID) matchInfo {
	gameID := createGameID()
	return matchInfo{
		player1: player1,
		player2: player2,
		gameID:  gameID,
	}
}

func createGameID() uuid.UUID {
	return uuid.New()
}

func (s *Server) ListenMatchWaiting() {
	buf := make(map[uuid.UUID]waiting)
	for {
		select {
		case w := <-s.waitingQueue:
			buf[w.userID] = w
			if len(buf) >= 2 {
				// two waiting players
				var w1, w2 waiting
				for _, val := range buf {
					if w1.userID == uuid.Nil {
						w1 = val
					} else {
						w2 = val
						break
					}
				}
				info := newMatchInfo(w1.userID, w2.userID)
				w1.replyCh <- info
				w2.replyCh <- info
				// clear 2 waiting players
				delete(buf, w1.userID)
				delete(buf, w2.userID)
				log.Printf("Matched! %s and %s\n", w1.userID, w2.userID)
				s.createGameSession(info.gameID)
			}
		case userID := <-s.removeQueue:
			if _, ok := buf[userID]; ok {
				delete(buf, userID)
				log.Printf("Removed %s from waiting queue\n", userID)
			}
		}
	}
}

// create game session
func (s *Server) createGameSession(gameID uuid.UUID) {
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

	// TODO parsing idString from path and change to uuid
	// idString := r.PathValue("id")
	idString := uuid.New()

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
		s.removeQueue <- idString
		close(replyCh)
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
