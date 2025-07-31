package game

import (
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/game"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

func (g grid) drawStones(stones []game.Stone, scale scale, data ControlData) {
	for _, stone := range stones {
		if !stone.IsOut {
			g.drawStone(stone, scale, data)
		}
	}
}

func (g grid) drawStone(stone game.Stone, scale scale, data ControlData) {
	x := stone.Position.X * scale.width
	y := stone.Position.Y * scale.height
	radiusW := stone.Radius * scale.width
	radiusH := stone.Radius * scale.height

	var circle string
	if stone.StoneType == game.White {
		circle = "●"
	} else {
		circle = "◯"
	}
	if data.SelectedStoneID == stone.ID {
		circle = lipgloss.NewStyle().Foreground(data.IndicatorColor).Render(circle)
	}
	g.drawCircle(x, y, radiusW, radiusH, circle)
}

func (g grid) drawAnimation(anim game.StonePath, curStep int, scale scale, stone game.Stone) {
	v := game.BlendVector(anim.EndPos, anim.StartPos,
		float64(curStep-anim.StartStep)/float64(anim.EndStep-anim.StartStep),
	)
	startX := anim.StartPos.X * scale.width
	startY := anim.StartPos.Y * scale.height
	curX := v.X * scale.width
	curY := v.Y * scale.height
	radiusW := stone.Radius * scale.width
	radiusH := stone.Radius * scale.height
	var circle string
	if stone.StoneType == game.White {
		circle = "●"
	} else {
		circle = "◯"
	}
	g.drawCircle(startX, startY, radiusW, radiusH, " ") // clear previous animation
	g.drawCircle(curX, curY, radiusW, radiusH, circle)
}

func (g grid) drawIndicator(stone game.Stone, scale scale, data ControlData) {
	const triangleSize = 1.5

	x := stone.Position.X * scale.width
	y := stone.Position.Y * scale.height
	radiusW := stone.Radius * scale.width
	radiusH := stone.Radius * scale.height

	switch data.Status {
	case ControlSelectStone:
		triangleH := triangleSize * scale.height
		triangle := lipgloss.NewStyle().Foreground(data.IndicatorColor).Render("▲")
		g.drawTriangle(x, y+radiusH*2+triangleH, triangleH, triangle)
	case ControlDirection, ControlCharging:
		// degrees based on x axis
		degrees := (data.Degrees - 90 + 360) % 360
		g.drawDirection(x, y, radiusW, radiusH, degrees, lipgloss.NewStyle().Foreground(color.GolangBlue).Render("+"))
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
	for k := range int(math.Round(height)) {
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
	thickness := (math.Abs(math.Cos(rad)*radiusH) + math.Abs(math.Sin(rad)*radiusW)) / 3
	for y := int(posY - 3*radiusH); y <= int(posY+3*radiusH); y++ {
		for x := int(posX - 3*radiusW); x <= int(posX+3*radiusW); x++ {
			if g.outOfBounds(x, y) {
				continue
			}
			// left: 0, down: 90, right: 180, up: 270
			if 315 <= degrees || degrees < 45 {
				if x-int(posX) <= 0 {
					continue
				}
			}
			if 45 <= degrees && degrees < 135 {
				if y-int(posY) <= 0 {
					continue
				}
			}
			if 135 <= degrees && degrees < 225 {
				if x-int(posX) >= 0 {
					continue
				}
			}
			if 225 <= degrees && degrees < 315 {
				if y-int(posY) >= 0 {
					continue
				}
			}
			dx := (posX - float64(x)) / radiusW
			dy := (posY - float64(y)) / radiusH
			r := dx*dx + dy*dy
			if r > 1.0 && r <= 8.0 {
				var dist float64
				switch degrees {
				case 0, 180:
					dist = math.Abs(float64(y) - posY)
				case 90, 270:
					dist = math.Abs(float64(x) - posX)
				default:
					dist = distancePointToLine(point{float64(x), float64(y)}, line{a, b})
				}
				if dist <= float64(thickness) {
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
	return math.Abs(l.a*p.x-p.y+l.b) / math.Sqrt(l.a*l.a+1)
}
