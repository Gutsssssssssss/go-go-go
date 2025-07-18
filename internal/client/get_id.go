package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/api"
)

func (c *Client) GetID() (uuid.UUID, error) {
	url := c.baseURL + "/api/user/id"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("new request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = res.Body.Close() }()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return uuid.Nil, fmt.Errorf("read body: %w", err)
	}
	var resp api.GetIDResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unmarshal: %w", err)
	}

	if resp.ID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("nil uuid received: %w", err)
	}
	return resp.ID, nil
}
