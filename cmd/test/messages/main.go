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
			From:    "Black",
			Message: "HI, I'm Black",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerChat,
			From:    "White",
			Message: "Hi, I'm White",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerGame,
			From:    "Server",
			Message: "game over",
		},
		{
			Time:    time.Now(),
			Type:    api.ServerChat,
			From:    "Black",
			Message: "Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message Very Long Message",
		},
	}

	fmt.Println(view.ServerMessages(messages, view.MessagesProps{
		Width:  100,
		Height: 50,
	}))
}
