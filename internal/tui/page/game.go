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
	help            help.Model
	game            *game.Game
	selectedStoneID int
	status          gameView.ControlStatus
	degrees         gameView.Degrees
	power           gameView.Power
	progressOn      progress.Model
	progressOff     progress.Model

	// Animation
	animationData *game.StoneAnimationsData
	currentStep   int
	isAnimating   bool
}

func NewGamePage() tea.Model {
	p := &gamePage{}
	p.help = help.New()
	p.game = game.NewGame()
	p.setProgresses()
	return p
}
func (p *gamePage) Init() tea.Cmd {
	p.game.AddPlayer("player1")
	p.game.AddPlayer("player2")
	p.game.StartGame()
	p.selectedStoneID = 0
	p.status = gameView.ControlSelectStone
	return nil
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
			return p, cmd(PagePopMsg{})
		case key.Matches(msg, keys.Up()):
			if p.status == gameView.ControlCharging {
				p.power += gameView.Power(1)
				if p.power > gameView.MaxPower {
					p.power = gameView.MaxPower
				}
			}
		case key.Matches(msg, keys.Down()):
			if p.status == gameView.ControlCharging {
				p.power -= gameView.Power(1)
				if p.power < gameView.MinPower {
					p.power = gameView.MinPower
				}
			}
		case key.Matches(msg, keys.Left()):
			switch p.status {
			case gameView.ControlSelectStone:
				p.selectedStoneID = p.game.GetLeftStone(0, p.selectedStoneID)
			case gameView.ControlDirection:
				p.degrees = p.degrees - 15
				if p.degrees < gameView.MinDegrees {
					p.degrees = gameView.MaxDegrees - 15
				}
			}
		case key.Matches(msg, keys.Right()):
			switch p.status {
			case gameView.ControlSelectStone:
				p.selectedStoneID = p.game.GetRightStone(0, p.selectedStoneID)
			case gameView.ControlDirection:
				p.degrees = p.degrees + 15
				if p.degrees > gameView.MaxDegrees {
					p.degrees = gameView.MinDegrees + 15
				}
			}
		case key.Matches(msg, keys.Enter()):
			switch p.status {
			case gameView.ControlSelectStone:
				p.status = gameView.ControlDirection
			case gameView.ControlDirection:
				p.status = gameView.ControlCharging
			case gameView.ControlCharging:
				// TODO: shoot stone
				velocity := game.ConvertToVelocity(float64(p.degrees), float64(p.power))
				evt := p.game.ShootStone(
					game.ShootData{
						PlayerID: 0,
						StoneID:  p.selectedStoneID,
						Velocity: velocity,
					},
				)
				return p.startAnimation(evt)
			}
		}
	case tickMsg:
		return p.updateAnimation()
	}
	return p, nil
}

type tickMsg struct{}

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
	sideWidth := (p.window.width - boardWidth) / 2
	p.progressOn.Width = sideWidth - 6
	p.progressOff.Width = sideWidth - 6

	return layout.Column(
		layout.Row(
			view.Board(
				view.BoardProps{
					Width:       sideWidth,
					Height:      boardHeight,
					BorderColor: t.PrimaryColor,
				},
				layout.Column(
					view.Board(
						view.BoardProps{
							Width:       sideWidth - 2, // minus for border
							Height:      boardHeight/2 - 1,
							BorderColor: getColor(p.status, gameView.ControlDirection),
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
							BorderColor: getColor(p.status, gameView.ControlCharging),
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
					BorderColor: t.PrimaryColor,
				},
				view.Board(
					view.BoardProps{
						Width:       boardWidth - 2,
						Height:      boardHeight - 2,
						BorderColor: getColor(p.status, gameView.ControlSelectStone),
					},
					gameView.View(
						p.game,
						gameView.Props{
							Width:            boardWidth - 4,
							Height:           boardHeight - 4,
							AnimationsData:   p.animationData,
							CurAnimationStep: p.currentStep,
							ControlData: gameView.ControlData{
								Status:          p.status,
								SelectedStoneID: p.selectedStoneID,
								Degrees:         p.degrees,
								Power:           gameView.Power(0),
							},
						}),
				),
			),
			view.Board(
				view.BoardProps{
					Width:       sideWidth,
					Height:      boardHeight,
					BorderColor: t.PrimaryColor,
				},
				layout.Column(
					"right",
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

func getColor(s gameView.ControlStatus, current gameView.ControlStatus) lipgloss.Color {
	if s == current {
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

func (p *gamePage) startAnimation(evt game.Event) (tea.Model, tea.Cmd) {
	if data, ok := evt.Data.(game.StoneAnimationsData); ok {
		p.animationData = &data
		p.currentStep = 0
		p.isAnimating = true
		// Start animation with a tick
		return p, tea.Tick(time.Millisecond*10, func(time.Time) tea.Msg {
			return tickMsg{}
		})
	}
	p.selectedStoneID = p.game.GetCurrentStone(0, p.selectedStoneID)
	p.status = gameView.ControlSelectStone // Fallback if no animation
	return p, nil
}

func (p *gamePage) updateAnimation() (tea.Model, tea.Cmd) {
	if p.isAnimating && p.animationData != nil {
		p.currentStep++
		if p.currentStep >= p.animationData.MaxStep {
			// Animation complete
			p.isAnimating = false
			p.animationData = nil
			p.currentStep = 0
			p.status = gameView.ControlSelectStone
			return p, nil
		}
		slog.Info("Animation", "currentStep", p.currentStep, "maxStep", p.animationData.MaxStep)
		// Continue animation
		return p, tea.Tick(time.Millisecond*10, func(time.Time) tea.Msg {
			return tickMsg{}
		})
	}
	return p, nil
}
