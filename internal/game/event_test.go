package game

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameEventUnmarshal(t *testing.T) {
	t.Run("game start", func(t *testing.T) {
		evt := Event{
			Type: StartGameEvent,
			Data: StartGameData{Turn: 0},
		}
		jsonData, err := json.Marshal(evt)
		require.NoError(t, err)

		var evt2 Event
		err = json.Unmarshal(jsonData, &evt2)
		require.NoError(t, err)
		require.Equal(t, evt, evt2)

		data, ok := evt2.Data.(StartGameData)
		require.True(t, ok)
		require.Equal(t, playerID(0), data.Turn)
	})
}
