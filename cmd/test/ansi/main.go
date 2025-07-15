package main

import (
	"fmt"
)

// ANSI escape code constants
type ANSI string

const (
	Reset        ANSI = "\033[0m" // Resets all styles and colors
	Bold         ANSI = "\033[1m" // Bold text
	Faint        ANSI = "\033[2m" // Faint text
	Italic       ANSI = "\033[3m" // Italic text
	Underline    ANSI = "\033[4m" // Underlined text
	BlinkSlow    ANSI = "\033[5m" // Blink slowly
	BlinkRapid   ANSI = "\033[6m" // Blink rapidly
	ReverseVideo ANSI = "\033[7m" // Reverse video
	Concealed    ANSI = "\033[8m" // Concealed text
	CrossedOut   ANSI = "\033[9m" // Crossed out text

	FgBlack   ANSI = "\033[30m" // Foreground black
	FgRed     ANSI = "\033[31m" // Foreground red
	FgGreen   ANSI = "\033[32m" // Foreground green
	FgYellow  ANSI = "\033[33m" // Foreground yellow
	FgBlue    ANSI = "\033[34m" // Foreground blue
	FgMagenta ANSI = "\033[35m" // Foreground magenta
	FgCyan    ANSI = "\033[36m" // Foreground cyan
	FgWhite   ANSI = "\033[37m" // Foreground white

	BgBlack   ANSI = "\033[40m" // Background black
	BgRed     ANSI = "\033[41m" // Background red
	BgGreen   ANSI = "\033[42m" // Background green
	BgYellow  ANSI = "\033[43m" // Background yellow
	BgBlue    ANSI = "\033[44m" // Background blue
	BgCyan    ANSI = "\033[46m" // Background cyan
	BgMagenta ANSI = "\033[45m" // Background magenta
	BgWhite   ANSI = "\033[47m" // Background white

	// extended colors
)

var (
	red   = RGB(255, 0, 0)
	green = RGB(0, 255, 0)
)

func RGB(r, g, b uint8) func(string) string {
	return func(s string) string {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, s)
	}
}

func main() {
	fmt.Println(FgBlack + "Black text" + Reset)
	fmt.Println(FgRed + "Red text" + Reset)
	fmt.Println(FgGreen + "Green text" + Reset)
	fmt.Println(FgBlue + "Blue text" + Reset)
	fmt.Println(FgYellow + "Yellow text" + Reset)
	fmt.Println(FgCyan + "Cyan text" + Reset)
	fmt.Println(FgMagenta + "Magenta text" + Reset)

	// Background colors
	fmt.Println(BgRed + "Text on red background" + Reset)
	fmt.Println(BgGreen + "Text on green background" + Reset)
	fmt.Println(BgBlue + "Text on blue background" + Reset)

	// Combining foreground and background
	fmt.Println(FgRed + BgYellow + "Red text on yellow background" + Reset)
	fmt.Println(FgGreen + BgBlue + "Green text on blue background" + Reset)

	// Add styles
	fmt.Println(Bold + FgRed + "Bold red text" + Reset)
	fmt.Println(Underline + FgBlue + "Underlined blue text" + Reset)
	fmt.Println(Bold + Underline + FgCyan + "Bold and underlined cyan text" + Reset)
	fmt.Println(Italic + FgRed + "Italic red text" + Reset)
	fmt.Println(Italic + Underline + FgBlue + "Italic and underlined blue text" + Reset)
	// other
	fmt.Println(BlinkSlow + FgRed + "Blink slow red text" + Reset)
	fmt.Println(BlinkRapid + FgBlue + "Blink rapid blue text" + Reset)
	fmt.Println(BlinkSlow + Underline + FgYellow + "Blink slow and underlined yellow text" + Reset)
	fmt.Println(CrossedOut + FgGreen + "Crossed out green text" + Reset)
	fmt.Println(CrossedOut + Underline + FgMagenta + "Crossed out and underlined magenta text" + Reset)
	fmt.Println(ReverseVideo + FgCyan + "Reverse video cyan text" + Reset)
	fmt.Println(Concealed + FgWhite + "Concealed white text" + Reset)
	fmt.Println(Bold + Faint + FgRed + "Bold faint black text" + Reset)

	// Example: Alkkagi-like game board snippet
	fmt.Println("\nAlkkagi Board Example:")
	fmt.Println(BgBlue + "    " + Reset + " " + BgBlue + "    " + Reset + " " + BgBlue + "    " + Reset)
	fmt.Println(FgRed + " O " + Reset + "   " + FgGreen + " X " + Reset + "   " + Reset) // Player 1 (O), Player 2 (X)
	fmt.Println(BgBlue + "    " + Reset + " " + BgBlue + "    " + Reset + " " + BgBlue + "    " + Reset)
	fmt.Println(Bold + FgYellow + "Player 1's turn" + Reset)

	// Extended colors
	fmt.Println("\nExtended colors:")
	for i := 0; i < 256; i++ {
		fmt.Printf("\033[38;5;%dm%3d\033[0m ", i, i)
	}

	// 24 bit colors
	fmt.Println("\n24 bit colors:")
	fmt.Println(RGB(255, 0, 0)("Red"))
	fmt.Println(RGB(0, 255, 0)("Green"))
}
