package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type UserInfo struct {
	Nickname string    `json:"nickname"`
	GameID   uuid.UUID `json:"gameID"`
}

// I think it have to be changed to websocket event
func (s *Server) HandleCreateGameRecord(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userInfo UserInfo
	if err := decoder.Decode(&userInfo); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	session := s.sessions[userInfo.GameID]

	if err := s.db.SaveGameRecord(userInfo.Nickname, session.Game.Record); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't save the game record", err)
		return
	}

}
