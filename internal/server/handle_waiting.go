package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/game"
	"github.com/yanmoyy/go-go-go/internal/server/ws"
)

type waiting struct {
	userID  uuid.UUID
	replyCh chan matchInfo
}

// matchInfo is used to send match info to the user
// right now, it's just a pair of userID
type matchInfo struct {
	sessionID uuid.UUID
	opponent  uuid.UUID
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
			sessionID := uuid.New() // random uuid
			w1.replyCh <- matchInfo{opponent: w2.userID, sessionID: sessionID}
			w2.replyCh <- matchInfo{opponent: w1.userID, sessionID: sessionID}
			s.createGameSession(sessionID)
			// clear waiting user
			delete(buf, w2.userID)
			slog.Info("Matched!", "user1", w1.userID, "user2", w2.userID)
		case userID := <-s.removeQueue:
			delete(buf, userID)
			slog.Info("Removed user from waiting queue", "userID", userID)
		}
	}
}

// create game session
func (s *Server) createGameSession(sessionID uuid.UUID) {
	game := game.NewGame()
	session := ws.NewGameSession(game)
	s.sessions[sessionID] = session
	slog.Debug("Start listening the game session", "sessionID", sessionID)
	go session.ListenSession()
}

func (s *Server) registerClientToSession(sessionID, clientID uuid.UUID, conn *websocket.Conn) {
	client := ws.NewClient(clientID, conn, s.sessions[sessionID])
	s.sessions[sessionID].Register(client)
	go client.WriteMessage()
	go client.ReadMessage()
}

func (s *Server) HandleWaiting(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request", err)
		return
	}

	const timeout = 5 * time.Second

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request", err)
		return
	}

	replyCh := make(chan matchInfo)
	s.waitingQueue <- waiting{userID: userID, replyCh: replyCh}
	select {
	case info := <-replyCh:
		err := conn.WriteJSON(
			api.QueueMessage{
				Message: api.QueueMessageMatchSuccess,
				Data: &api.QueueMessageData{
					Opponent: info.opponent,
				},
			})
		if err != nil {
			ws.SendCloseWithError(conn, "couldn't send JSON", err)
			return
		}
		s.registerClientToSession(info.sessionID, userID, conn)

	case <-time.After(timeout):
		s.removeQueue <- userID
		close(replyCh)
		err := conn.WriteJSON(
			api.QueueMessage{Message: api.QueueMessageMatchFailed},
		)
		if err != nil {
			ws.SendCloseWithError(conn, "couldn't send JSON", err)
			return
		}
		ws.SendCloseMessage(conn, "match failed")
	}
}
