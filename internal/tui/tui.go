package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yanmoyy/go-go-go/internal/tui/layout"
	"github.com/yanmoyy/go-go-go/internal/tui/page"
	"github.com/yanmoyy/go-go-go/internal/util/ds"
)

type appModel struct {
	pages     map[page.PageID]tea.Model
	pageStack *ds.Stack[page.PageID]
	window    struct {
		width, height int
	}
}

func New() *appModel {
	initialPage := page.StartPage
	stack := ds.NewStack[page.PageID]()
	stack.Push(initialPage)
	return &appModel{
		pageStack: stack,
		pages: map[page.PageID]tea.Model{
			page.StartPage: page.NewStartPage(),
			page.LobbyPage: page.NewLobbyPage(),
		},
	}
}

func (a *appModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmd := a.pages[a.pageStack.Top()].Init()
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (a *appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.window.width = msg.Width
		a.window.height = msg.Height
		if sizable, ok := a.pages[a.pageStack.Top()].(layout.Sizable); ok {
			sizable.SetSize(msg.Width, msg.Height)
			cmds = append(cmds, cmd)
		}
	case page.PagePushMsg:
		return a, a.pushPage(msg.ID)
	case page.PagePopMsg:
		return a, a.popPage()
	}
	currentPage := a.pageStack.Top()
	a.pages[currentPage], cmd = a.pages[currentPage].Update(msg)
	cmds = append(cmds, cmd)
	return a, tea.Batch(cmds...)
}

func (a *appModel) View() string {
	return a.pages[a.pageStack.Top()].View()
}

func (a *appModel) pushPage(pageID page.PageID) tea.Cmd {
	var cmds []tea.Cmd
	a.pageStack.Push(pageID)
	if sizable, ok := a.pages[pageID].(layout.Sizable); ok {
		sizable.SetSize(a.window.width, a.window.height)
	}
	cmds = append(cmds, a.pages[pageID].Init())
	return tea.Batch(cmds...)
}

func (a *appModel) popPage() tea.Cmd {
	if a.pageStack.Len() > 1 {
		_, _ = a.pageStack.Pop()
		currentPage := a.pageStack.Top()
		if sizable, ok := a.pages[currentPage].(layout.Sizable); ok {
			sizable.SetSize(a.window.width, a.window.height)
		}
	}
	return nil
}
