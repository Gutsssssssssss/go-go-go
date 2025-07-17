package ws

import "testing"

func TestParseMessage(t *testing.T) {
	message := []byte("hello world")
	_, err := parseMessage(message)
	if err != nil {
		t.Error("failed to parse message")
	}
}
