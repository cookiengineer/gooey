//go:build wasm

package webgl

type RenderbufferFormat uint

const (
	RenderbufferFormatDepthComponent16 RenderbufferFormat = 0x81A5
	RenderbufferFormatDepthComponent24 RenderbufferFormat = 0x81A6
	RenderbufferFormatDepth24Stencil8  RenderbufferFormat = 0x88F0
	RenderbufferFormatR8               RenderbufferFormat = 0x8229
	RenderbufferFormatRGB565           RenderbufferFormat = 0x8D62
	RenderbufferFormatRGB5A1           RenderbufferFormat = 0x8057
	RenderbufferFormatRGBA4            RenderbufferFormat = 0x8056
	RenderbufferFormatRGBA8            RenderbufferFormat = 0x8058
	RenderbufferFormatStencilIndex8    RenderbufferFormat = 0x8D48
)
