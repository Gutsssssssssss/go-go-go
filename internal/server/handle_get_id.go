package server

import (
	"net/http"

	"github.com/google/uuid"
)

// create userID and return it
func (s *Server) HandleGetID(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID uuid.UUID `json:"id"`
	}

	id := uuid.New()
	respondWithJSON(w, 200, response{ID: id})
}
