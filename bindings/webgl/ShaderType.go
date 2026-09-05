//go:build wasm

package webgl

type ShaderType uint

const (
	ShaderTypeFragment ShaderType = 0x8B30
	ShaderTypeVertex   ShaderType = 0x8B31
)
