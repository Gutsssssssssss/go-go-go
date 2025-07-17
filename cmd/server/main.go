package main

import (
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/database"
	"github.com/yanmoyy/go-go-go/internal/logging"
	"github.com/yanmoyy/go-go-go/internal/server"
	"github.com/yanmoyy/go-go-go/internal/server/middleware"
)

func main() {
	const port = "8080"
	// temporary database
	db := database.NewMemoryDatabase()
	s := server.NewServer(db)
	go s.ListenMatchWaiting()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/user/id", middleware.Log(s.HandleGetID))
	mux.HandleFunc("/ws/waiting/{id}", middleware.Log(s.HandleWaiting))

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
