//go:build wasm

package webgl

type PixelStoreParameter uint

const (
	PixelStoreParameterPackAlignment              PixelStoreParameter = 0x0D05
	PixelStoreParameterUnpackAlignment            PixelStoreParameter = 0x0CF5
	PixelStoreParameterUnpackColorspaceConversion PixelStoreParameter = 0x9243
	PixelStoreParameterUnpackFlipY                PixelStoreParameter = 0x9240
	PixelStoreParameterUnpackPremultiplyAlpha     PixelStoreParameter = 0x9241
	PixelStoreParameterUnpackRowLength            PixelStoreParameter = 0x0CF2
)
