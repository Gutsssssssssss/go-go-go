package game

import (
	"fmt"
)

type Player struct {
	ID           int    `json:"id"`
	StoneType StoneType `json:"stoneType"`
}
type Size struct {
	Width float64 `json:"width"`	
	Height float64 `json:"height"`
}

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Stone struct {
	ID        int     `json:"id"`
	StoneType StoneType `json:"stoneType"`
	Position  Vector2 `json:"position"`
	IsOut     bool `json:"isOut"`
	Radius    float64 `json:"radius"`
}

type StoneAnimation struct {
	StoneID   int     `json:"stoneID"`
	StartStep int     `json:"startStep"`
	EndStep   int     `json:"endStep"`
	StartPos  Vector2 `json:"startPos"`
	EndPos    Vector2 `json:"endPos"`
}

var zeroVelocity = Vector2{X: 0, Y: 0}

func (v Vector2) isZero() bool {
	return v.X == 0 && v.Y == 0
}
func (a StoneAnimation) String() string {
	return fmt.Sprintf("Animation{StoneID: %d, StartStep: %d, EndStep: %d, StartPos: %v, EndPos: %v}", a.StoneID, a.StartStep, a.EndStep, a.StartPos, a.EndPos)
}

type StoneType int

const (
	White StoneType = iota
	Black
)


func (s Stone) String() string {
	return fmt.Sprintf("Stone{ID: %d, Position: %v, Radius: %f}", s.ID, s.Position, s.Radius)
}
