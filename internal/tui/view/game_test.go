package view

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDrawCircle(t *testing.T) {
	// case 1: Midium circle
	grid := createGrid(5, 5)
	drawCircle(grid, 2, 2, 1, 1, "r")
	t.Log(grid)
	require.Equal(t, "     \n  r  \n rrr \n  r  \n     ", grid.String())

	// case 2: Small circle
	grid = createGrid(3, 3)
	drawCircle(grid, 1, 1, 0.5, 0.5, "r")
	t.Log(grid)
	require.Equal(t, "   \n r \n   ", grid.String())
}

// r:3 circle (1,1)
//   010
//   111
//   010

// 1 circle
//   000
//   010
//   000
