//go:build wasm

package interfaces

type Component interface {
	Render()
	String() string
}
