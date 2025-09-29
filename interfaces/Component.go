//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type Component interface {
	Enable()                 bool
	Disable()                bool

	Mount()                  bool
	Unmount()                bool

	Query(string)            Component

	Render()                 *dom.Element
	String()                 string

}
