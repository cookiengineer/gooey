//go:build wasm

package webgl

type TextureFilter int

const (
	TextureFilterLinear               TextureFilter = 0x2601
	TextureFilterLinearMipmapLinear   TextureFilter = 0x2703
	TextureFilterLinearMipmapNearest  TextureFilter = 0x2701
	TextureFilterNearest              TextureFilter = 0x2600
	TextureFilterNearestMipmapLinear  TextureFilter = 0x2702
	TextureFilterNearestMipmapNearest TextureFilter = 0x2700
)
