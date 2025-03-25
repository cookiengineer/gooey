//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type View interface {
	Init(string, string, string)
	GetElement(string) *dom.Element
	SetElement(string, *dom.Element)
	RemoveElement(string) bool
	Enter() bool
	Leave() bool
	Render()

	// Required for UI integration
	Properties() (string, string, string)
}

