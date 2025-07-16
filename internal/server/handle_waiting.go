package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type waiting struct {
	userID  uuid.UUID
	replyCh chan matchInfo
}

// matchInfo is used to send match info to the user
// right now, it's just a pair of userID
type matchInfo struct {
	opponent uuid.UUID
}

func (s *Server) ListenMatchWaiting() {
	buf := make(map[uuid.UUID]waiting)
	for {
		select {
		case w1 := <-s.waitingQueue:
			if len(buf) == 0 {
				buf[w1.userID] = w1
				continue
			}
			var w2 waiting
			for _, val := range buf {
				if w2.userID == uuid.Nil {
					w2 = val
					break
				}
			}
			// sessionID := uuid.New() // random uuid
			w1.replyCh <- matchInfo{opponent: w2.userID}
			w2.replyCh <- matchInfo{opponent: w1.userID}
			// clear waiting user
			delete(buf, w2.userID)
			slog.Info("Matched!", "user1", w1.userID, "user2", w2.userID)
			// s.createGameSession(sessionID)
		case userID := <-s.removeQueue:
			if _, ok := buf[userID]; ok {
				delete(buf, userID)
				slog.Info("Removed user from waiting queue", "userID", userID)
			}
		}
	}
}

// create game session
func (s *Server) createGameSession(sessionID uuid.UUID) {
	game := game.NewGame()
	session := newGameSession(game)
	s.sessions[sessionID] = session
	slog.Debug("Start listening the game session", "sessionID", sessionID)
	go session.ListenSession()
}

func (s *Server) HandleWaiting(w http.ResponseWriter, r *http.Request) {
	const timeout = 5 * time.Second

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}

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
		err := conn.WriteJSON(
			api.QueueMessage{
				Message: api.QueueMessageMatchSuccess,
				Data: api.QueueMessageData{
					Opponent: info.opponent,
				},
			})
		if err != nil {
			slog.Error("sending JSON", "err", err)
			sendCloseMessage(conn)
			conn.Close()
		}
		return
	case <-time.After(timeout):
		s.removeQueue <- userID
		close(replyCh)
		slog.Info("Can not find other player", "userID", userID)
		err := conn.WriteJSON(
			api.QueueMessage{Message: api.QueueMessageMatchFailed},
		)
		if err != nil {
			slog.Error("sending JSON", "err", err)
		}
		sendCloseMessage(conn)
		conn.Close()
		return
	}
}
