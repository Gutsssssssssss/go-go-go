package ws

import "fmt"

type MessageType int

const (
	GameEvent MessageType = iota
	ChatEvent
)

// getMessageType returns the message type and length of the message which is read from the websocket connection
func getMessageType(payload []byte) (MessageType, int, error) {
	if len(payload) < 1 {
		return 0, 0, fmt.Errorf("message too short")
	}
	switch payload[0] {
	case 'G':
		return GameEvent, 1, nil
	case 'C':
		return ChatEvent, 1, nil
	default:
		return 0, 0, fmt.Errorf("unknown message type")
	}
}

func parseGameEvent(data []byte) error {
}
