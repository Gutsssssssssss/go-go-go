package game

import "fmt"

type StartGameData struct {
	Turn playerID `json:"turn"`
}

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var zeroVelocity = Vector2{X: 0, Y: 0}

func (v Vector2) isZero() bool {
	return v.X == 0 && v.Y == 0
}

type ShootData struct {
	PlayerID int     `json:"playerID"`
	StoneID  int     `json:"stoneID"`
	Velocity Vector2 `json:"velocity"`
}

type StoneAnimation struct {
	StoneID   int     `json:"stoneID"`
	StartStep int     `json:"startStep"`
	EndStep   int     `json:"endStep"`
	StartPos  Vector2 `json:"startPos"`
	EndPos    Vector2 `json:"endPos"`
}

func (a StoneAnimation) String() string {
	return fmt.Sprintf("Animation{StoneID: %d, StartStep: %d, EndStep: %d, StartPos: %v, EndPos: %v}", a.StoneID, a.StartStep, a.EndStep, a.StartPos, a.EndPos)
}

type StoneAnimationsData struct {
	InitialStones []Stone          `json:"initialStones"`
	Animations    []StoneAnimation `json:"animations"`
	MaxStep       int              `json:"maxStep"`
}

type GameOverData struct {
	WinnerID int `json:"winnerID"`
}
