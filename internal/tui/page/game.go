package page

import (
	"log/slog"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/game"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
	"github.com/yanmoyy/go-go-go/internal/tui/view"
)

const GamePage PageID = "game"

type gamePage struct {
	page
	help help.Model
	game *game.Game
}

func NewGamePage() tea.Model {
	p := &gamePage{}
	p.help = help.New()
	p.game = game.NewGame()
	p.game.StartGame()
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
	if p.window.width == 0 || p.window.height == 0 {
		return ""
	}
	t := theme.GetTheme()
	const (
		PADDING    = 1
		MARGIN     = 1
		HELPHEIGHT = 1
	)
	boardHeight := p.window.height - MARGIN - HELPHEIGHT
	boardWidth := 2 * boardHeight
	return layout.Column(
		view.Board(
			view.BoardProps{
				Width:       2 * boardHeight,
				Height:      boardHeight,
				BorderColor: t.PrimaryColor,
			},
			view.Game(
				p.game,
				view.GameProps{
					Width:  boardWidth - 2,
					Height: boardHeight - 2,
				}),
		),
		view.Help(&p.help, view.HelpProps{
			KeyMap: keys.GetGameKeys(),
			Width:  p.window.width,
		}),
	)
}
