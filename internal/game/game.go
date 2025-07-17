package game

type Game struct {
	Record any
}

type GameEvent struct {
}

func NewGame() *Game {
	return &Game{}
}
