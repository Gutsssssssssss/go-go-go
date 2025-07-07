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
	mux := http.NewServeMux()
	mux.HandleFunc("/queue/enter", server.HandleQueueEnter)

	srv := &http.Server{
		Addr:        ":" + port,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
