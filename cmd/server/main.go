package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yanmoyy/go-go-go/internal/server"
)

func main() {
	const port = "8080"
	s := server.NewServer()
	go s.RunMatcher()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/queue/enter/", s.HandleQueueEnter)

	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
