//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type View interface {

	// Required for App Components
	Name()  string
	Label() string
	Path()  string

	// Lifecycle Methods
	Enable()  bool
	Disable() bool
	Mount()   bool
	Unmount() bool
	Enter()   bool
	Leave()   bool

	// Component Methods
	Query(string) Component
	Render()      *dom.Element
	String()      string

}

