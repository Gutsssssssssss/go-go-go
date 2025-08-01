package view

import (
	"unicode"

	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/api"
	"github.com/yanmoyy/go-go-go/internal/tui/color"
)

const (
	padding      = 1
	minWidthLine = " 12:34 |  [B] " // for calculating min width
)

type MessagesProps struct {
	Width  int
	Height int
}

var minWidth = lipgloss.Width(minWidthLine)

// ServerMessages is a view for messages. It shows the last n lines of the message.
func ServerMessages(messages []api.ServerMessage, props MessagesProps) string {
	if len(messages) == 0 || props.Width < minWidth+(2*padding) {
		return ""
	}
	style := lipgloss.NewStyle().
		Width(props.Width).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(color.Gray).
		Padding(0, padding).
		AlignHorizontal(lipgloss.Left)

	h := lipgloss.Height(style.Render("temp"))
	n := min(props.Height/h, len(messages))
	messages = messages[len(messages)-n:]

	strs := make([]string, len(messages))

	white := lipgloss.NewStyle().Foreground(color.White)
	gray := lipgloss.NewStyle().Foreground(color.Gray)
	blue := lipgloss.NewStyle().Foreground(color.GolangBlue)
	for i, msg := range messages {
		content := msg.Message
		if len(content) > props.Width-minWidth-2*padding {
			content = content[:props.Width-minWidth-2*padding-3] + "..."
		}
		var line string
		line += gray.Render(msg.Time.Format("15:04 | "))
		switch msg.Type {
		case api.ServerChat:
			content = white.Render(content)
		case api.ServerGame:
			content = gray.Render(content)
		}
		line += content
		if len(msg.From) > 0 {
			initial := unicode.ToUpper(rune(msg.From[0]))
			switch initial {
			case 'B':
				line += white.Render(" [", gray.Render("B"), white.Render("]"))
			case 'W':
				line += white.Render(" [", white.Render("W"), white.Render("]"))
			case 'S':
				line += white.Render(" [", blue.Render("S"), white.Render("]"))
			}
		}
		strs[i] = style.Render(line)
	}
	return lipgloss.JoinVertical(lipgloss.Left, strs...)
}
