package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDrawCircle(t *testing.T) {
	// case 1: Midium circle
	grid := createGrid(5, 5)
	grid.drawCircle(2, 2, 1, 1, "r")
	require.Equal(t, "     \n  r  \n rrr \n  r  \n     ", grid.String())

	// case 2: Small circle
	grid = createGrid(3, 3)
	grid.drawCircle(1, 1, 0.5, 0.5, "r")
	require.Equal(t, "   \n r \n   ", grid.String())
}

func TestDrawArrow(t *testing.T) {
	// case 1: small triangle
	//   t
	grid := createGrid(3, 3)
	grid.drawTriangle(1, 1, 1, "t")
	require.Equal(t, "   \n t \n   ", grid.String())

	// case 2: medium triangle
	//   t
	//  ttt
	grid = createGrid(3, 3)
	grid.drawTriangle(1, 0, 2, "t")
	require.Equal(t, " t \nttt\n   ", grid.String())
}

func TestDrawDirection(t *testing.T) {
	// case 1: small size
	// ↖ ↑ ↗
	// ← 0 →
	// ↙ ↓ ↘
	grid := createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 0, "d")
	require.Equal(t, "   \n  d\n   ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 45, "d")
	require.Equal(t, "   \n   \n  d", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 90, "d")
	require.Equal(t, "   \n   \n d ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 135, "d")
	require.Equal(t, "   \n   \nd  ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 180, "d")
	require.Equal(t, "   \nd  \n   ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 225, "d")
	require.Equal(t, "d  \n   \n   ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 270, "d")
	require.Equal(t, " d \n   \n   ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 315, "d")
	require.Equal(t, "  d\n   \n   ", grid.String())

	grid = createGrid(3, 3)
	grid.drawDirection(1, 1, 0.5, 0.5, 360, "d")
	require.Equal(t, "   \n  d\n   ", grid.String())
}
