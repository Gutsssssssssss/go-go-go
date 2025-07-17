package ws

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
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
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("client: cannot read message", "err", err)
			}
			break
		}
		c.session.Send(message)
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
