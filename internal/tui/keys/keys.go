package keys

import "github.com/charmbracelet/bubbles/key"

func Up() key.Binding {
	return key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	)
}

func Down() key.Binding {
	return key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	)
}

func Right() key.Binding {
	return key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	)
}

func Left() key.Binding {
	return key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	)
}

func Enter() key.Binding {
	return key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter"),
	)
}

func Help() key.Binding {
	return key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	)
}

func Quit() key.Binding {
	return key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	)
}

func Back() key.Binding {
	return key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	)
}

type DefaultKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Right key.Binding
	Left  key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
	Back  key.Binding
}

func GetDefaultKeyMap() DefaultKeyMap {
	return DefaultKeyMap{
		Up:    Up(),
		Down:  Down(),
		Right: Right(),
		Left:  Left(),
		Enter: Enter(),
		Help:  Help(),
		Quit:  Quit(),
		Back:  Back(),
	}
}

func (k DefaultKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Left, k.Right, k.Enter, k.Help, k.Quit}
}

func (k DefaultKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}
