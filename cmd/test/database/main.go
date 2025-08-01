package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yanmoyy/go-go-go/internal/database"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type dummy struct {
	ID   int
	Name string
	Age  int
}

func main() {
	dummy := dummy{
		ID:   1,
		Name: "jeon",
		Age:  20,
	}
	jsonData, err := json.Marshal(dummy)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn := "host=localhost port=5432 user=postgres password=123456 dbname=gogogo sslmode=disable"
	db, err := database.NewDB(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = database.CreateGameRecord(db, "test", []game.Event{
		game.Event{
			Type: game.PlayerStartGame,
			Data: dummy,
		},
	})

	fmt.Println(string(jsonData))
}
