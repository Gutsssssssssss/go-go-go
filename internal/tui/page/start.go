package page

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
	"github.com/yanmoyy/go-go-go/internal/tui/drawing"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
)

const StartPage PageID = "Start"

const (
	choiceStart int = iota
	choiceOptions
	choiceHelp
	choiceQuit
)

type startPage struct {
	Page
	quitting bool
	choices  []string
	selected int
}

type startKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
}

func (k startKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.Enter,
		k.Quit,
	}
}

func (k startKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

var startKeys = startKeyMap{
	Up:    keys.Up(),
	Down:  keys.Down(),
	Enter: keys.Enter(),
	Quit:  keys.Quit(),
}

func NewStartPage() tea.Model {
	p := &startPage{
		choices: []string{
			"Start Game",
			"Options",
			"Help",
			"Quit",
		},
	}
	p.help = help.New()
	return p
}

func (p *startPage) Init() tea.Cmd {
	return nil
}

func (p *startPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, startKeys.Quit):
			if !p.quitting {
				p.quitting = true
			}
			return p, tea.Quit
		case key.Matches(msg, startKeys.Up):
			if p.selected > 0 {
				p.selected--
			}
		case key.Matches(msg, startKeys.Down):
			if p.selected < len(startKeys.ShortHelp())-1 {
				p.selected++
			}
		case key.Matches(msg, startKeys.Enter):
			switch p.selected {
			case choiceStart:
				return p, func() tea.Msg {
					return PagePushMsg{ID: LobbyPage}
				}
			case choiceOptions:
				return p, nil
			case choiceHelp:
				return p, nil
			case choiceQuit:
				p.quitting = true
				return p, tea.Quit
			}
		}
	}
	return p, nil
}

func (p *startPage) View() string {
	if p.quitting {
		return ""
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		// main
		p.mainView(),
		// help
		p.helpView(),
	)
}

func (p *startPage) helpView() string {
	p.help.Styles.ShortKey = lipgloss.NewStyle().Foreground(color.White)
	p.help.Styles.ShortDesc = lipgloss.NewStyle().Foreground(color.Gray)
	return lipgloss.NewStyle().
		MarginLeft(2).MarginBottom(1).
		Render(p.help.View(startKeys))
}

func (p *startPage) mainView() string {
	const (
		PADDING    = 1
		MARGIN     = 1
		HELPHEIGHT = 1
	)
	t := theme.GetTheme()
	return lipgloss.NewStyle().
		Width(p.window.width-2*PADDING).
		Height(p.window.height-PADDING*2-MARGIN*2-HELPHEIGHT).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color(t.PrimaryColor)).
		AlignHorizontal(lipgloss.Center).
		Padding(0, PADDING, 0, PADDING).
		Margin(0, MARGIN, 0, MARGIN).
		Render(
			layout.Column(
				lipgloss.NewStyle().Foreground(t.PrimaryColor).
					Render(drawing.Logo),
				layout.Row(
					lipgloss.NewStyle().Foreground(color.White).
						Render(drawing.BlackStone),
					lipgloss.NewStyle().Foreground(t.PrimaryColor).Padding(0, 3, 0, 3).
						Render(drawing.Arrow),
					lipgloss.NewStyle().Foreground(color.White).
						Render(drawing.WhiteStone),
				),
				// Buttons
				lipgloss.NewStyle().
					Margin(3).
					Render(
						layout.Column(
							choiceView(p.choices[0], p.selected == 0),
							choiceView(p.choices[1], p.selected == 1),
							choiceView(p.choices[2], p.selected == 2),
							choiceView(p.choices[3], p.selected == 3),
						),
					),
			),
		)
}

func choiceView(text string, enabled bool) string {
	t := theme.GetTheme()
	style := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Padding(0, 1).
		Margin(0, 1).
		Width(20).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(color.White).
		Foreground(color.White)
	if enabled {
		style = style.BorderForeground(t.PrimaryColor).
			Foreground(t.PrimaryColor)
	}
	return style.Render(text)
}
