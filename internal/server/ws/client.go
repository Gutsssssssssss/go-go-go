package ws

import (
	"bytes"
	"log/slog"
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

type Client struct {
	id        uuid.UUID
	conn      *websocket.Conn
	session   *Session
	messageCh chan []byte
}

func NewClient(id uuid.UUID, conn *websocket.Conn, session *Session) *Client {
	return &Client{
		id:        id,
		conn:      conn,
		session:   session,
		messageCh: make(chan []byte, 256),
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.session.unregisterCh <- c
		SendCloseMessage(c.conn, "The session closed the message channel")
		c.conn.Close()
	}()

	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("could not read the message", "err", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.ReplaceAll(message, []byte{'\n'}, []byte{' '}))
		c.session.broadcastCh <- message
	}
}

func (c *Client) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		SendCloseMessage(c.conn, "The session closed the message channel")
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.messageCh:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				if err := c.conn.WriteJSON(api.QueueMessage{
					Message: "The session closed the message channel",
				}); err != nil {
					slog.Error("sending JSON", "err", err)
					SendCloseMessage(c.conn, "The session closed the message channel")
				}
			}

			// we have to decide the message form
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.messageCh)
			for range n {
				w.Write([]byte{'\n'})
				w.Write(<-c.messageCh)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
