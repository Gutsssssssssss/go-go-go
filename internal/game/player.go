package game

type player struct {
	id           int
	movableStone StoneType
}

func newPlayer(id int, stoneType StoneType) player {
	return player{
		id:           id,
		movableStone: stoneType,
	}
}
