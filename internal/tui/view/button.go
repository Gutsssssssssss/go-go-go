package view

import (
	"github.com/charmbracelet/lipgloss"
)

type ButtonProps struct {
	Text  string
	Color lipgloss.Color
}

func Button(p ButtonProps) string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Padding(0, 1).
		Margin(0, 1).
		Width(20).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(p.Color).
		Foreground(p.Color)
	return style.Render(p.Text)
}
