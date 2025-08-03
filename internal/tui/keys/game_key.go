package keys

import "github.com/charmbracelet/bubbles/key"

var gameKeys = newKeyMap(Up(), Down(), Left(), Right(), Enter(), Input(), Escape(), Quit())

func GetGameKeys() gameKeyMap {
	return gameKeyMap{
		KeyMap: gameKeys,
	}
}

// basicKeyMap is just a wrapper around KeyMap pointer
type gameKeyMap struct {
	*KeyMap
}

func (k gameKeyMap) Up() key.Binding {
	return k.bindings[0]
}

func (k gameKeyMap) Down() key.Binding {
	return k.bindings[1]
}

func (k gameKeyMap) Left() key.Binding {
	return k.bindings[2]
}

func (k gameKeyMap) Right() key.Binding {
	return k.bindings[3]
}

func (k gameKeyMap) Enter() key.Binding {
	return k.bindings[4]
}

func (k gameKeyMap) Input() key.Binding {
	return k.bindings[5]
}

func (k gameKeyMap) Escape() key.Binding {
	return k.bindings[6]
}

func (k gameKeyMap) Quit() key.Binding {
	return k.bindings[7]
}
