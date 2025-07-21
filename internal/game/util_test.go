package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplyFriction(t *testing.T) {
	v := Vector2{X: 1, Y: 1}
	v = applyFriction(v, 0.1)
	require.Equal(t, 0.9, v.X)
	require.Equal(t, 0.9, v.Y)

	v = Vector2{X: 1, Y: 1}
	v = applyFriction(v, 0.5)
	require.Equal(t, 0.5, v.X)
	require.Equal(t, 0.5, v.Y)

	v = Vector2{X: 1, Y: 1}
	v = applyFriction(v, 1)
	require.Equal(t, 0.0, v.X)
	require.Equal(t, 0.0, v.Y)

	v = Vector2{X: 1, Y: 1}
	v = applyFriction(v, 2)
	require.Equal(t, 0.0, v.X)
	require.Equal(t, 0.0, v.Y)

	v = Vector2{X: -1, Y: -1}
	v = applyFriction(v, 0.1)
	require.Equal(t, -0.9, v.X)
	require.Equal(t, -0.9, v.Y)

	v = Vector2{X: -1, Y: -1}
	v = applyFriction(v, -1)
	require.Equal(t, -1.0, v.X)
	require.Equal(t, -1.0, v.Y)
}

func TestIsCollision(t *testing.T) {
	// case1 : intersaction
	stone1 := Stone{
		Position: Vector2{X: 0, Y: 0},
		Radius:   1,
	}
	stone2 := Stone{
		Position: Vector2{X: 1, Y: 1},
		Radius:   1,
	}
	require.True(t, isCollision(stone1, stone2))

	// case2 : contact
	stone1 = Stone{
		Position: Vector2{X: 0, Y: 0},
		Radius:   1,
	}
	stone2 = Stone{
		Position: Vector2{X: 2, Y: 0},
		Radius:   1,
	}
	require.True(t, isCollision(stone1, stone2))

	// case3 : no intersaction
	stone1 = Stone{
		Position: Vector2{X: 0, Y: 0},
		Radius:   1,
	}
	stone2 = Stone{
		Position: Vector2{X: 2, Y: 1},
		Radius:   1,
	}
	require.False(t, isCollision(stone1, stone2))

	// case4 : inside
	stone1 = Stone{
		Position: Vector2{X: 0, Y: 0},
		Radius:   1,
	}
	stone2 = Stone{
		Position: Vector2{X: 0, Y: 0},
		Radius:   2,
	}
	require.True(t, isCollision(stone1, stone2))
}
