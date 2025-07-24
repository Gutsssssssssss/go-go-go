package page

import (
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
	help            help.Model
	game            *game.Game
	selectedStoneID int
}

func NewGamePage() tea.Model {
	p := &gamePage{}
	p.help = help.New()
	p.game = game.NewGame()
	p.game.AddPlayer("player1")
	p.game.AddPlayer("player2")
	p.game.StartGame()
	return p
}

func (p *gamePage) Init() tea.Cmd {
	p.selectedStoneID = 10
	return nil
}

func (p *gamePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// update user stones with sorting by x coordinate
	keys := keys.GetGameKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, cmd(PagePopMsg{})
		case key.Matches(msg, keys.Up()):
		case key.Matches(msg, keys.Down()):
		case key.Matches(msg, keys.Left()):
			p.selectedStoneID = p.game.GetLeftStone(p.selectedStoneID)
		case key.Matches(msg, keys.Right()):
			p.selectedStoneID = p.game.GetRightStone(p.selectedStoneID)
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
					Width:           boardWidth - 2,
					Height:          boardHeight - 2,
					IndicatorColor:  t.PrimaryColor,
					SelectedStoneID: p.selectedStoneID,
				}),
		),
		view.Help(&p.help, view.HelpProps{
			KeyMap: keys.GetGameKeys(),
			Width:  p.window.width,
		}),
	)
}
