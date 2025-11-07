//go:build wasm

package interfaces

type Controller interface {

	// Required for App Components
	Name() string

	// State Transition Methods
	Enter() bool
	Leave() bool

	// Controller Methods
	Update()
	Render()

}
