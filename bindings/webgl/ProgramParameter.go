//go:build wasm

package webgl

type ProgramParameter uint

const (
	ProgramParameterActiveAttributes ProgramParameter = 0x8B89
	ProgramParameterActiveUniforms   ProgramParameter = 0x8B86
	ProgramParameterAttachedShaders  ProgramParameter = 0x8B85
	ProgramParameterDeleteStatus     ProgramParameter = 0x8B80
	ProgramParameterLinkStatus       ProgramParameter = 0x8B82
	ProgramParameterValidateStatus   ProgramParameter = 0x8B83
)
