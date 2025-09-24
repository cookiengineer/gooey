//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type Component interface {
	Enable()                 bool
	Disable()                bool

	Query(string)            Component
	SetChildren([]Component) bool

	Mount()                  bool
	Unmount()                bool
	Render()                 *dom.Element
	String()                 string
}
