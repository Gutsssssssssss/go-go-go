package ws

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMessage(t *testing.T) {
	t.Run("get message type", func(t *testing.T) {
		message := []byte{'G'}
		messageType, messageLength, err := getMessageType(message)
		require.NoError(t, err)
		require.Equal(t, GameEvent, messageType)
		require.Equal(t, 1, messageLength)
	})
}
