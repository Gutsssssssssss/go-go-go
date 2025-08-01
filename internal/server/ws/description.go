package ws

import (
	"fmt"

	"github.com/yanmoyy/go-go-go/internal/game"
)

func getDescription(evt game.Event) string {
	switch evt.Type {
	case game.ShootResult:
		data := evt.Data.(game.ShootResultData)
		return fmt.Sprintf("%s player shoots!", data.Shooter)
	}
	return ""
}
