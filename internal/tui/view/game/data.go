package game

import "github.com/charmbracelet/lipgloss"

type ControlStatus int

const (
	ControlSelectStone ControlStatus = iota
	ControlDirection
	ControlCharging
)

const (
	MaxPower = 10
	MinPower = 0

	MaxDegrees = 180
	MinDegrees = -180
)

type Degrees int // degrees are starting from 0 to 360 (clockwise, from upper (↑))

type ControlData struct {
	Status          ControlStatus
	SelectedStoneID int
	Degrees         Degrees
	IndicatorColor  lipgloss.Color
}
