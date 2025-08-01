package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/yanmoyy/go-go-go/internal/logging"
	"github.com/yanmoyy/go-go-go/internal/tui"
	"github.com/yanmoyy/go-go-go/internal/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	util.EnsureEnvExist("HTTP_BASE")
	util.EnsureEnvExist("WS_BASE")
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()
	if *debug {
		f, err := tea.LogToFile("debug.log", "")
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		prettyHandler := logging.NewPrettyHandler(f, opts)
		logger := slog.New(prettyHandler)
		slog.SetDefault(logger)
		slog.SetLogLoggerLevel(slog.LevelDebug)
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer func() { _ = f.Close() }()
	} else {
		// discard all logs
		logger := slog.New(slog.DiscardHandler)
		slog.SetDefault(logger)
	}
	p := tea.NewProgram(tui.New())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
