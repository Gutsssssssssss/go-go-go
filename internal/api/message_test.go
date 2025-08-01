package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yanmoyy/go-go-go/internal/game"
)

func TestUnmarshalMessage(t *testing.T) {
	t.Run("GameEventMessage", func(t *testing.T) {
		m := Message{
			Type: GameEventMsg,
			Data: game.Event{
				Type: game.PlayerShoot,
				Data: game.PlayerShootData{
					PlayerID: 0,
					StoneID:  1,
					Velocity: game.Vector2{X: 1, Y: 0},
				},
			},
		}
		jsonData, err := json.Marshal(m)
		require.NoError(t, err)
		var msg Message
		err = json.Unmarshal(jsonData, &msg)
		require.NoError(t, err)
		require.Equal(t, m, msg)
	})

	t.Run("ResponseMessage", func(t *testing.T) {
		m := Message{
			Type: ResponseMsg,
			Data: Response{
				ID:      "id",
				Status:  "status",
				Message: "message",
			},
		}
		jsonData, err := json.Marshal(m)
		require.NoError(t, err)
		var msg Message
		err = json.Unmarshal(jsonData, &msg)
		require.NoError(t, err)
		require.Equal(t, m, msg)
	})
}
