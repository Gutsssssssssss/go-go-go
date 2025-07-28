package ws

import (
	"log/slog"

	"github.com/gorilla/websocket"
)

func SendCloseWithError(conn *websocket.Conn, msg string, err error) {
	if err != nil {
		slog.Error(msg, "err", err)
	}
	_ = conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, msg),
	)
	_ = conn.Close()
}

func SendCloseMessage(conn *websocket.Conn, msg string) {
	slog.Debug(msg)
	_ = conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, msg),
	)
	_ = conn.Close()
}
