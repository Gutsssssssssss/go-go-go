package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

func Help(m *help.Model, k help.KeyMap) string {
	m.Styles.ShortKey = lipgloss.NewStyle().Foreground(color.White)
	m.Styles.ShortDesc = lipgloss.NewStyle().Foreground(color.Gray)
	return lipgloss.NewStyle().
		MarginLeft(2).MarginBottom(1).
		Render(m.View(k))
}
