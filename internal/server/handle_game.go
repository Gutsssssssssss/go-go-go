package server

import (
	"net/http"
)

func (s *Server) HandleGame(w http.ResponseWriter, r *http.Request) {
	// conn, err := s.upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Printf("Websocket upgrade failed: %s\n", err)
	// 	respondWithError(w, 400, "Bad request", nil)
	// 	return
	// }

	// read gameID from client

	// find gameID in map
	// game start at goroutine

}
