package server

import (
	"bytes"
	"encoding/json"

	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait / 9) * 10
)

type sessionInfo struct {
	GameID uuid.UUID `json:"game_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (p *Client) readMessage() {
	defer func() {
		p.session.unregisterCh <- p
		sendCloseMessage(p.conn, "The session closed the message channel")
		p.conn.Close()
	}()

	p.conn.SetReadLimit(512)
	p.conn.SetReadDeadline(time.Now().Add(pongWait))
	p.conn.SetPongHandler(func(string) error { p.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("could not read the message", "err", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.ReplaceAll(message, []byte{'\n'}, []byte{' '}))
		p.session.broadcastCh <- message
	}
}

func (p *Client) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		sendCloseMessage(p.conn, "The session closed the message channel")
		p.conn.Close()
	}()
	for {
		select {
		case message, ok := <-p.messageCh:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				if err := p.conn.WriteJSON(api.QueueMessage{
					Message: "The session closed the message channel",
				}); err != nil {
					slog.Error("sending JSON", "err", err)
					sendCloseMessage(p.conn, "The session closed the message channel")
				}
			}

			// we have to decide the message form
			w, err := p.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(p.messageCh)
			for range n {
				w.Write([]byte{'\n'})
				w.Write(<-p.messageCh)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			p.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := p.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (s *Server) HandleGame(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}

	info := &sessionInfo{}
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	_, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			slog.Error("could not read the message", "err", err)
		}
		return
	}
	err = json.Unmarshal(message, info)
	if err != nil {
		slog.Error("could not parse the messages", "err", err)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "bad_request",
		}); err != nil {
			slog.Error("sending JSON", "err", err)
		}
		sendCloseMessage(conn, "bad request")
		return
	}

	gameID := info.GameID
	userID := info.UserID
	slog.Debug("Entered the game", "userID", userID, "gameID", gameID)

	session := s.sessions[gameID]
	if session == nil {
		slog.Debug("could not find the session", "gameID", gameID)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "not_found",
		}); err != nil {
			slog.Error("sending JSON", "err", err)
		}
		sendCloseMessage(conn, "not found")
		return
	}
	client := newClient(userID, conn, session)
	session.registerCh <- client

	go client.writeMessage()
	go client.readMessage()
}
