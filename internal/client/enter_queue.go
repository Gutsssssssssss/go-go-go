package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yanmoyy/go-go-go/internal/api"
)

func (c *Client) EnterQueue(id string) error {
	// web socket
	u := url.URL{Scheme: "ws", Host: c.wsHost, Path: "/api/waiting/" + id}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	// get response from server with conn
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var msg api.QueueMessage
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("Error unmarshalling JSON: %s\n", err)
				return
			}
			if msg.Message == "match_success" {
				log.Printf("Successfully matched! %s\n", msg.GameID)
				// TODO: Game start
				return
			}
			if msg.Message == "match_failed" {
				log.Printf("Failed to match! %s\n", msg.GameID)
				return
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return nil
		case t := <-ticker.C:
			if err := conn.WriteMessage(websocket.TextMessage, []byte(t.String())); err != nil {
				return err
			}
		}
	}
}
