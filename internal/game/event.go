package game

import (
	"encoding/json"
	"fmt"
)

type EventType int

const (
	StartGameEvent EventType = iota
	ShootEvent
	StoneAnimationsEvent
	GameOverEvent
)

type Event struct {
	Type EventType `json:"type"`
	Data any       `json:"data"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type tempEvent struct {
		Type EventType       `json:"type"`
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

func unmarshalData(t EventType, data []byte) (any, error) {
	switch t {
	case StartGameEvent:
		var d StartGameData
		err := json.Unmarshal(data, &d)
		return d, err
	case ShootEvent:
		var d ShootData
		err := json.Unmarshal(data, &d)
		return d, err
	case StoneAnimationsEvent:
		var d StoneAnimationsData
		err := json.Unmarshal(data, &d)
		return d, err
	case GameOverEvent:
		var d GameOverData
		err := json.Unmarshal(data, &d)
		return d, err
	}
	return nil, fmt.Errorf("unknown event type: %d", t)
}
