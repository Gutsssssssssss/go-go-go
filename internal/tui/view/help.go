package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

type HelpProps struct {
	KeyMap help.KeyMap
	Width  int
}

func Help(m *help.Model, props HelpProps) string {
	m.Styles.ShortKey = lipgloss.NewStyle().Foreground(color.White)
	m.Styles.ShortDesc = lipgloss.NewStyle().Foreground(color.Gray)
	return lipgloss.NewStyle().
		MarginLeft(2).MarginBottom(1).
		Width(props.Width).
		AlignHorizontal(lipgloss.Center).
		Render(m.View(props.KeyMap))
}
