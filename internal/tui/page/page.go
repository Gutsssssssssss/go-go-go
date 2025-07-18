package page

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
)

type PageID string

type page struct {
	tea.Model
	layout.Sizable
	window struct {
		width, height int
	}
}

func (p *page) SetSize(width, height int) {
	p.window.width = width
	p.window.height = height
}

func (p *page) GetSize() (int, int) {
	return p.window.width, p.window.height
}

// Messages

type PagePushMsg struct {
	ID PageID
}

type PagePopMsg struct{}

func cmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
