package game

import (
	"math"
)

func applyFriction(velocity Vector2, friction float64) Vector2 {
	if friction <= 0 {
		return velocity
	}
	if velocity.X >= 0 {
		velocity.X = math.Max(0, velocity.X-friction)
	} else {
		velocity.X = math.Min(0, velocity.X+friction)
	}

	if velocity.Y >= 0 {
		velocity.Y = math.Max(0, velocity.Y-friction)
	} else {
		velocity.Y = math.Min(0, velocity.Y+friction)
	}

	return velocity
}

func isCollision(s1, s2 Stone) bool {
	squredDistance := (s1.Position.X-s2.Position.X)*
		(s1.Position.X-s2.Position.X) + (s1.Position.Y-s2.Position.Y)*(s1.Position.Y-s2.Position.Y)
	return squredDistance <= (s1.Radius+s2.Radius)*(s1.Radius+s2.Radius)
}

func dotProduct(v1, v2 Vector2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// returns velocities after collision
func computeCollisionVelocities(v1, v2, p1, p2 Vector2) (Vector2, Vector2) {
	const e = 0.75 // coefficient of restitution
	normal := normalizeVector(p1, p2)
	if normal.X == 0 && normal.Y == 0 {
		return v1, v2
	}

	v1Normal := dotProduct(v1, normal)
	v2Normal := dotProduct(v2, normal)
	v_rel := v1Normal - v2Normal
	J := -(1 + e) * v_rel / 2
	impulse := Vector2{X: J * normal.X, Y: J * normal.Y}

	afterV1 := Vector2{X: v1.X + impulse.X, Y: v1.Y + impulse.Y}
	afterV2 := Vector2{X: v2.X - impulse.X, Y: v2.Y - impulse.Y}
	return afterV1, afterV2
}

// returns collision direction unit vector
func normalizeVector(p1, p2 Vector2) Vector2 {
	normal := Vector2{X: p1.X - p2.X, Y: p1.Y - p2.Y}
	magnitude := math.Sqrt(normal.X*normal.X + normal.Y*normal.Y)
	if magnitude == 0 {
		return Vector2{X: 0, Y: 0}
	}
	return Vector2{X: normal.X / magnitude, Y: normal.Y / magnitude}
}

// Degrees: Max = 180, Min = -180
func ConvertToVelocity(degrees, speed float64) Vector2 {
	// 0 -> 90, 90 -> 0, 180 -> -90, -90 -> 180
	degrees = -(degrees - 90)
	dx := math.Cos(degrees / 180 * math.Pi)
	dy := math.Sin(degrees / 180 * math.Pi)
	dx = math.Round(dx*100) / 100
	dy = -math.Round(dy*100) / 100
	return Vector2{X: dx * speed, Y: dy * speed}
}

// BlenVector blends two vectors
// value is the amount of how the v1 should be blended (0.0 ~ 1.0);
// Ex) 0.0 -> v1, 1.0 -> v2
func BlendVector(v1, v2 Vector2, value float64) Vector2 {
	value = math.Min(1.0, math.Max(0.0, value))
	x := v1.X*value + v2.X*(1-value)
	y := v1.Y*value + v2.Y*(1-value)
	return Vector2{X: x, Y: y}
}
