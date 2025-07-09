package client

import (
	"net/http"
)

type Client struct {
	baseURL string
	wsHost  string
	client  *http.Client
}

func NewClient(baseURL string, wsHost string) *Client {
	return &Client{
		baseURL: baseURL,
		wsHost:  wsHost,
		client:  http.DefaultClient,
	}
}
