package page

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
)

const LobbyPage PageID = "lobby"

type lobbyPage struct {
	Page
	finding bool
}

func NewLobbyPage() tea.Model {
	p := &lobbyPage{}
	return p
}

func (p *lobbyPage) Init() tea.Cmd {
	p.finding = true
	return nil
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
	}
	return p, nil
}

func (p *lobbyPage) View() string {
	if p.finding {
		return lipgloss.NewStyle().
			Width(p.window.width).
			Height(p.window.height).
			Align(lipgloss.Center, lipgloss.Center).
			Render("Finding...")
	}
	return ""
}
