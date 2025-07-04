package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	PADDING    = 1
	MARGIN     = 1
	HELPHEIGHT = 1
)

type window struct {
	width, height int
}

type buttonType int

const (
	btnStartGame buttonType = iota
	btnOptions
	btnHelp
	btnQuit
)

func (b buttonType) String() string {
	switch b {
	case btnStartGame:
		return "Start Game"
	case btnOptions:
		return "Options"
	case btnHelp:
		return "Help"
	case btnQuit:
		return "Quit"
	}
	return ""
}

type Board struct {
	Title    string
	quitting bool
	help     help.Model
	window   window

	buttons  []Button
	selected buttonType
}

type Button struct {
	text string
}

func (b Button) View(selected bool) string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Padding(0, PADDING, 0, PADDING).
		Margin(0, MARGIN, 0, MARGIN).
		Width(20).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(white)).
		Foreground(lipgloss.Color(white))
	if selected {
		style = style.BorderForeground(lipgloss.Color(t.PrimaryColor)).
			Foreground(lipgloss.Color(t.PrimaryColor))
	}
	return style.Render(b.text)
}

func NewBoard(title string) *Board {
	b := &Board{
		Title: title,
		buttons: []Button{
			{text: btnStartGame.String()},
			{text: btnOptions.String()},
			{text: btnHelp.String()},
			{text: btnQuit.String()},
		},
	}
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
		case key.Matches(msg, keys.Up):
			if b.selected == btnStartGame {
				return b, nil
			}
			b.selected--
		case key.Matches(msg, keys.Down):
			if b.selected == btnQuit {
				return b, nil
			}
			b.selected++
		case key.Matches(msg, keys.Enter):
			switch b.selected {
			case btnStartGame:
				return b, nil
			case btnOptions:
				return b, nil
			case btnHelp:
				return b, nil
			case btnQuit:
				b.quitting = true
				return b, tea.Quit
			}
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
		MarginLeft(MARGIN + PADDING).MarginBottom(MARGIN).
		Render(b.help.View(keys))
}

func (b *Board) mainView() string {
	t := GetTheme()
	return lipgloss.NewStyle().
		Width(b.window.width-PADDING*2-MARGIN*2).
		Height(b.window.height-PADDING*2-MARGIN*2-HELPHEIGHT).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(t.PrimaryColor)).
		AlignHorizontal(lipgloss.Center).
		Padding(0, PADDING, 0, PADDING).
		Margin(0, MARGIN, 0, MARGIN).
		Render(
			Column(
				lipgloss.NewStyle().Foreground(lipgloss.Color(t.PrimaryColor)).
					Render(logo),
				Row(
					lipgloss.NewStyle().Foreground(lipgloss.Color(white)).
						Render(blackStone),
					lipgloss.NewStyle().Foreground(lipgloss.Color(t.PrimaryColor)).Padding(0, 3, 0, 3).
						Render(arrow),
					lipgloss.NewStyle().Foreground(lipgloss.Color(white)).
						Render(whiteStone),
				),
				// Buttons
				lipgloss.NewStyle().
					Margin(3).
					Render(
						Column(
							b.buttonViews()...,
						),
					),
			),
		)
}

func (b *Board) buttonViews() []string {
	var views []string
	for i, btn := range b.buttons {
		views = append(views, btn.View(i == int(b.selected)))
	}
	return views
}

func Column(children ...string) string {
	return lipgloss.JoinVertical(lipgloss.Center, children...)
}

func Row(children ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center, children...)
}
