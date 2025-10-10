//go:build wasm

package interfaces

type Controller interface {

	// Required for App Components
	Name() string

	// Controller Methods
	Update()
	Render()

}
