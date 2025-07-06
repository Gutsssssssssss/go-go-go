package page

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
)

type PageID string

type PagePushMsg struct {
	ID PageID
}

type PagePopMsg struct{}

type Page struct {
	tea.Model
	layout.Sizable
	ID     PageID
	help   help.Model
	window struct {
		width, height int
	}
}

func (p *Page) SetSize(width, height int) {
	p.window.width = width
	p.window.height = height
}

func (p *Page) GetSize() (int, int) {
	return p.window.width, p.window.height
}
