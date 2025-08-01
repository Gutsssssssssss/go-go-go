package game

import (
	"github.com/yanmoyy/go-go-go/internal/game"
)

// data for game state
type GameData struct {
	Turn   int
	Player game.Player
	Stones []game.Stone // all stones info (position, radius, etc)
	Size   game.Size
}
