//go:build wasm

package webgl

type FramebufferTarget uint

const (
	FramebufferTargetDraw        FramebufferTarget = 0x8CA9
	FramebufferTargetFramebuffer FramebufferTarget = 0x8D40
	FramebufferTargetRead        FramebufferTarget = 0x8CA8
)
