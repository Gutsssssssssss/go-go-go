package page

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
)

const LobbyPage PageID = "lobby"

type lobbyPage struct {
	Page
	spinner spinner.Model
	finding bool
}

func NewLobbyPage() tea.Model {
	p := &lobbyPage{}
	return p
}

func (p *lobbyPage) Init() tea.Cmd {
	t := theme.GetTheme()
	p.finding = true
	p.spinner = t.WaitingSpinner
	return p.spinner.Tick
}

func (p *lobbyPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, func() tea.Msg {
				return PagePopMsg{}
			}
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		p.spinner, cmd = p.spinner.Update(msg)
		return p, cmd
	}
	return p, nil
}

func (p *lobbyPage) View() string {
	if p.finding {
		return lipgloss.NewStyle().
			Width(p.window.width).
			Height(p.window.height).
			Align(lipgloss.Center, lipgloss.Center).
			Render(
				layout.Column(
					lipgloss.NewStyle().Render("Finding a another player..."),
					layout.GapV(2),
					p.spinner.View(),
				),
			)
	}
	return ""
}
