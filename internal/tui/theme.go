package tui

import "github.com/charmbracelet/lipgloss"

type theme struct {
	PrimaryColor lipgloss.Color // main color of the theme
}

func newTheme() *theme {
	return &theme{
		PrimaryColor: lipgloss.Color(golangBlue),
	}
}
func GetTheme() *theme {
	return t
}

var t *theme

func init() {
	t = newTheme()
}
