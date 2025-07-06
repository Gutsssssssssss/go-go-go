package cmp

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
)

type Button struct {
	text    string
	padding int
	margin  int
}

func (b Button) View(selected bool) string {
	t := theme.GetTheme()
	style := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Padding(0, b.padding).
		Margin(0, b.margin).
		Width(20).
		Border(lipgloss.RoundedBorder(), true)
	if selected {
		style = style.BorderForeground(lipgloss.Color(t.PrimaryColor)).
			Foreground(lipgloss.Color(t.PrimaryColor))
	}
	return style.Render(b.text)
}
