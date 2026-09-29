//go:build wasm

package interfaces

import "github.com/cookiengineer/gooey/bindings/dom"

type Component interface {

	// Lifecycle Methods
	Enable()  bool
	Disable() bool
	Mount()   bool
	Unmount() bool

	// Scheduler Propagation
	SetScheduler(scheduler Scheduler)

	// Component Methods
	Query(string) Component
	Render()      *dom.Element
	String()      string

}
