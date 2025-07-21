package game

import "math"

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

func isCollision(stone1, stone2 Stone) bool {
	squredDistance := (stone1.Position.X-stone2.Position.X)*
		(stone1.Position.X-stone2.Position.X) + (stone1.Position.Y-stone2.Position.Y)*(stone1.Position.Y-stone2.Position.Y)
	return squredDistance <= (stone1.Radius+stone2.Radius)*(stone1.Radius+stone2.Radius)
}
