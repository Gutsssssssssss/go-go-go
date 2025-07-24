package game

type playerID int
type player struct {
	id           playerID
	movableStone StoneType
}

func newPlayer(id playerID, stoneType StoneType) player {
	return player{
		id:           id,
		movableStone: stoneType,
	}
}
