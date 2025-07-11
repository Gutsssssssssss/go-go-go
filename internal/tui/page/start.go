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
	"github.com/yanmoyy/go-go-go/internal/tui/view"
)

const StartPage PageID = "Start"

type startPage struct {
	page
	help     help.Model
	quitting bool
	choices  []string
	selected int
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
	keys := keys.GetBasicKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			if !p.quitting {
				p.quitting = true
			}
			return p, tea.Quit
		case key.Matches(msg, keys.Up()):
			if p.selected > 0 {
				p.selected--
			}
		case key.Matches(msg, keys.Down()):
			if p.selected < len(p.choices)-1 {
				p.selected++
			}
		case key.Matches(msg, keys.Enter()):
			switch p.selected {
			case 0: // Start Game
				return p, func() tea.Msg {
					return PagePushMsg{ID: LobbyPage}
				}
			case 1: // Options
				return p, nil
			case 2: // Help
				return p, nil
			case 3: // Quit
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
	return layout.Column(
		// main
		p.mainView(),
		// help
		view.Help(&p.help, keys.GetBasicKeys()),
	)
}
func (p *startPage) mainView() string {
	const (
		PADDING    = 1
		MARGIN     = 1
		HELPHEIGHT = 1
	)
	t := theme.GetTheme()
	boardHeight := p.window.height - MARGIN*2 - HELPHEIGHT
	return view.Board(
		view.BoardProps{
			Width:       p.window.width - 2*PADDING,
			Height:      boardHeight,
			BorderColor: t.PrimaryColor,
			Padding:     PADDING,
			Margin:      MARGIN,
		},
		layout.FlexVertical(
			boardHeight-2*PADDING-2*MARGIN,
			layout.Fixed(
				layout.Column(
					lipgloss.NewStyle().Foreground(t.PrimaryColor).
						Render(drawing.Logo),
					layout.Row(
						lipgloss.NewStyle().Foreground(color.White).
							Render(drawing.BlackStone),
						lipgloss.NewStyle().Foreground(t.PrimaryColor).
							Padding(0, 3, 0, 3).
							Render(drawing.Arrow),
						lipgloss.NewStyle().Foreground(color.White).
							Render(drawing.WhiteStone),
					),
				),
			),
			layout.Expanded(""),
			layout.Fixed(
				// Buttons
				lipgloss.NewStyle().
					Margin(2).
					Render(
						layout.Column(
							p.choiceView(p.choices[0], p.selected == 0),
							p.choiceView(p.choices[1], p.selected == 1),
							p.choiceView(p.choices[2], p.selected == 2),
							p.choiceView(p.choices[3], p.selected == 3),
						),
					),
			),
		),
	)
}

func (p *startPage) choiceView(text string, selected bool) string {
	t := theme.GetTheme()
	c := t.PrimaryColor
	if !selected {
		c = t.DefaultColor
	}
	return view.Button(view.ButtonProps{
		Text:  text,
		Color: c,
	})
}
