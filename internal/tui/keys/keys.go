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

func Input() key.Binding {
	return key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "chat"),
	)
}

func Escape() key.Binding {
	return key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "escape chat"),
	)
}

type KeyMap struct {
	bindings []key.Binding
}

func newKeyMap(bindings ...key.Binding) *KeyMap {
	return &KeyMap{
		bindings: bindings,
	}
}

func (k *KeyMap) ShortHelp() []key.Binding {
	return k.bindings
}

func (k *KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
