package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yanmoyy/go-go-go/internal/database"
	"github.com/yanmoyy/go-go-go/internal/logging"
	"github.com/yanmoyy/go-go-go/internal/server"
	"github.com/yanmoyy/go-go-go/internal/server/middleware"
	"github.com/yanmoyy/go-go-go/internal/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverPort := util.EnsureEnvExist("SERVER_PORT")

	var cfg server.Config
	useDB := flag.Bool("db", false, "use database")
	flag.Parse()
	if *useDB {
		dbHost := util.EnsureEnvExist("DB_HOST")
		dbPort := util.EnsureEnvExist("DB_PORT")
		dbUser := util.EnsureEnvExist("DB_USER")
		dbPassword := util.EnsureEnvExist("DB_PASSWORD")
		dbName := util.EnsureEnvExist("DB_NAME")
		conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
		db, err := database.NewDB(conn)
		if err != nil {
			log.Fatal(err)
		}
		cfg = server.Config{UseDB: true, DB: db}
	}

	s := server.NewServer(cfg)
	go s.ListenMatchWaiting()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/user/id", middleware.Log(s.HandleGetID))
	mux.HandleFunc("/ws/waiting/{id}", middleware.Log(s.HandleWaiting))

	srv := &http.Server{
		Addr:        ":" + serverPort,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	// set logger
	logging.SetPrettyDebugLogger()

	log.Printf("Listening on port %s\n", serverPort)
	log.Fatal(srv.ListenAndServe())
}
