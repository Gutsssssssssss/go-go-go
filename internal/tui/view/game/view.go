package game

import (
	gameClient "github.com/yanmoyy/go-go-go/internal/client/game"
	"github.com/yanmoyy/go-go-go/internal/game"
)

type scale struct {
	width  float64
	height float64
}

type Props struct {
	Width       int
	Height      int
	GameData    *gameClient.GameData
	ControlData ControlData

	// Animation
	AnimationsData   *game.StoneAnimationsData
	CurAnimationStep int
}

func View(props Props) string {
	if props.GameData == nil || props.GameData.Size.Width == 0 || props.GameData.Size.Height == 0 || len(props.GameData.Stones) == 0 {
		return ""
	}
	// Get the size of the game board
	gameW, gameH := props.GameData.Size.Width, props.GameData.Size.Height

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

	// All stones
	grid.drawStones(props.GameData.Stones, scale, props.ControlData)

	// Indicator
	selectedStone := props.GameData.Stones[props.ControlData.SelectedStoneID]
	grid.drawIndicator(selectedStone, scale, props.ControlData)

	return grid.String()
}
