package game

import (
	"encoding/json"
	"fmt"
)

type DataType string

const (
	StartGame       DataType = "start_game"
	PlayerStartGame DataType = "player_start_game"
	Shoot           DataType = "shoot"
	StoneAnimations DataType = "stone_animations"
	GameOver        DataType = "game_over"
)

type Event struct {
	Type DataType `json:"type"`
	Data any      `json:"data"`
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Type: %s, Data: %+v}", e.Type, e.Data)
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type tempEvent struct {
		Type DataType        `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	var temp tempEvent
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	d, err := unmarshalData(temp.Type, temp.Data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	e.Type = temp.Type
	e.Data = d
	return nil
}

func unmarshalData(t DataType, data []byte) (any, error) {
	switch t {
	case StartGame:
		var d StartGameData
		err := json.Unmarshal(data, &d)
		return d, err
	case PlayerStartGame:
		var d PlayerStartGameData
		err := json.Unmarshal(data, &d)
		return d, err
	case Shoot:
		var d ShootData
		err := json.Unmarshal(data, &d)
		return d, err
	case StoneAnimations:
		var d StoneAnimationsData
		err := json.Unmarshal(data, &d)
		return d, err
	case GameOver:
		var d GameOverData
		err := json.Unmarshal(data, &d)
		return d, err
	}
	return nil, fmt.Errorf("unknown event type: %s", t)
}

// Data for client
type StartGameData struct {
	Turn   int     `json:"turn"`
	Stones []Stone `json:"stones"`
}

type PlayerStartGameData struct {
	Turn   playerID `json:"turn"`
	Player Player   `json:"player"`
	Stones []Stone  `json:"stones"`
	Size   `json:"size"`
}

type ShootData struct {
	PlayerID int     `json:"playerID"`
	StoneID  int     `json:"stoneID"`
	Velocity Vector2 `json:"velocity"`
}

type StoneAnimationsData struct {
	InitialStones []Stone          `json:"initialStones"`
	FinalStones   []Stone          `json:"finalStones"`
	Animations    []StoneAnimation `json:"animations"`
	MaxStep       int              `json:"maxStep"`
}

type GameOverData struct {
	WinnerID int `json:"winnerID"`
}
