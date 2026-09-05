//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleTexture() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	// a 2x2 texture, one RGBA pixel per row of four bytes
	pixels := []byte{
		255, 0, 0, 255,
		0, 255, 0, 255,
		0, 0, 255, 255,
		255, 255, 255, 255,
	}

	texture := context.CreateTexture()

	context.ActiveTexture(TextureUnit0)
	context.BindTexture(TextureTarget2D, texture)

	// a 2x2 RGBA row is 8 bytes wide, so the default 4-byte row alignment
	// would read two bytes of padding that are not there
	context.PixelStoreInt(PixelStoreParameterUnpackAlignment, 1)

	err := context.TexImage2D(TextureTarget2D, 0, TextureFormatRGBA, 2, 2, TextureFormatRGBA, DataTypeUnsignedByte, pixels)

	if err != nil {
		return
	}

	context.TexParameterFilter(TextureTarget2D, TextureParameterMagFilter, TextureFilterNearest)
	context.TexParameterFilter(TextureTarget2D, TextureParameterMinFilter, TextureFilterNearest)
	context.TexParameterWrap(TextureTarget2D, TextureParameterWrapS, TextureWrapClampToEdge)
	context.TexParameterWrap(TextureTarget2D, TextureParameterWrapT, TextureWrapClampToEdge)

	context.BindTexture(TextureTarget2D, nil)
	context.DeleteTexture(texture)

}
