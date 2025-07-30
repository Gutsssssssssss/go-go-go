package api

import "github.com/google/uuid"

const (
	MatchSuccess = "match_success"
	MatchFailed  = "match_failed"
)

type MatchData struct {
	Status   string    `json:"status"`
	Opponent uuid.UUID `json:"opponent,omitempty"`
}
