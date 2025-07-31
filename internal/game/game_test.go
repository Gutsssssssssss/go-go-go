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
	require.Equal(t, playerID(0), g.idMap["player1"])
	require.Equal(t, Player{ID: 0, StoneType: White}, g.players[0])

	needStart, err = g.AddPlayer("player2")
	require.NoError(t, err)
	require.True(t, needStart)
	require.Equal(t, playerID(1), g.idMap["player2"])
	require.Equal(t, Player{ID: 1, StoneType: Black}, g.players[1])
}

func TestSimulateCollision(t *testing.T) {
	// Test 1: collision only one
	pos1, v1 := Vector2{X: 0, Y: 0}, Vector2{X: 1, Y: 0}
	pos2, v2 := Vector2{X: 1, Y: 0}, Vector2{X: 0, Y: 0}
	stones := []Stone{
		{ID: 0, Position: pos1, Radius: 1},
		{ID: 1, Position: pos2, Radius: 1},
	}
	movings := []moving{
		{id: stones[0].ID, startPos: pos1, velocity: v1},
	}
	v1 = applyFriction(v1, friction*0.1)
	v1, v2 = computeCollisionVelocities(v1, v2, pos1, pos2)
	animations := []StonePath{}

	afterPos := Vector2{X: 0.1, Y: 0}

	movings, animations = simulateCollision(movings, stones, animations, 0.1)
	require.Equal(t, 2, len(movings))
	require.Equal(t, 1, len(animations))
	require.Equal(t, []moving{
		{id: 1, startPos: pos2, startStep: 1, curStep: 1, velocity: v2, inCollision: true},
		{id: 0, startPos: afterPos, startStep: 1, curStep: 1, velocity: v1, inCollision: true},
	}, movings)
	require.Equal(t, []StonePath{
		{StoneID: 0, StartStep: 0, EndStep: 1, StartPos: Vector2{X: 0, Y: 0}, EndPos: Vector2{X: 0.1, Y: 0}},
	}, animations)

	// Test 2: out of board (striking stone)
	stones = []Stone{
		{ID: 0, Position: Vector2{X: 1000, Y: 0}, Radius: 1},
	}
	movings, animations = simulateCollision([]moving{
		{id: stones[0].ID, startPos: stones[0].Position, startStep: 0, curStep: 0, velocity: zeroVelocity},
	}, stones, []StonePath{}, 0.1)
	require.Equal(t, 0, len(movings))
	require.Equal(t, 1, len(animations))
}

func TestShootStone(t *testing.T) {
	g := NewGame()
	g.stones = []Stone{
		{ID: 0, Position: Vector2{X: 0, Y: 0}, Radius: 1},
		{ID: 1, Position: Vector2{X: 1, Y: 0}, Radius: 1},
	}
	evt, _ := g.ShootStone(PlayerShootData{PlayerID: 0, StoneID: 0, Velocity: Vector2{X: 1, Y: 0}})
	checkIDResult(t, idMap{0: 2, 1: 1}, evt)
	checkInitialStones(t, []Stone{
		{ID: 0, Position: Vector2{X: 0, Y: 0}, Radius: 1},
		{ID: 1, Position: Vector2{X: 1, Y: 0}, Radius: 1},
	}, evt)

	g.stones = []Stone{
		{ID: 0, Position: Vector2{X: 0, Y: 0}, Radius: 1},
		{ID: 1, Position: Vector2{X: 3, Y: 0}, Radius: 1},
		{ID: 2, Position: Vector2{X: 6, Y: 0}, Radius: 1},
	}
	evt, _ = g.ShootStone(PlayerShootData{PlayerID: 0, StoneID: 0, Velocity: Vector2{X: 6, Y: 0}})
	t.Log(evt)
	checkIDResult(t, idMap{0: 2, 1: 2, 2: 1}, evt)
	checkInitialStones(t, []Stone{
		{ID: 0, Position: Vector2{X: 0, Y: 0}, Radius: 1},
		{ID: 1, Position: Vector2{X: 3, Y: 0}, Radius: 1},
		{ID: 2, Position: Vector2{X: 6, Y: 0}, Radius: 1},
	}, evt)
}

type idMap map[int]int

func checkIDResult(t *testing.T, expected idMap, evt Event) {
	paths := evt.Data.(ShootResultData).Animation.Paths
	ids := make(idMap)
	for _, p := range paths {
		ids[p.StoneID] += 1
	}
	require.Equal(t, expected, ids)
}

func checkInitialStones(t *testing.T, expected []Stone, evt Event) {
	stones := evt.Data.(ShootResultData).Animation.InitialStones
	require.Equal(t, expected, stones)
}

func TestCheckGameOver(t *testing.T) {
	stones := []Stone{
		{StoneType: White, IsOut: true},
		{StoneType: Black},
	}
	isOver, winner := checkGameOver(stones)
	require.True(t, isOver)
	require.Equal(t, "Black", winner)

	stones = []Stone{
		{StoneType: Black},
		{StoneType: White},
	}
	isOver, winner = checkGameOver(stones)
	require.False(t, isOver)
	require.Equal(t, "", winner)
}
