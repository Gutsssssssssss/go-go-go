package page

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
	"github.com/yanmoyy/go-go-go/internal/tui/view"
)

const LobbyPage PageID = "lobby"

type lobbyPage struct {
	page
	help           help.Model
	spinner        spinner.Model
	choices        []lobbyChoice
	selected       int
	finding        bool
	seeMoreOptions bool
}

type lobbyChoice struct {
	name    string
	enabled bool
}

func NewLobbyPage() tea.Model {
	p := &lobbyPage{
		choices: []lobbyChoice{
			{"Play with Player", false},
			{"Play with Bot", true},
			{"Retry Finding", false},
			{"Quit", true},
		},
		selected: 0,
	}
	p.help = help.New()
	return p
}

func (p *lobbyPage) Init() tea.Cmd {
	t := theme.GetTheme()
	p.finding = true
	connectionTimer := time.NewTimer(time.Second * 30)
	seeMoreTimer := time.NewTimer(time.Second * 3)
	go func() {
		select {
		case <-connectionTimer.C:
			p.finding = false
		case <-seeMoreTimer.C:
			p.seeMoreOptions = true
		}
	}()
	p.spinner = t.WaitingSpinner
	return p.spinner.Tick
}

func (p *lobbyPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keys := keys.GetBasicKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, func() tea.Msg {
				return PagePopMsg{}
			}
		case key.Matches(msg, keys.Up()):
			if p.selected > 0 {
				p.selected--
			}
		case key.Matches(msg, keys.Down()):
			if p.selected < len(p.choices)-1 {
				p.selected++
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
	const (
		PADDING    = 1
		MARGIN     = 1
		HELPHEIGHT = 1
	)
	t := theme.GetTheme()
	boardHeight := p.window.height - MARGIN*2 - HELPHEIGHT
	return layout.Column(
		view.Board(
			view.BoardProps{
				Width:       p.window.width - 2*PADDING,
				Height:      boardHeight,
				BorderColor: t.PrimaryColor,
				Padding:     PADDING,
				Margin:      MARGIN,
			},
			layout.FlexVertical(
				boardHeight-2*PADDING-2*MARGIN,
				layout.Expanded(""),
				// spinner
				layout.Fixed(
					layout.Column(
						"Finding another Player",
						layout.GapV(2),
						p.spinner.View(),
					),
				),
				layout.Expanded(""),
				// buttons
				layout.Fixed(
					layout.Column(
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
			),
		),
		view.Help(&p.help, keys.GetBasicKeys()),
	)
}

func (p *lobbyPage) choiceView(choice lobbyChoice, selected bool) string {
	t := theme.GetTheme()
	c := t.DefaultColor
	if !choice.enabled {
		c = t.DisabledColor
	}
	if selected {
		c = t.PrimaryColor
	}
	return view.Button(view.ButtonProps{
		Text:  choice.name,
		Color: c,
	})
}
