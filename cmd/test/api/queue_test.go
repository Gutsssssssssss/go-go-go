package main

import (
	"testing"

	"github.com/yanmoyy/go-go-go/internal/client"
)

func TestEnterQueue(t *testing.T) {
	c := client.NewClient(
		"http://localhost:8080",
		"localhost:8080",
	)
	id := "randomID"
	err := c.EnterQueue(id)
	if err != nil {
		t.Fatal(err)
	}
}
