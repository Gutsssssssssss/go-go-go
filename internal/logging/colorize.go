package logging

import "fmt"

type ansi int

const (
	reset = "\033[0m"
)

const (
	black ansi = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

type extendedAnsi int

const (
	darkGray extendedAnsi = 240
	gray     extendedAnsi = 248
)

func (a ansi) Color(s string) string {
	return fmt.Sprintf("\033[%dm%s%s", a, s, reset)
}

func (a extendedAnsi) Color(s string) string {
	return fmt.Sprintf("\033[38;5;%dm%s%s", a, s, reset)
}
