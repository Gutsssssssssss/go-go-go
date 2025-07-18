package page

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	"github.com/yanmoyy/go-go-go/internal/client"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
	"github.com/yanmoyy/go-go-go/internal/tui/view"
)

const LobbyPage PageID = "lobby"

type lobbyStatus int

const (
	lobbyInitial lobbyStatus = iota
	lobbyWaiting
	lobbyNotFoundPlayer
	lobbyEnteringGame
	lobbyFailedGetID
	lobbyConnectionErr
)

type lobbyPage struct {
	page
	help     help.Model
	spinner  spinner.Model
	choices  []string
	selected int
	client   *client.Client
	data     struct {
		id         uuid.UUID
		opponentID uuid.UUID
	}
	status         lobbyStatus
	findingSeconds int
	message        string
	messageTimer   *time.Timer
}

func NewLobbyPage() tea.Model {
	p := &lobbyPage{
		choices: []string{
			"Retry Finding",
			"Play with Bot",
			"Quit",
		},
		selected: 0,
	}
	p.help = help.New()
	p.client = client.DefaultClient()
	return p
}

func (p *lobbyPage) fetchUserID() {
	id, err := p.client.GetID()
	if err != nil {
		p.status = lobbyFailedGetID
		return
	}
	p.data.id = id
}

func (p *lobbyPage) startWaiting() {
	if p.status == lobbyWaiting || p.data.id == uuid.Nil {
		return
	}
	const timeout = time.Second * 10

	p.status = lobbyWaiting
	p.findingSeconds = int(timeout.Seconds())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	go func() {
		defer cancel()
		opponentID, err := p.client.StartWaiting(p.data.id, ctx)
		if err != nil {
			slog.Error("Connection Error", "err", err)
			p.status = lobbyConnectionErr
			return
		}
		if opponentID == uuid.Nil {
			slog.Info("Not Found Player")
			p.status = lobbyNotFoundPlayer
			return
		}
		slog.Info("Found Player", "id", opponentID)
		p.data.opponentID = opponentID
		p.status = lobbyEnteringGame
	}()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.findingSeconds--
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (p *lobbyPage) Init() tea.Cmd {
	p.selected = 1
	if p.status != lobbyInitial {
		return p.spinner.Tick
	}
	p.status = lobbyInitial
	if p.data.id == uuid.Nil {
		p.fetchUserID()
	}
	t := theme.GetTheme()
	p.startWaiting()
	p.spinner = t.WaitingSpinner
	return p.spinner.Tick
}

func (p *lobbyPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	const (
		retry int = iota
		withBot
		quit
	)
	keys := keys.GetBasicKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, cmd(PagePopMsg{})
		case key.Matches(msg, keys.Up()):
			if p.selected > 0 {
				p.selected--
			}
		case key.Matches(msg, keys.Down()):
			if p.selected < len(p.choices)-1 {
				p.selected++
			}
		case key.Matches(msg, keys.Enter()):
			if !p.isChoiceEnabled(p.selected) {
				p.showMessage("Please select a valid option")
				break
			}
			switch p.selected {
			case retry:
				p.startWaiting()
			case withBot:
				p.showMessage("Play with Bot")
			case quit:
				return p, cmd(PagePopMsg{})
			}
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		p.spinner, cmd = p.spinner.Update(msg)
		return p, cmd
	}

	if p.status == lobbyEnteringGame {
		return p, tea.Batch(
			cmd(PagePushMsg{ID: GamePage}),
		)
	}
	return p, nil
}

func (p *lobbyPage) showMessage(msg string) {
	p.message = msg
	if p.messageTimer == nil {
		p.messageTimer = time.AfterFunc(time.Second*1, func() {
			p.message = ""
			p.messageTimer = nil
		})
	} else {
		p.messageTimer.Reset(time.Second * 1)
	}
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
					p.statusView(),
				),
				layout.Expanded(""),
				layout.Fixed(
					p.messageView(),
				),
				// buttons
				layout.Fixed(
					layout.Column(
						lipgloss.NewStyle().
							Margin(2).
							Render(
								layout.Column(
									p.choiceView(0),
									p.choiceView(1),
									p.choiceView(2),
								),
							),
					),
				),
			),
		),
		view.Help(&p.help, keys.GetBasicKeys()),
	)
}
func (p *lobbyPage) isChoiceEnabled(idx int) bool {
	const (
		retry int = iota
		withBot
		quit
	)
	switch idx {
	case retry:
		return p.status == lobbyNotFoundPlayer || p.status == lobbyConnectionErr
	case withBot:
		return true
	case quit:
		return true
	default:
		return true
	}
}

func (p *lobbyPage) choiceView(idx int) string {
	t := theme.GetTheme()
	textColor := t.DefaultColor
	borderColor := t.DefaultColor
	text := p.choices[idx]
	enabled := p.isChoiceEnabled(idx)
	selected := p.selected == idx
	if !enabled {
		textColor = t.DisabledColor
		borderColor = t.DisabledColor
	}
	if selected {
		borderColor = t.PrimaryColor
		if enabled {
			textColor = t.PrimaryColor
		}
	}
	return view.Button(view.ButtonProps{
		Text:        text,
		TextColor:   textColor,
		BorderColor: borderColor,
	})
}

func (p *lobbyPage) statusView() string {
	switch p.status {
	case lobbyWaiting:
		return layout.Column(
			fmt.Sprintf("Finding another Player (%ds)", p.findingSeconds),
			layout.GapV(2),
			p.spinner.View(),
		)
	case lobbyEnteringGame:
		return "Found Player! Entering Game..."
	case lobbyNotFoundPlayer:
		return "Not Found Player"
	case lobbyFailedGetID:
		return "Connection Error (start again)"
	case lobbyConnectionErr:
		return "Connection Error"
	}
	return ""
}

func (p *lobbyPage) messageView() string {
	return lipgloss.NewStyle().
		Margin(0, 0, 2, 0).
		Render(p.message)
}
