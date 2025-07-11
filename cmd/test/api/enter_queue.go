package main

import (
	"fmt"

	"github.com/yanmoyy/go-go-go/internal/client"
)

func testEnterQueue() error {
	c := client.NewClient(
		"http://localhost:8080",
		"localhost:8080",
	)
	id := "randomID"
	err := c.StartWaiting(id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := testEnterQueue()
	if err != nil {
		fmt.Println("Test failed")
		fmt.Println(err)
	}
}
