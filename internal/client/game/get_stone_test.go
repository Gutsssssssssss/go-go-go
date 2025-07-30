package game

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yanmoyy/go-go-go/internal/game"
)

func TestGetFilteredStones(t *testing.T) {
	stones := []game.Stone{
		{ID: 0, Position: game.Vector2{X: 0, Y: 0}, StoneType: game.White},
		{ID: 1, Position: game.Vector2{X: 1, Y: 0}, StoneType: game.White},
		{ID: 2, Position: game.Vector2{X: 2, Y: 0}, StoneType: game.Black},
		{ID: 3, Position: game.Vector2{X: 3, Y: 0}, StoneType: game.White, IsOut: true},
	}
	res := getFilteredStones(stones, game.White)
	require.Equal(t, []game.Stone{
		{ID: 0, Position: game.Vector2{X: 0, Y: 0}, StoneType: game.White},
		{ID: 1, Position: game.Vector2{X: 1, Y: 0}, StoneType: game.White},
	}, res)

	res = getFilteredStones(stones, game.Black)
	require.Equal(t, []game.Stone{
		{ID: 2, Position: game.Vector2{X: 2, Y: 0}, StoneType: game.Black},
	}, res)
}

func TestGetNextStone(t *testing.T) {
	// Test 1: getNextStone
	c := NewGameClient(nil)
	c.gameData.Stones = []game.Stone{
		{ID: 0, Position: game.Vector2{X: 1, Y: 0}, StoneType: game.White},
		{ID: 1, Position: game.Vector2{X: 0, Y: 0}, StoneType: game.White},
		{ID: 2, Position: game.Vector2{X: 2, Y: 0}, StoneType: game.White},
	}

	// order: 1 0 2
	// case 1: normal
	nxtID, err := c.getNextStone(0, -1)
	require.NoError(t, err)
	require.Equal(t, 1, nxtID)

	// case 2: not found
	nxtID, err = c.getNextStone(3, -1)
	require.NoError(t, err)
	require.Equal(t, 2, nxtID)

	// Test 2: GetLeftStone
	nxtID = c.GetLeftStone(2)
	require.Equal(t, 0, nxtID)
	nxtID = c.GetLeftStone(1)
	require.Equal(t, 2, nxtID)

	// Test 3: GetRightStone
	nxtID = c.GetRightStone(1)
	require.Equal(t, 0, nxtID)
	nxtID = c.GetRightStone(2)
	require.Equal(t, 1, nxtID)
}
