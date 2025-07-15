package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
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
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return uuid.Nil, fmt.Errorf("read body: %w", err)
	}
	type response struct {
		ID uuid.UUID `json:"id"`
	}
	var resp response
	err = json.Unmarshal(data, &res)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unmarshal: %w", err)
	}
	return resp.ID, nil
}
