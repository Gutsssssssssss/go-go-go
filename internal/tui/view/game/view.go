package game

import (
	"github.com/yanmoyy/go-go-go/internal/game"
)

type scale struct {
	width  float64
	height float64
}

type Props struct {
	Width            int
	Height           int
	ControlData      ControlData
	AnimationsData   *game.StoneAnimationsData
	CurAnimationStep int
}

func View(g *game.Game, props Props) string {
	// Get the size of the game board
	gameW, gameH := g.GetSize()

	// Scaling factors for 100x100 board to TUI grid
	scale := scale{width: float64(props.Width) / gameW, height: float64(props.Height) / gameH}

	// Initialize a grid to represent the board (height rows x width columns)
	grid := createGrid(props.Width, props.Height)
	if props.AnimationsData != nil {
		initialStones := props.AnimationsData.InitialStones
		grid.drawStones(initialStones, scale, props.ControlData)
		for _, anim := range props.AnimationsData.Animations {
			if anim.StartStep > props.CurAnimationStep {
				continue
			}
			grid.drawAnimation(anim, props.CurAnimationStep, scale, initialStones[anim.StoneID])
		}
		return grid.String()
	}

	stones := g.GetStones()
	grid.drawStones(stones, scale, props.ControlData)

	selectedStone := stones[props.ControlData.SelectedStoneID]
	grid.drawIndicator(selectedStone, scale, props.ControlData)

	return grid.String()
}
