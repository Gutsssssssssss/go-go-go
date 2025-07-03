package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.NewBoard("Hello World"))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
