package server

import (
	"log/slog"
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
				slog.Info("Matched!", "user1", w1.userID, "user2", w.userID)
				s.createGameSession(info.gameID)
			}
		case userID := <-s.removeQueue:
			if _, ok := buf[userID]; ok {
				delete(buf, userID)
				slog.Info("Removed %s from waiting queue", "userID", userID)
			}
		}
	}
}

// create game session
func (s *Server) createGameSession(gameID uuid.UUID) {
	session := NewSession()
	s.sessions[gameID] = session
	slog.Debug("Start listening the session", "gameID", gameID)
	go session.ListenSession()
}

func (s *Server) HandleWaiting(w http.ResponseWriter, r *http.Request) {
	// TODO: check if user's waiting is exist
	// if user stop waiting, remove user from waiting queue
	const timeout = 5 * time.Second

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}

	defer conn.Close()

	userID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.Error("parsing userID", "err", err)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "not_found",
		}); err != nil {
			slog.Error("Sending JSON", "err", err)
		}
		sendCloseMessage(conn)
		return
	}

	replyCh := make(chan matchInfo)
	s.waitingQueue <- waiting{userID: userID, replyCh: replyCh}

	select {
	case info := <-replyCh:
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "match_success", GameID: info.gameID,
		}); err != nil {
			slog.Error("Sending json", "err", err)
			sendCloseMessage(conn)
		}
		return
	case <-time.After(timeout):
		s.removeQueue <- userID
		close(replyCh)
		slog.Debug("Can not find other player", "userID", userID)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "match_failed",
		}); err != nil {
			slog.Error("sending JSON", "err", err)
		}
		sendCloseMessage(conn)
		return
	}
}
