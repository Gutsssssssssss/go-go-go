package game

type StoneType int

const (
	White StoneType = iota
	Black
)

type Stone struct {
	ID        int
	StoneType StoneType
	Position  Vector2
	Velocity  Vector2
	Radius    float64
}
