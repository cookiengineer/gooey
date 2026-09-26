//go:build wasm

package webgl

type BlendFactor uint

const (
	BlendFactorConstantAlpha         BlendFactor = 0x8003
	BlendFactorConstantColor         BlendFactor = 0x8001
	BlendFactorDstAlpha              BlendFactor = 0x0304
	BlendFactorDstColor              BlendFactor = 0x0306
	BlendFactorOne                   BlendFactor = 1
	BlendFactorOneMinusConstantAlpha BlendFactor = 0x8004
	BlendFactorOneMinusConstantColor BlendFactor = 0x8002
	BlendFactorOneMinusDstAlpha      BlendFactor = 0x0305
	BlendFactorOneMinusDstColor      BlendFactor = 0x0307
	BlendFactorOneMinusSrcAlpha      BlendFactor = 0x0303
	BlendFactorOneMinusSrcColor      BlendFactor = 0x0301
	BlendFactorSrcAlpha              BlendFactor = 0x0302
	BlendFactorSrcAlphaSaturate      BlendFactor = 0x0308
	BlendFactorSrcColor              BlendFactor = 0x0300
	BlendFactorZero                  BlendFactor = 0
)
