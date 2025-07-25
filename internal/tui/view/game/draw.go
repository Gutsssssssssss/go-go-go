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
	triangle := lipgloss.NewStyle().Foreground(data.IndicatorColor).Render("▲")
	if data.SelectedStoneID == stone.ID {
		g.drawTriangle(x, y+radiusH*2+triangleH, triangleH, triangle)
		switch data.Status {
		case ControlDirection:
			// degrees based on x axis
			degrees := (data.Degrees - 90 + 360) % 360
			g.drawDirection(x, y, radiusW, radiusH, degrees, "+")
		case ControlCharging:
			// TODO: draw Charging indicator
		}
		circle = lipgloss.NewStyle().Foreground(data.IndicatorColor).Render(circle)
	}
	g.drawCircle(x, y, radiusW, radiusH, circle)
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

// degrees are based on x axis
func (g grid) drawDirection(posX, posY, radiusW, radiusH float64, degrees Degrees, symbol string) {
	// y = ax + b, y = ax - b
	rad := math.Pi / 180 * float64(degrees)
	a := math.Tan(rad)
	b := posY - a*posX
	// thick:
	//  0  | 180 : radiusH
	//  90 | 270 : radiusW
	thickness := (math.Abs(math.Cos(rad)*radiusH) + math.Abs(math.Sin(rad)*radiusW)) / 4
	for y := int(posY - 2*radiusH); y <= int(posY+2*radiusH); y++ {
		for x := int(posX - 2*radiusW); x <= int(posX+2*radiusW); x++ {
			if g.outOfBounds(x, y) {
				continue
			}
			if 0 <= degrees && degrees < 90 {
				if x-int(posX) < 0 || y-int(posY) < 0 {
					continue
				}
			}
			if 90 <= degrees && degrees < 180 {
				if x-int(posX) > 0 || y-int(posY) < 0 {
					continue
				}
			}
			if 180 <= degrees && degrees < 270 {
				if x-int(posX) > 0 || y-int(posY) > 0 {
					continue
				}
			}
			if 270 <= degrees && degrees <= 360 {
				if x-int(posX) < 0 || y-int(posY) > 0 {
					continue
				}
			}
			dx := (posX - float64(x)) / radiusW
			dy := (posY - float64(y)) / radiusH
			r := dx*dx + dy*dy
			if r > 1.0 && r < 5.0 {
				var dist float64
				switch degrees {
				case 0, 180:
					dist = math.Abs(float64(y) - posY)
				case 90, 270:
					dist = math.Abs(float64(x) - posX)
				default:
					dist = distancePointToLine(point{float64(x), float64(y)}, line{a, b})
				}
				if dist <= thickness {
					g[y][x] = symbol
				}
			}
		}
	}
}

type point struct {
	x, y float64
}

type line struct {
	a, b float64
}

func distancePointToLine(p point, l line) float64 {
	return math.Abs(l.a*p.x-p.y+l.b) / math.Sqrt(math.Pow(l.a, 2)+1)
}
