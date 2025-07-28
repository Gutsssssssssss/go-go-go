package game

import (
	"fmt"
	"sort"
)

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
	turn    playerID
	idMap   map[string]playerID // key: userID, value: playerID (to hide userID)
	stones  []Stone
}

func NewGame() *Game {
	return &Game{
		record:  []Event{},
		players: []player{},
		idMap:   make(map[string]playerID),
		stones:  []Stone{},
	}
}

func (g *Game) AddPlayer(uuid string) (needStart bool, err error) {
	if len(g.players) >= maxPlayers {
		return false, fmt.Errorf("game is full (max players: %d)", maxPlayers)
	}
	// id is an index of g.players slice
	id := playerID(len(g.players))
	g.idMap[uuid] = id
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
	g.turn = 1
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

// GetPlayerStones returns stones of the player with the given playerID
// It sorts stones by x coordinate
func (g *Game) getPlayerStones(id playerID) []Stone {
	var stones []Stone
	for _, stone := range g.stones {
		if stone.StoneType == g.players[id].movableStone && !stone.IsOut {
			stones = append(stones, stone)
		}
	}
	sort.Slice(stones, func(i, j int) bool {
		return stones[i].Position.X < stones[j].Position.X
	})
	return stones
}

func (g *Game) GetSize() (float64, float64) {
	return boardWidth, boardHeight
}

func (g *Game) getNextStone(playerID playerID, selectedStoneID int, direction int) (nextStoneID int, err error) {
	stones := g.getPlayerStones(playerID)
	if len(stones) == 0 {
		return selectedStoneID, fmt.Errorf("no stones found")
	}
	idx := findIdx(stones, selectedStoneID)
	if idx == -1 {
		idx = 0
	}
	nextIdx := (idx + direction + len(stones)) % len(stones)
	return stones[nextIdx].ID, nil
}

func (g *Game) GetLeftStone(playerID playerID, selectedStoneID int) int {
	nxt, err := g.getNextStone(playerID, selectedStoneID, -1)
	if err != nil {
		return selectedStoneID
	}
	return nxt
}

func (g *Game) GetRightStone(playerID playerID, selectedStoneID int) int {
	nxt, err := g.getNextStone(playerID, selectedStoneID, 1)
	if err != nil {
		return selectedStoneID
	}
	return nxt
}
func (g *Game) GetCurrentStone(playerID playerID, selectedStoneID int) int {
	cur, err := g.getNextStone(playerID, selectedStoneID, 0)
	if err != nil {
		return selectedStoneID
	}
	return cur
}

func findIdx(stones []Stone, stoneID int) int {
	for i, stone := range stones {
		if stone.ID == stoneID {
			return i
		}
	}
	return -1
}

type moving struct {
	id          int
	startPos    Vector2
	velocity    Vector2
	startStep   int
	curStep     int
	inCollision bool
}

func (m moving) String() string {
	return fmt.Sprintf("moving{id: %d, startPos: %v, velocity: %v, startStep: %d, curStep: %d}", m.id, m.startPos, m.velocity, m.startStep, m.curStep)
}

func addAnimation(animations []StoneAnimation, mov moving, endPosition Vector2) []StoneAnimation {
	return append(animations, StoneAnimation{
		StoneID:   mov.id,
		StartStep: mov.startStep,
		EndStep:   mov.curStep,
		StartPos:  mov.startPos,
		EndPos:    endPosition,
	})
}

func simulateCollision(movings []moving, stones []Stone, animations []StoneAnimation, dt float64) ([]moving, []StoneAnimation) {
	var nextMovings []moving
	if len(movings) == 0 {
		return nextMovings, animations
	}
	for _, mov := range movings {
		id := mov.id
		stones[id].Position.X += mov.velocity.X * dt
		stones[id].Position.Y += mov.velocity.Y * dt
		mov.curStep += 1
		if outOfBoard(stones[id].Position) {
			stones[id].IsOut = true
			animations = addAnimation(animations, mov, stones[id].Position)
			continue
		}
		mov.velocity = applyFriction(mov.velocity, friction*dt)
		if mov.velocity.isZero() {
			animations = addAnimation(animations, mov, stones[id].Position)
			continue
		}
		hasCollision := false // check collision only once
		for _, target := range stones {
			if target.ID == id || target.IsOut {
				continue
			}
			if isCollision(stones[id], target) {
				hasCollision = true
				if mov.inCollision {
					break
				}
				v1, v2 := computeCollisionVelocities(mov.velocity, zeroVelocity, stones[id].Position, target.Position)
				mov.velocity = v1
				if !v2.isZero() {
					nextMovings = append(nextMovings,
						moving{
							id:          target.ID,
							startPos:    stones[target.ID].Position,
							velocity:    v2,
							startStep:   mov.curStep,
							curStep:     mov.curStep,
							inCollision: true,
						},
					)
				}
				animations = addAnimation(animations, mov, stones[id].Position)
				break
			}
		}
		if mov.inCollision {
			if !hasCollision {
				mov.inCollision = false
			}
		} else {
			if hasCollision {
				mov.inCollision = true
				mov.startStep = mov.curStep
				mov.startPos = stones[id].Position
			}
		}
		nextMovings = append(nextMovings, mov)
	}
	return nextMovings, animations
}

func (g *Game) ShootStone(shootData ShootData) Event {
	// TODO: add shootData another field. for better client side abstraction
	striking := g.stones[shootData.StoneID]
	if striking.IsOut {
		return Event{}
	}
	initialStones := make([]Stone, len(g.stones))
	copy(initialStones, g.stones)
	animations := []StoneAnimation{}
	movings := []moving{
		{id: striking.ID, startPos: striking.Position, velocity: shootData.Velocity,
			startStep: 0, curStep: 0, inCollision: false},
	}
	for {
		movings, animations = simulateCollision(movings, g.stones, animations, 0.1)
		if len(movings) == 0 {
			break
		}
	}
	maxStep := 0
	for _, anim := range animations {
		if anim.EndStep > maxStep {
			maxStep = anim.EndStep
		}
	}
	evt := Event{Type: StoneAnimationsEvent, Data: StoneAnimationsData{
		InitialStones: initialStones,
		Animations:    animations,
		MaxStep:       maxStep,
	}}
	g.record = append(g.record, evt)
	return evt
}

func outOfBoard(pos Vector2) bool {
	return pos.X < 0 || pos.X > boardWidth || pos.Y < 0 || pos.Y > boardHeight
}
