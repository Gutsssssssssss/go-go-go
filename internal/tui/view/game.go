package view

import (
	"strings"

	"github.com/yanmoyy/go-go-go/internal/game"
)

type GameProps struct {
	Width  int
	Height int
}

func Game(game *game.Game, props GameProps) string {
	// Get stones from the game
	stones := game.GetStones()
	gameW, gameH := game.GetSize()

	// Scaling factors for 100x100 board to TUI grid
	scaleW := float64(props.Width) / gameW
	scaleH := float64(props.Height) / gameH

	// Initialize a grid to represent the board (height rows x width columns)
	grid := createGrid(props.Width, props.Height)

	// // Map stones to the grid
	for _, stone := range stones {
		drawStone(grid, scaleW, scaleH, stone)
	}

	return grid.String()
}

type grid [][]string

func (g grid) String() string {
	var b strings.Builder
	for i := range g {
		for j := range g[i] {
			b.WriteString(g[i][j])
		}
		if i < len(g)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func createGrid(width, height int) grid {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = " " // Default to empty space
		}
	}
	return grid
}

func drawStone(grid grid, scaleW, scaleH float64, stone game.Stone) {
	x := stone.Position.X * scaleW
	y := stone.Position.Y * scaleH
	radiusW := stone.Radius * scaleW
	radiusH := stone.Radius * scaleH

	var symbol string
	if stone.StoneType == game.White {
		symbol = "●"
	} else {
		symbol = "◯"
	}
	drawCircle(grid, x, y, radiusW, radiusH, symbol)
}

// drawCircle draws a circle on the grid
func drawCircle(grid grid, posX, posY, radiusW, radiusH float64, symbol string) {
	if radiusW == 0 || radiusH == 0 {
		return
	}
	for y := int(posY - radiusH); y <= int(posY+radiusH); y++ {
		for x := int(posX - radiusW); x <= int(posX+radiusW); x++ {
			if outOfBounds(x, y, len(grid[0]), len(grid)) {
				continue
			}
			dx := (posX - float64(x)) / radiusW
			dy := (posY - float64(y)) / radiusH
			if dx*dx+dy*dy <= 1.0 {
				grid[y][x] = symbol
			}
		}
	}
}

func outOfBounds(x, y int, width, height int) bool {
	return x < 0 || x >= width || y < 0 || y >= height
}
