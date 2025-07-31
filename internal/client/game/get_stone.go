package game

import (
	"fmt"
	"sort"

	"github.com/yanmoyy/go-go-go/internal/game"
)

func (c *GameClient) GetLeftStone(selectedStoneID int) int {
	nxt, err := c.getNextStone(selectedStoneID, -1)
	if err != nil {
		return selectedStoneID
	}
	return nxt
}

func (c *GameClient) GetRightStone(selectedStoneID int) int {
	nxt, err := c.getNextStone(selectedStoneID, 1)
	if err != nil {
		return selectedStoneID
	}
	return nxt
}

func (c *GameClient) GetCurrentStone(selectedStoneID int) int {
	cur, _ := c.getNextStone(selectedStoneID, 0)
	return cur
}

// GetPlayerStones returns stones of the player with the given playerID
// It sorts stones by x coordinate
func getFilteredStones(stones []game.Stone, stoneType game.StoneType) []game.Stone {
	var result []game.Stone
	for _, stone := range stones {
		if stone.StoneType == stoneType && !stone.IsOut {
			result = append(result, stone)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Position.X < result[j].Position.X
	})
	return result
}

func (c *GameClient) getNextStone(selectedStoneID int, direction int) (nextStoneID int, err error) {
	if c.gameData == nil {
		return selectedStoneID, fmt.Errorf("no game data")
	}
	stones := getFilteredStones(c.gameData.Stones, c.gameData.Player.StoneType)
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

func findIdx(stones []game.Stone, stoneID int) int {
	for i, stone := range stones {
		if stone.ID == stoneID {
			return i
		}
	}
	return -1
}
