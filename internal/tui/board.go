package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding    = 1
	margin     = 1
	helpHeight = 1
)

type window struct {
	width, height int
}

type Board struct {
	Title    string
	quitting bool
	help     help.Model
	window   window
}

func NewBoard(title string) *Board {
	b := &Board{Title: title}
	b.help = help.New()
	return b
}
func (b *Board) setWindow(width, height int) {
	b.window.width = width
	b.window.height = height
}

func (b *Board) Init() tea.Cmd {
	return nil
}

func (b *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.setWindow(msg.Width, msg.Height)
		return b, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			if !b.quitting {
				b.quitting = true
			}
			return b, tea.Quit
		}
	}
	return b, nil
}

func (b *Board) View() string {
	if b.quitting {
		return ""
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		// main
		b.mainView(),
		// help
		b.helpView(),
	)
}

func (b *Board) helpView() string {
	b.help.Styles.ShortKey = lipgloss.NewStyle().Foreground(lipgloss.Color(white))
	b.help.Styles.ShortDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(gray))
	return lipgloss.NewStyle().
		MarginLeft(margin + padding).MarginBottom(margin).
		Render(b.help.View(keys))
}

func (b *Board) mainView() string {
	t := GetTheme()
	// board
	return lipgloss.NewStyle().
		Width(b.window.width-padding*2-margin*2).
		Height(b.window.height-padding*2-margin*2-helpHeight).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(t.PrimaryColor)).
		AlignHorizontal(lipgloss.Center).
		Padding(0, padding, 0, padding).
		Margin(0, margin, 0, margin).Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			// logo
			lipgloss.NewStyle().
				Foreground(lipgloss.Color(t.PrimaryColor)).Render(logo),
			// drawings
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				// black stone
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(white)).Render(blackStone),
				// arrow
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(t.PrimaryColor)).Padding(0, 3, 0, 3).
					Render(arrow),
				// white stone
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(white)).Render(whiteStone),
			),
		),
	)
}
