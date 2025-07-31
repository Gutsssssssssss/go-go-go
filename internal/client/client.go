package client

import (
	"net/http"

	game "github.com/yanmoyy/go-go-go/internal/client/game"
)

type Client struct {
	httpBase string
	wsBase   string
	http     *http.Client     // http client
	game     *game.GameClient // game client
}

func newClient(httpBase string, wsBase string) *Client {
	return &Client{
		http:     http.DefaultClient,
		httpBase: httpBase,
		wsBase:   wsBase,
	}
}

var singleton *Client

// singleton pattern
func getClient() *Client {
	if singleton != nil {
		return singleton
	}
	singleton = newClient("http://localhost:8080", "ws://localhost:8080")
	return singleton
}

func GetGameClient() *game.GameClient {
	c := getClient()
	if c.game == nil {
		panic("game client is nil")
	}
	return c.game
}
