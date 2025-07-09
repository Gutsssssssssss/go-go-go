package api

type QueueMessage struct {
	Message string `json:"message"`
	GameID  string `json:"game_id"`
}