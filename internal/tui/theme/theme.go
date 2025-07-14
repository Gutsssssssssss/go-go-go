package theme

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

type theme struct {
	PrimaryColor   lipgloss.Color // default: GolangBlue
	DefaultColor   lipgloss.Color // default: White
	DisabledColor  lipgloss.Color // default: Gray
	WaitingSpinner spinner.Model  // default: spinner.Points with primary color
}

func newTheme() *theme {
	return &theme{
		PrimaryColor:  color.GolangBlue,
		DefaultColor:  color.White,
		DisabledColor: color.Gray,
		WaitingSpinner: spinner.New(spinner.WithSpinner(spinner.Points), spinner.WithStyle(
			lipgloss.NewStyle().Foreground(color.GolangBlue)),
		),
	}
}
func GetTheme() *theme {
	return t
}

var t *theme

func init() {
	t = newTheme()
}
