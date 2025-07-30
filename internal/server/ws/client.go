package ws

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
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

func (c *Client) Listen() {
	go c.ReadPump()
	go c.WritePump()
}

// ReadPump reads messages from the websocket connection
func (c *Client) ReadPump() {
	defer func() {
		c.session.Unregister(c)
		SendCloseMessage(c.conn, "client read finished")
	}()
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(
		func(string) error {
			_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		},
	)
	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("client: unexpected close error", "err", err)
			}
			slog.Debug("client: read message", "err", err)
			break
		}
		c.session.Send(message{clientID: c.id, payload: payload})
	}
}

// WritePump writes messages to the websocket connection
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		SendCloseMessage(c.conn, "The session closed the message channel")
	}()
	for {
		select {
		case message, ok := <-c.messageCh:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			slog.Debug("WritePump", "message", string(message))
			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					SendCloseWithError(c.conn, "connection write message failed", err)
				}
			}
			// we have to decide the message form
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				slog.Error("NextWriter", "err", err)
				return
			}
			_, err = w.Write(message)
			if err != nil {
				slog.Error("w.Write", "err", err)
				return
			}

			n := len(c.messageCh)
			for range n {
				_, _ = w.Write([]byte{'\n'})
				_, _ = w.Write(<-c.messageCh)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
