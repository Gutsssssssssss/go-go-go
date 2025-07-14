package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/tui"
)

func main() {

	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	if *debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
	p := tea.NewProgram(tui.New())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
