package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddPlayer(t *testing.T) {
	g := NewGame()
	needStart, err := g.AddPlayer("player1")
	require.NoError(t, err)
	require.False(t, needStart)
	require.Equal(t, 0, g.idMap["player1"])
	require.Equal(t, newPlayer(0, white), g.players[0])

	needStart, err = g.AddPlayer("player2")
	require.NoError(t, err)
	require.True(t, needStart)
	require.Equal(t, 1, g.idMap["player2"])
	require.Equal(t, newPlayer(1, black), g.players[1])
}
