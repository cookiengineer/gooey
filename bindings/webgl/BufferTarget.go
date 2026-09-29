//go:build wasm

package webgl

type BufferTarget uint

const (
	BufferTargetArray             BufferTarget = 0x8892
	BufferTargetCopyRead          BufferTarget = 0x8F36
	BufferTargetCopyWrite         BufferTarget = 0x8F37
	BufferTargetElementArray      BufferTarget = 0x8893
	BufferTargetPixelPack         BufferTarget = 0x88EB
	BufferTargetPixelUnpack       BufferTarget = 0x88EC
	BufferTargetTransformFeedback BufferTarget = 0x8C8E
	BufferTargetUniform           BufferTarget = 0x8A11
)
