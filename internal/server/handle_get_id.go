package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/api"
)

// create userID and return it
func (s *Server) HandleGetID(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	respondWithJSON(w, 200, api.GetIDResponse{ID: id})
}
