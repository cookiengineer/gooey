//go:build wasm

package webgl

type ShaderParameter uint

const (
	ShaderParameterCompileStatus ShaderParameter = 0x8B81
	ShaderParameterDeleteStatus  ShaderParameter = 0x8B80
	ShaderParameterShaderType    ShaderParameter = 0x8B4F
)
