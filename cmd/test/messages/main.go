package main

import (
	"fmt"
	"time"

	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/tui/view"
)

// test messages view

func main() {
	println("test messages view")
	messages := []api.ServerMessage{
		{
			Time:    time.Now(),
			Type:    api.ServerChat,
			Message: "game start",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerGame,
			Message: "Someone shoot stone",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerChat,
			Message: "Hi you are so cool",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerGame,
			Message: "game over",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerChat,
			Message: "Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message",
		},
	}

	fmt.Println(view.ServerMessages(messages, view.MessagesProps{
		Width:  100,
		Height: 50,
	}))
}
