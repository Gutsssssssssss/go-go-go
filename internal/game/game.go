package game

import "fmt"

const (
	maxPlayers       = 2
	boardWidth       = 100
	boardHeight      = 100
	maxStones        = 10
	whiteStoneStartY = boardWidth / 4
	blackStoneStartY = boardWidth / 4 * 3
	startX           = boardWidth / 12
	stoneGap         = boardWidth / 12
)

type Game struct {
	record  []Event
	players []player
	turn    int
	idMap   map[string]int // key: userID, value: playerID (to hide userID)
	stones  []Stone
}

func NewGame() *Game {
	return &Game{
		record:  []Event{},
		players: []player{},
		turn:    0,
		idMap:   make(map[string]int),
		stones:  []Stone{},
	}
}

func (g *Game) AddPlayer(key string) (needStart bool, err error) {
	if len(g.players) >= maxPlayers {
		return false, fmt.Errorf("game is full (max players: %d)", maxPlayers)
	}
	// id is an index of g.players slice
	id := len(g.players)
	g.idMap[key] = id
	stone := white
	if len(g.players) == 1 {
		stone = black
	}
	g.players = append(g.players, newPlayer(id, stone))
	if len(g.players) == 2 {
		return true, nil
	}
	return false, nil
}

func (g *Game) StartGame() Event {
	g.placeStones()
	g.turn = 0
	evt := Event{Type: StartGameEvent, Data: StartGameData{Turn: g.turn}}
	g.record = append(g.record, evt)
	return evt
}

func (g *Game) placeStones() {
	// TODO: add stones placement game logic for better gameplay

	// white player stones
	for i := range maxStones {
		g.stones = append(g.stones,
			Stone{
				ID:        i,
				StoneType: white,
				Position: Vector2{
					X: startX + stoneGap*i,
					Y: whiteStoneStartY,
				},
			})
	}

	//black player stones
	for i := range maxStones {
		g.stones = append(g.stones,
			Stone{
				ID:        i + maxStones,
				StoneType: black,
				Position: Vector2{
					X: startX + stoneGap*i,
					Y: blackStoneStartY,
				},
			})
	}
}
