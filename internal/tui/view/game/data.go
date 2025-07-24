package game

import "github.com/charmbracelet/lipgloss"

type ControlStatus int

const (
	ControlSelectStone ControlStatus = iota
	ControlDirection
	ControlCharging
)

type Direction struct{}

type Power float64

type ControlData struct {
	Status          ControlStatus
	IndicatorColor  lipgloss.Color
	SelectedStoneID int
	Direction       Direction
	Power           Power
}
