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

type Sizable interface {
	SetSize(width, height int)
	GetSize() (int, int)
}
