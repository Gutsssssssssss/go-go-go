package server

import (
	"bytes"
	"encoding/json"
	"log"
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

func (p *Player) readMessage() {
	defer func() {
		p.session.unregisterCh <- p
		sendCloseMessage(p.conn)
		p.conn.Close()
	}()

	p.conn.SetReadLimit(512)
	p.conn.SetReadDeadline(time.Now().Add(pongWait))
	p.conn.SetPongHandler(func(string) error { p.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v\n", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.ReplaceAll(message, []byte{'\n'}, []byte{' '}))
		p.session.broadcastCh <- message
	}
}

func (p *Player) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		sendCloseMessage(p.conn)
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
					log.Printf("Error sending JSON: %s\n", err)
					sendCloseMessage(p.conn)
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
		log.Printf("Websocket upgrade failed: %s\n", err)
		respondWithError(w, 400, "Bad request", nil)
		return
	}

	info := &sessionInfo{}
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	_, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("Error: %v\n", err)
		}
		return
	}
	err = json.Unmarshal(message, info)
	if err != nil {
		log.Printf("Can not parse the messages: %s", err)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "bad_request",
		}); err != nil {
			log.Printf("Error sending JSON: %s\n", err)
		}
		sendCloseMessage(conn)
		return
	}

	gameID := info.GameID
	userID := info.UserID
	log.Printf("User %s entered Game %s", userID, gameID)

	session := s.sessions[gameID]
	if session == nil {
		log.Printf("Can't find the session, gameID: %s\n", gameID)
		if err := conn.WriteJSON(api.QueueMessage{
			Message: "not_found",
		}); err != nil {
			log.Printf("Error sending JSON: %s\n", err)
		}
		sendCloseMessage(conn)
		return
	}
	player := &Player{id: userID, conn: conn, session: session, messageCh: make(chan []byte)}
	session.registerCh <- player

	go player.writeMessage()
	go player.readMessage()
}
