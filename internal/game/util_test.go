package game

import (
	"math"
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


func TestNormalizeVector(t *testing.T) {
	// case1 : different position
	p1 := Vector2{X: 3, Y: 4}
	p2 := Vector2{X: 0, Y: 0}
	require.Equal(t, Vector2{X: 0.6, Y: 0.8}, normalizeVector(p1, p2))

	// case2 : same position
	p1 = Vector2{X: 0, Y: 0}
	p2 = Vector2{X: 0, Y: 0}
	require.Equal(t, Vector2{X: 0, Y: 0}, normalizeVector(p1, p2))
}

func TestComputeCollisionVelocities(t *testing.T) {
	// case 1: same velocity
	v1 := Vector2{X: 1, Y: 1}
	v2 := Vector2{X: 1, Y: 1}
	p1 := Vector2{X: 0, Y: 0}
	p2 := Vector2{X: 1, Y: 1}

	actual1, actual2 := computeCollisionVelocities(v1, v2, p1, p2)
	require.Equal(t, Vector2{X: 1, Y: 1}, actual1)
	require.Equal(t, Vector2{X: 1, Y: 1}, actual2)

	// case 2: collision
	v1 = Vector2{X: 1, Y: 1}
	v2 = Vector2{X: 0, Y: 0}
	p1 = Vector2{X: 0, Y: 0}
	p2 = Vector2{X: 1, Y: 1}
	actual1, actual2 = computeCollisionVelocities(v1, v2, p1, p2)
	expected1 := Vector2{X: 0.125, Y: 0.125}
	expected2 := Vector2{X: 0.875, Y: 0.875}
	require.True(t, withinTolerance(actual1.X, expected1.X, 0.0001))
	require.True(t, withinTolerance(actual1.Y, expected1.Y, 0.0001))
	require.True(t, withinTolerance(actual2.X, expected2.X, 0.0001))
	require.True(t, withinTolerance(actual2.Y, expected2.Y, 0.0001))
}

func withinTolerance(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestDotProduct(t *testing.T) {
	velocity1 := Vector2{X: 1, Y: 2}
	velocity2 := Vector2{X: 3, Y: 4}
	require.Equal(t, 11.0, dotProduct(velocity1, velocity2))
}


func TestBlendVector(t *testing.T) {
	v1 := Vector2{X: 1, Y: 0}
	v2 := Vector2{X: 0, Y: 1}
	require.Equal(t, Vector2{X: 0.5, Y: 0.5}, BlendVector(v1, v2, 0.5))
	require.Equal(t, Vector2{X: 0, Y: 1}, BlendVector(v1, v2, 0))
	require.Equal(t, Vector2{X: 1, Y: 0}, BlendVector(v1, v2, 1))
}
func TestConvertToVelocity(t *testing.T) {
	// Test 1: 0 power
	require.Equal(t, Vector2{X: 0, Y: 0}, ConvertToVelocity(0, 0))
	require.Equal(t, Vector2{X: 0, Y: 0}, ConvertToVelocity(180, 0))

	// Test 2: 10 power
	require.Equal(t, Vector2{X: 0, Y: -10}, ConvertToVelocity(0, 10))
	require.Equal(t, Vector2{X: 0, Y: 10}, ConvertToVelocity(180, 10))
	require.Equal(t, Vector2{X: 10, Y: 0}, ConvertToVelocity(90, 10))
	require.Equal(t, Vector2{X: -10, Y: 0}, ConvertToVelocity(-90, 10))

	// Test 3: Random Cases
	require.Equal(t, 1.0, ConvertToVelocity(30, 2).X)
	require.Equal(t, -1.0, ConvertToVelocity(60, 2).Y)
}