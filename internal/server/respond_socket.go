package server

import (
	"log/slog"

	"github.com/gorilla/websocket"
)

func sendCloseMessage(conn *websocket.Conn) {
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		slog.Error("sending close message", "err", err)
	}
	conn.Close()
}

