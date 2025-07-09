package game

type Game struct {
	players []Player
}

type Player struct {
	id string
}

func NewGame() *Game {
	return &Game{}
}
