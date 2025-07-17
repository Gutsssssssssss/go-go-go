package ws

type MessageType int

const (
	GameEvent MessageType = iota
	PlayerEvent
	ChatEvent
)

type Message struct {
	Type MessageType
	Data []byte
}

// parseMessage parses a message from a client
func parseMessage(message []byte) (Message, error) {
	return Message{}, nil
}
