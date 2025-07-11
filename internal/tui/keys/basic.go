package keys

import "github.com/charmbracelet/bubbles/key"

var basicKeys = newKeyMap(Up(), Down(), Enter(), Quit())

func GetBasicKeys() basicKeyMap {
	return basicKeyMap{
		KeyMap: basicKeys,
	}
}

// basicKeyMap is just a wrapper around KeyMap pointer
type basicKeyMap struct {
	*KeyMap
}

func (k basicKeyMap) Up() key.Binding {
	return k.bindings[0]
}

func (k basicKeyMap) Down() key.Binding {
	return k.bindings[1]
}

func (k basicKeyMap) Enter() key.Binding {
	return k.bindings[2]
}

func (k basicKeyMap) Quit() key.Binding {
	return k.bindings[3]
}
