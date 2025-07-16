package server

import (
	"log/slog"

	"github.com/gorilla/websocket"
)

func sendCloseWithError(conn *websocket.Conn, msg string, err error) {
	if err != nil {
		slog.Error("sendCloseMessage", "err", err)
	}
	err = conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, msg),
	)
	if err != nil {
		slog.Error("sending close message", "err", err)
	}
	conn.Close()
}

func sendCloseMessage(conn *websocket.Conn, msg string) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, msg),
	)
	if err != nil {
		slog.Error("sending close message", "err", err)
	}
	conn.Close()
}
