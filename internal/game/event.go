package game

import (
	"encoding/json"
	"time"
)

type EventType int

const (
	GameStart EventType = iota
	Shoot
	StoneAnimation
	GameOver
)

type GameEvent struct {
	Type EventType
	Data any
}

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

func (e GameEvent) ToJSON() ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return b, nil
}
