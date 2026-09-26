//go:build wasm

package webgl

type BufferUsage uint

const (
	BufferUsageDynamicCopy BufferUsage = 0x88EA
	BufferUsageDynamicDraw BufferUsage = 0x88E8
	BufferUsageDynamicRead BufferUsage = 0x88E9
	BufferUsageStaticCopy  BufferUsage = 0x88E6
	BufferUsageStaticDraw  BufferUsage = 0x88E4
	BufferUsageStaticRead  BufferUsage = 0x88E5
	BufferUsageStreamCopy  BufferUsage = 0x88E2
	BufferUsageStreamDraw  BufferUsage = 0x88E0
	BufferUsageStreamRead  BufferUsage = 0x88E1
)
