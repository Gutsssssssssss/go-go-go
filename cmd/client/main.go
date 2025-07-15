package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/logging"
	"github.com/yanmoyy/go-go-go/internal/tui"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()
	if *debug {
		f, err := tea.LogToFile("debug.log", "")
		prettyHandler := logging.NewPrettyHandler(f, logging.PrettyHandlerOptions{
			SlogOpts: slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		})
		logger := slog.New(prettyHandler)
		slog.SetDefault(logger)
		slog.SetLogLoggerLevel(slog.LevelDebug)
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
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
