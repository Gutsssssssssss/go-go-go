package game

import (
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/game"
)

func (g grid) drawStone(stone game.Stone, scale scale, data ControlData) {
	const triangleSize = 1.0

	x := stone.Position.X * scale.width
	y := stone.Position.Y * scale.height
	radiusW := stone.Radius * scale.width
	radiusH := stone.Radius * scale.height

	triangleH := triangleSize * scale.height

	var circle string
	if stone.StoneType == game.White {
		circle = "●"
	} else {
		circle = "◯"
	}
	g.drawCircle(x, y, radiusW, radiusH, circle)
	triangle := lipgloss.NewStyle().Foreground(data.IndicatorColor).Render("▲")
	if data.SelectedStoneID == stone.ID {
		g.drawTriangle(x, y+radiusH*2+triangleH, triangleH, triangle)
		switch data.Status {
		case ControlDirection:
			// TODO: draw Direction indicator

		case ControlCharging:
			// TODO: draw Charging indicator
		}
	}
}

// drawCircle draws a circle on the grid
func (g grid) drawCircle(posX, posY, radiusW, radiusH float64, symbol string) {
	if radiusW == 0 || radiusH == 0 {
		return
	}
	for y := int(posY - radiusH); y <= int(posY+radiusH); y++ {
		for x := int(posX - radiusW); x <= int(posX+radiusW); x++ {
			if g.outOfBounds(x, y) {
				continue
			}
			dx := (posX - float64(x)) / radiusW
			dy := (posY - float64(y)) / radiusH
			if dx*dx+dy*dy <= 1.0 {
				g[y][x] = symbol
			}
		}
	}
}

func (g grid) drawTriangle(posX, posY, height float64, symbol string) {
	if height == 0 {
		return
	}
	for k := 0; k <= int(height); k++ {
		y := int(posY) + k
		for x := int(math.Round(posX)) - k; x <= int(math.Round(posX))+k; x++ {
			if g.outOfBounds(x, y) {
				continue
			}
			g[y][x] = symbol
		}
	}
}
