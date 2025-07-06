package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

type theme struct {
	PrimaryColor lipgloss.Color // main color of the theme
}

func newTheme() *theme {
	return &theme{
		PrimaryColor: color.GolangBlue,
	}
}
func GetTheme() *theme {
	return t
}

var t *theme

func init() {
	t = newTheme()
}
