package server

import "github.com/gorilla/websocket"

func sendCloseMessage(conn *websocket.Conn) error {
	return conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
