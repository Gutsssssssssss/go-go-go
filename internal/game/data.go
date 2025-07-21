package game

import "time"

type StartGameData struct {
	Turn int
}

type Vector2 struct {
	X float64
	Y float64
}

type ShootData struct {
	PlayerID int
	StoneID  int
	Velocity Vector2
}

type StoneAnimationData struct {
	StoneID   int
	StartTime time.Time
	EndTime   time.Time
	StartPos  Vector2
	EndPos    Vector2
}

type GameOverData struct {
	WinnerID int
}
