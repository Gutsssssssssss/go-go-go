package theme

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

type theme struct {
	PrimaryColor   lipgloss.Color // main color of the theme
	WaitingSpinner spinner.Model
}

func newTheme() *theme {
	return &theme{
		PrimaryColor: color.GolangBlue,
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
