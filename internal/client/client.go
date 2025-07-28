package client

import (
	"net/http"

	game "github.com/yanmoyy/go-go-go/internal/client/game"
)

type Client struct {
	baseURL string
	wsHost  string
	http *http.Client // http client
	game *game.GameClient // game client
}

func newClient(baseURL string, wsHost string) *Client {
	return &Client{
		baseURL: baseURL,
		wsHost:  wsHost,
		http:  http.DefaultClient,
	}
}

var singleton *Client
// singleton pattern
func getClient() *Client {
	if singleton != nil {
		return singleton
	}
	singleton = newClient("http://localhost:8080", "localhost:8080")
	return singleton 
}

func GetGameClient() *game.GameClient {
	c := getClient()
	return c.game
}
