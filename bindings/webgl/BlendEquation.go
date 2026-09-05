//go:build wasm

package webgl

type BlendEquation uint

const (
	BlendEquationAdd             BlendEquation = 0x8006
	BlendEquationMax             BlendEquation = 0x8008
	BlendEquationMin             BlendEquation = 0x8007
	BlendEquationReverseSubtract BlendEquation = 0x800B
	BlendEquationSubtract        BlendEquation = 0x800A
)
