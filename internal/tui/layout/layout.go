package layout

import (
	"github.com/charmbracelet/lipgloss"
)

func Column(children ...string) string {
	return lipgloss.JoinVertical(lipgloss.Center, children...)
}

func Row(children ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, children...)
}

func GapV(size int) string {
	return lipgloss.NewStyle().Height(size).Render("")
}

func GapH(size int) string {
	return lipgloss.NewStyle().Width(size).Render("")
}

type Sizable interface {
	SetSize(width, height int)
	GetSize() (int, int)
}
