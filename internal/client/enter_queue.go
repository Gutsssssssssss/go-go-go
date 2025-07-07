package client

import (
	"io"
	"net/http"
)

func (c *Client) EnterQueue() error {
	req, err := http.NewRequest("POST", c.baseURL+"/queue/enter", nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// print response body
	println(string(body))
	return nil
}
