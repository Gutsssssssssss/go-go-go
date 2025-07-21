package game

type StoneType int

const (
	white StoneType = iota
	black
)

type Stone struct {
	ID        int
	StoneType StoneType
	Position  Vector2
	Velocity  Vector2
	Radius    float64
}
