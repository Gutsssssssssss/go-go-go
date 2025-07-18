package game

import "time"

type StartGameData struct {
	Turn int
}

type Vector2 struct {
	X int
	Y int
}

type ShootData struct {
	PlayerID  int
	StoneID   int
	Power     int
	Direction Vector2
}

type StoneAnimationData struct {
	StoneID   int
	StartTime time.Time
	EndTime   time.Time
	Speed     float64
	Direction Vector2
}

type GameOverData struct {
	WinnerID int
}
