//go:build wasm

package interfaces

type View interface {
	Enter() bool
	Leave() bool
	Render()

	// Required for UI integration
	GetProperty(string) string
	SetProperty(string, string) bool
}

