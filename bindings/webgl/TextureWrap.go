//go:build wasm

package webgl

type TextureWrap int

const (
	TextureWrapClampToEdge    TextureWrap = 0x812F
	TextureWrapMirroredRepeat TextureWrap = 0x8370
	TextureWrapRepeat         TextureWrap = 0x2901
)
