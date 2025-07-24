package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDrawCircle(t *testing.T) {
	// case 1: Midium circle
	grid := createGrid(5, 5)
	grid.drawCircle(2, 2, 1, 1, "r")
	t.Log(grid)
	require.Equal(t, "     \n  r  \n rrr \n  r  \n     ", grid.String())

	// case 2: Small circle
	grid = createGrid(3, 3)
	grid.drawCircle(1, 1, 0.5, 0.5, "r")
	t.Log(grid)
	require.Equal(t, "   \n r \n   ", grid.String())
}

func TestDrawArrow(t *testing.T) {
	// case 1: small triangle
	//   t
	grid := createGrid(3, 3)
	grid.drawTriangle(1, 1, 1, "t")
	t.Log(grid)
	require.Equal(t, "   \n t \n   ", grid.String())

	// case 2: medium triangle
	//   t
	//  ttt
	grid = createGrid(3, 3)
	grid.drawTriangle(1, 0, 2, "t")
	t.Log(grid)
	require.Equal(t, " t \nttt\n   ", grid.String())
}
