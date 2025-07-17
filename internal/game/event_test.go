package game

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameEventToJSON(t *testing.T) {
	t.Run("game start", func(t *testing.T) {
		evt := GameEvent{Type: GameStart, Data: StartGameData{Turn: 0}}
		jsonData, err := evt.ToJSON()
		require.NoError(t, err)
		var evt2 GameEvent
		err = json.Unmarshal(jsonData, &evt2)
		require.NoError(t, err)
		require.Equal(t, evt, evt2)
	})
}
