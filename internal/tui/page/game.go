package page

import (
	"log/slog"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/client"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
)

const GamePage PageID = "game"

type gamePage struct {
	page
	help   help.Model
	client *client.Client
}

func NewGamePage() tea.Model {
	p := &gamePage{}
	p.client = client.DefaultClient()
	return p
}

func (p *gamePage) Init() tea.Cmd {
	return nil
}

func (p *gamePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keys := keys.GetGameKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, cmd(PagePopMsg{})
		case key.Matches(msg, keys.Up()):
			slog.Info("Up")
		case key.Matches(msg, keys.Down()):
			slog.Info("Down")
		case key.Matches(msg, keys.Left()):
			slog.Info("Left")
		case key.Matches(msg, keys.Right()):
			slog.Info("Right")
		}
	}
	return p, nil
}

func (p *gamePage) View() string {
	return "Game Page"
}
