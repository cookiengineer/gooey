//go:build wasm

package interfaces

// Scheduler coalesces invalidated components and renders them in batches.
type Scheduler interface {

	// Schedule marks a component as needing a render.
	Schedule(component Component)

	// Unschedule removes a component that is no longer mounted.
	Unschedule(component Component)

}

