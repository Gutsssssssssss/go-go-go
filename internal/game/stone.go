package game

import "fmt"

type StoneType int

const (
	White StoneType = iota
	Black
)

type Stone struct {
	ID        int
	StoneType StoneType
	Position  Vector2
	isOut       bool
	Radius    float64
}

func (s Stone) String() string {
	return fmt.Sprintf("Stone{ID: %d, Position: %v, Radius: %f}", s.ID, s.Position, s.Radius)
}
