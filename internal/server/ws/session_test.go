package ws

import "testing"

func TestGameEventToJSON(t *testing.T) {
	t.Run("game start", func(t *testing.T) {
		evt := GameEvent{Type: GameStart, Data: StartGameData{Turn: 0}}
		jsonData, err := evt.ToJSON()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(jsonData))
	})
}
