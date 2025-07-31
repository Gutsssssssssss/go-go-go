package page

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/client"
	gameClient "github.com/yanmoyy/go-go-go/internal/client/game"
	"github.com/yanmoyy/go-go-go/internal/game"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
	"github.com/yanmoyy/go-go-go/internal/tui/keys"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/theme"
	"github.com/yanmoyy/go-go-go/internal/tui/view"
	gameView "github.com/yanmoyy/go-go-go/internal/tui/view/game"
	"github.com/yanmoyy/go-go-go/internal/util"
)

const GamePage PageID = "game"

type gamePage struct {
	page
	help help.Model

	// game Client
	client *gameClient.GameClient
	done   chan struct{}

	selectedStoneID int
	status          gameView.ControlStatus
	degrees         int
	power           int

	// progress bar
	progressOn  progress.Model
	progressOff progress.Model

	// Animation
	animationData *game.AnimationData
	currentStep   int
	isAnimating   bool
}

type animationMsg struct {
	data *game.AnimationData
}

type gameStartedMsg struct{}

type gameOverMsg struct {
	winner string
}

type tickMsg struct{}

func NewGamePage() tea.Model {
	p := &gamePage{}
	p.help = help.New()
	return p
}

func (p *gamePage) Init() tea.Cmd {
	p.setProgresses()
	p.done = make(chan struct{})
	p.client = client.GetGameClient()
	err := p.client.StartListenConn(p.done)
	if err != nil {
		slog.Error("failed to start game", "err", err)
	}
	p.status = gameView.ControlSelectStone
	// listen for animation
	return p.ListenClient()
}

func (p *gamePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// update user stones with sorting by x coordinate
	keys := keys.GetGameKeys()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if p.isAnimating {
			return p, nil
		}
		switch {
		case key.Matches(msg, keys.Quit()):
			return p, p.onEscapePressed()
		case key.Matches(msg, keys.Up()):
			return p, p.onUpPressed()
		case key.Matches(msg, keys.Down()):
			return p, p.onDownPressed()
		case key.Matches(msg, keys.Left()):
			return p, p.onLeftPressed()
		case key.Matches(msg, keys.Right()):
			return p, p.onRightPressed()
		case key.Matches(msg, keys.Enter()):
			return p, p.onEnterPressed()
		}
	case animationMsg:
		var cmds []tea.Cmd
		if msg.data != nil {
			cmds = append(cmds, p.startAnimation(msg.data))
		}
		cmds = append(cmds, p.ListenClient())
		return p, tea.Batch(cmds...)
	case tickMsg:
		return p.updateAnimation()
	case gameStartedMsg:
		p.selectedStoneID = p.client.GetCurrentStone(0)
		return p, p.ListenClient()
	case gameOverMsg:
		// TODO: show winner dialog
		return p, nil
	}
	return p, nil
}

func (p *gamePage) View() string {
	if p.window.width == 0 || p.window.height == 0 {
		return ""
	}
	if p.client == nil {
		return ""
	}
	const (
		PADDING    = 1
		MARGIN     = 1
		HELPHEIGHT = 1
	)
	boardHeight := p.window.height - MARGIN - HELPHEIGHT
	boardWidth := 2 * boardHeight
	sideWidth := (p.window.width - boardWidth) / 2
	p.progressOn.Width = sideWidth - 6
	p.progressOff.Width = sideWidth - 6

	return layout.Column(
		layout.Row(
			view.Board(
				view.BoardProps{
					Width:       sideWidth,
					Height:      boardHeight,
					BorderColor: p.getTurnColor(),
				},
				layout.Column(
					view.Board(
						view.BoardProps{
							Width:       sideWidth - 2, // minus for border
							Height:      boardHeight/2 - 1,
							BorderColor: p.getStatusColor(gameView.ControlDirection),
						},
						layout.Column(
							fmt.Sprintf("Degrees (%d°)", p.degrees),
							layout.GapV(2),
							p.getProgress(p.status, gameView.ControlDirection).ViewAs(util.GetPercentage(int(p.degrees))),
						),
					),
					view.Board(
						view.BoardProps{
							Width:       sideWidth - 2, // minus for border
							Height:      boardHeight/2 - 1,
							BorderColor: p.getStatusColor(gameView.ControlCharging),
						},
						layout.Column(
							fmt.Sprintf("Power (%d)", p.power),
							layout.GapV(2),
							p.getProgress(p.status, gameView.ControlCharging).ViewAs(float64(p.power)/gameView.MaxPower),
						),
					),
				),
			),
			view.Board(
				view.BoardProps{
					Width:       boardWidth,
					Height:      boardHeight,
					BorderColor: p.getTurnColor(),
				},
				view.Board(
					view.BoardProps{
						Width:       boardWidth - 2,
						Height:      boardHeight - 2,
						BorderColor: p.getStatusColor(gameView.ControlSelectStone),
					},
					gameView.View(
						gameView.Props{
							Width:            boardWidth - 4,
							Height:           boardHeight - 4,
							GameData:         p.client.GetGameData(),
							AnimationsData:   p.animationData,
							CurAnimationStep: p.currentStep,
							ControlData: gameView.ControlData{
								Status:          p.status,
								SelectedStoneID: p.selectedStoneID,
								Degrees:         gameView.Degrees(p.degrees),
								IndicatorColor:  p.getTurnColor(),
							},
						},
					),
				),
			),
			view.Board(
				view.BoardProps{
					Width:       sideWidth - 2,
					Height:      boardHeight,
					BorderColor: p.getTurnColor(),
				},
				layout.Column(
					"",
				),
			),
		),
		view.Help(&p.help, view.HelpProps{
			KeyMap: keys.GetGameKeys(),
			Width:  p.window.width,
		}),
	)
}

func (p *gamePage) setProgresses() {
	p.progressOn = progress.New(
		progress.WithScaledGradient(string(color.GolangBlue), string(color.GolangLightBlue)),
	)
	p.progressOn.ShowPercentage = false
	p.progressOn.EmptyColor = string(color.Gray)

	p.progressOff = progress.New(
		progress.WithScaledGradient(string(color.Darkgray), string(color.Gray)),
	)
	p.progressOff.ShowPercentage = false
	p.progressOff.EmptyColor = string(color.Gray)
}

func (p *gamePage) getStatusColor(status gameView.ControlStatus) lipgloss.Color {
	if p.status == status {
		return p.getTurnColor()
	}
	return theme.GetTheme().DisabledColor
}

func (p *gamePage) getTurnColor() lipgloss.Color {
	if p.client != nil && p.client.IsPlayerTurn() {
		return theme.GetTheme().PrimaryColor
	}
	return theme.GetTheme().DisabledColor
}

func (p *gamePage) getProgress(s gameView.ControlStatus, current gameView.ControlStatus) progress.Model {
	if s == current {
		return p.progressOn
	}
	return p.progressOff
}

func (p *gamePage) shootStone() {
	err := p.client.ShootStone(
		p.selectedStoneID,
		p.degrees,
		p.power,
	)
	if err != nil {
		slog.Error("failed to shoot stone", "err", err)
		p.selectedStoneID = p.client.GetCurrentStone(p.selectedStoneID)
		p.status = gameView.ControlSelectStone // Fallback if no animation
	}
}

func (p *gamePage) ListenClient() tea.Cmd {
	return func() tea.Msg {
		select {
		case data, ok := <-p.client.AnimationCh:
			if !ok {
				return nil
			}
			return animationMsg{data: data}
		case state, ok := <-p.client.GameStateCh:
			if !ok {
				return nil
			}
			switch state.State {
			case gameClient.GameStateStart:
				return gameStartedMsg{}
			case gameClient.GameStateOver:
				return gameOverMsg{winner: state.Data["winner"]}
			}
			return nil
		}
	}
}

func (p *gamePage) startAnimation(data *game.AnimationData) tea.Cmd {
	p.animationData = data
	p.currentStep = 0
	p.isAnimating = true
	// Start animation with a tick
	return tea.Tick(time.Millisecond*10, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (p *gamePage) updateAnimation() (tea.Model, tea.Cmd) {
	if p.isAnimating && p.animationData != nil {
		p.currentStep++
		if p.currentStep >= p.animationData.MaxAnimationStep {
			// Animation complete
			p.isAnimating = false
			p.animationData = nil
			p.currentStep = 0
			p.status = gameView.ControlSelectStone
			return p, nil
		}
		// Continue animation
		return p, tea.Tick(time.Millisecond*10, func(time.Time) tea.Msg {
			return tickMsg{}
		})
	}
	return p, nil
}

// ########## Key Press Handler ##########

func (p *gamePage) onLeftPressed() tea.Cmd {
	if !p.client.IsPlayerTurn() {
		return nil
	}
	switch p.status {
	case gameView.ControlSelectStone:
		p.selectedStoneID = p.client.GetLeftStone(p.selectedStoneID)
	case gameView.ControlDirection:
		p.degrees = p.degrees - 15
		if p.degrees < gameView.MinDegrees {
			p.degrees = gameView.MaxDegrees - 15
		}
	}
	return nil
}

func (p *gamePage) onRightPressed() tea.Cmd {
	if !p.client.IsPlayerTurn() {
		return nil
	}
	switch p.status {
	case gameView.ControlSelectStone:
		p.selectedStoneID = p.client.GetRightStone(p.selectedStoneID)
	case gameView.ControlDirection:
		p.degrees = p.degrees + 15
		if p.degrees > gameView.MaxDegrees {
			p.degrees = gameView.MinDegrees + 15
		}
	}
	return nil
}

func (p *gamePage) onUpPressed() tea.Cmd {
	if p.status == gameView.ControlCharging {
		p.power += 1
		if p.power > gameView.MaxPower {
			p.power = gameView.MaxPower
		}
	}
	return nil
}

func (p *gamePage) onDownPressed() tea.Cmd {
	if p.status == gameView.ControlCharging {
		p.power -= 1
		if p.power < gameView.MinPower {
			p.power = gameView.MinPower
		}
	}
	return nil
}

func (p *gamePage) onEnterPressed() tea.Cmd {
	if !p.client.IsPlayerTurn() {
		return nil
	}
	switch p.status {
	case gameView.ControlSelectStone:
		p.status = gameView.ControlDirection
	case gameView.ControlDirection:
		p.status = gameView.ControlCharging
	case gameView.ControlCharging:
		p.shootStone()
	}
	return nil
}

func (p *gamePage) onEscapePressed() tea.Cmd {
	switch p.status {
	case gameView.ControlSelectStone:
		// TODO: quit dialog
		p.client.Close()
		return cmd(PagePopMsg{})
	case gameView.ControlDirection:
		p.status = gameView.ControlSelectStone
	case gameView.ControlCharging:
		p.status = gameView.ControlDirection
	}
	return nil
}
