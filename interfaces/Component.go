//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type Component interface {
	Enable()  bool
	Disable() bool
	Render()  *dom.Element
	String()  string
}
