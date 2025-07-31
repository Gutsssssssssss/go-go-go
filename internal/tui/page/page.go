package page

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
)

type PageID string

type page struct {
	tea.Model
	layout.Sizable
	window struct {
		width, height int
	}

	// message
	message      string
	messageTimer *time.Timer
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

type PageSwitchMsg struct {
	ID PageID
}

type PagePopMsg struct{}

func cmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func (p *page) showMessage(msg string) {
	p.message = msg
	if p.messageTimer == nil {
		p.messageTimer = time.AfterFunc(time.Second*1, func() {
			p.message = ""
			p.messageTimer = nil
		})
	} else {
		p.messageTimer.Reset(time.Second * 1)
	}
}

func messageView(message string) string {
	return lipgloss.NewStyle().
		Margin(0, 0, 2, 0).
		Render(message)
}
