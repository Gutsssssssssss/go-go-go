package main

import (
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/database"
	"github.com/yanmoyy/go-go-go/internal/logging"
	"github.com/yanmoyy/go-go-go/internal/server"
)

func main() {
	const port = "8080"
	// temporary database
	db := database.NewMemoryDatabase()
	s := server.NewServer(db)
	go s.ListenMatchWaiting()
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/waiting/{id}", s.HandleWaiting)
	mux.HandleFunc("/ws/game", s.HandleGame)
	mux.HandleFunc("GET /api/user/id", s.HandleGetID)

	mux.HandleFunc("POST /api/game", s.HandleCreateGameRecord)

	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	// set logger
	logging.SetPrettyDebugLogger()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
