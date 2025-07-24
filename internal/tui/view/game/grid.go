package game

import "strings"

type grid [][]string

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

func (g grid) outOfBounds(x, y int) bool {
	return x < 0 || x >= len(g[0]) || y < 0 || y >= len(g)
}
