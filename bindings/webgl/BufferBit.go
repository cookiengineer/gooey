//go:build wasm

package webgl

type BufferBit uint

const (
	BufferBitColor   BufferBit = 0x00004000
	BufferBitDepth   BufferBit = 0x00000100
	BufferBitStencil BufferBit = 0x00000400
)
