//go:build wasm

package webgl

// TextureFormat covers both the internal format and the format that
// TexImage2D() takes. WebGL2 accepts the unsized formats of WebGL1 as well as
// the sized ones, and both are given here, because a binding that only offered
// one of them would refuse valid calls.
type TextureFormat uint

const (
	TextureFormatAlpha             TextureFormat = 0x1906
	TextureFormatDepthComponent    TextureFormat = 0x1902
	TextureFormatDepthComponent16  TextureFormat = 0x81A5
	TextureFormatDepthComponent24  TextureFormat = 0x81A6
	TextureFormatDepthComponent32F TextureFormat = 0x8CAC
	TextureFormatDepthStencil      TextureFormat = 0x84F9
	TextureFormatDepth24Stencil8   TextureFormat = 0x88F0
	TextureFormatLuminance         TextureFormat = 0x1909
	TextureFormatLuminanceAlpha    TextureFormat = 0x190A
	TextureFormatR8                TextureFormat = 0x8229
	TextureFormatR32F              TextureFormat = 0x822E
	TextureFormatRed               TextureFormat = 0x1903
	TextureFormatRedInteger        TextureFormat = 0x8D94
	TextureFormatRG                TextureFormat = 0x8227
	TextureFormatRG8               TextureFormat = 0x822B
	TextureFormatRGB               TextureFormat = 0x1907
	TextureFormatRGB8              TextureFormat = 0x8051
	TextureFormatRGB16F            TextureFormat = 0x881B
	TextureFormatRGB32F            TextureFormat = 0x8815
	TextureFormatRGBA              TextureFormat = 0x1908
	TextureFormatRGBA8             TextureFormat = 0x8058
	TextureFormatRGBA16F           TextureFormat = 0x881A
	TextureFormatRGBA32F           TextureFormat = 0x8814
	TextureFormatRGBAInteger       TextureFormat = 0x8D99
	TextureFormatRGBInteger        TextureFormat = 0x8D98
	TextureFormatSRGB8             TextureFormat = 0x8C41
	TextureFormatSRGB8Alpha8       TextureFormat = 0x8C43
)
