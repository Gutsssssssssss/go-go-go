package view

import "github.com/charmbracelet/lipgloss"

type BoardProps struct {
	Width       int
	Height      int
	BorderColor lipgloss.Color
	Padding     int
	Margin      int
}

// Board is a wrapper around lipgloss.Style.Render
// It always has a border, so width and height are reduced by 2 each
func Board(p BoardProps, content string) string {
	return lipgloss.NewStyle().
		Width(p.Width-2).
		Height(p.Height-2).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(p.BorderColor).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Top).
		Padding(p.Padding).
		Margin(p.Margin).
		Render(content)
}
