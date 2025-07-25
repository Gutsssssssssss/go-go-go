package game

import (
	"github.com/yanmoyy/go-go-go/internal/game"
)

type scale struct {
	width  float64
	height float64
}

type Props struct {
	Width       int
	Height      int
	ControlData ControlData
}

func View(game *game.Game, props Props) string {
	// Get stones from the game
	stones := game.GetStones()
	gameW, gameH := game.GetSize()

	// Scaling factors for 100x100 board to TUI grid
	scale := scale{width: float64(props.Width) / gameW, height: float64(props.Height) / gameH}

	// Initialize a grid to represent the board (height rows x width columns)
	grid := createGrid(props.Width, props.Height)

	// Map stones to the grid
	for _, stone := range stones {
		if !stone.IsOut {
			grid.drawStone(stone, scale, props.ControlData)
		}
	}
	selectedStone := stones[props.ControlData.SelectedStoneID]
	grid.drawIndicator(selectedStone, scale, props.ControlData)

	return grid.String()
}
