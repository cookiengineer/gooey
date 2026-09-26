//go:build wasm

package webgl

type TextureParameter uint

const (
	TextureParameterBaseLevel   TextureParameter = 0x813C
	TextureParameterCompareFunc TextureParameter = 0x884D
	TextureParameterCompareMode TextureParameter = 0x884C
	TextureParameterMagFilter   TextureParameter = 0x2800
	TextureParameterMaxLevel    TextureParameter = 0x813D
	TextureParameterMaxLOD      TextureParameter = 0x813B
	TextureParameterMinFilter   TextureParameter = 0x2801
	TextureParameterMinLOD      TextureParameter = 0x813A
	TextureParameterWrapR       TextureParameter = 0x8072
	TextureParameterWrapS       TextureParameter = 0x2802
	TextureParameterWrapT       TextureParameter = 0x2803
)
