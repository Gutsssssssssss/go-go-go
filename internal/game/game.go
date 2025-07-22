package game

import "fmt"

const (
	maxPlayers               = 2
	boardWidth       float64 = 100.0
	boardHeight      float64 = 100.0
	maxStones                = 10
	whiteStoneStartY         = boardHeight / 4
	blackStoneStartY         = boardHeight / 4 * 3
	startX                   = boardWidth / 11
	stoneGap                 = boardWidth / 11
	friction                 = 0.1
	stoneRadius              = 1.5
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
	stone := White
	if len(g.players) == 1 {
		stone = Black
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
				StoneType: White,
				Radius:    stoneRadius,
				Position: Vector2{
					X: startX + stoneGap*float64(i),
					Y: whiteStoneStartY,
				},
			})
	}

	//black player stones
	for i := range maxStones {
		g.stones = append(g.stones,
			Stone{
				ID:        i + maxStones,
				StoneType: Black,
				Radius:    stoneRadius,
				Position: Vector2{
					X: startX + stoneGap*float64(i),
					Y: blackStoneStartY,
				},
			})
	}
}

func (g *Game) GetStones() []Stone {
	return g.stones
}
func (g *Game) GetSize() (float64, float64) {
	return boardWidth, boardHeight
}

func (g *Game) shootStone(shootData ShootData) {
	var movingStone []int
	g.stones[shootData.StoneID].Velocity = shootData.Velocity
	movingStone = append(movingStone, shootData.StoneID)
	for {
		if len(movingStone) == 0 {
			break
		}
		for _, stoneID := range movingStone {
			g.stones[stoneID].Position.X += g.stones[stoneID].Velocity.X
			g.stones[stoneID].Position.Y += g.stones[stoneID].Velocity.Y
			g.stones[stoneID].Velocity = applyFriction(g.stones[stoneID].Velocity, friction)
			for _, stone := range g.stones {
				if stone.ID == stoneID {
					continue
				}
				// TODO : collision detection
			}
		}
	}
}
