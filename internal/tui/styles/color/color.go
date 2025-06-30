package color

import (
	"fmt"
	"strconv"
	"strings"
)

// reference: https://github.com/fatih/color/

// Attribute defines a single SGR Code
type Attribute int

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Forground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const escape = "\x1b"

type Color struct {
	params []Attribute
}

// New returns a newly created color object.
func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	c.Add(value...)
	return c
}

// Add is used to chain SGR Codes and apply them all at once
func (c *Color) Add(value ...Attribute) {
	c.params = append(c.params, value...)
}

func (c *Color) Sprint(a ...any) string {
	return c.wrap(fmt.Sprint(a...))
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}
	return strings.Join(format, ";")
}

// wrap wraps the given string with the color codes. The string is ready to
// be printed.
func (c *Color) wrap(s string) string {
	return c.format() + s + c.unformat()
}

// format part of the output
func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

// unformat part of the output
func (c *Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}
