//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleContext_Clear() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	context.Viewport(0, 0, int(canvas.Width), int(canvas.Height))
	context.ClearColor(0.1, 0.1, 0.15, 1.0)
	context.ClearDepth(1.0)
	context.Clear(BufferBitColor | BufferBitDepth)

}

func ExampleContext_Enable() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	context.Enable(CapabilityDepthTest)
	context.DepthFunc(DepthFuncLEqual)
	context.DepthMask(true)

	context.Enable(CapabilityCullFace)
	context.CullFace(CullFaceBack)
	context.FrontFace(FrontFaceCounterClockwise)

	context.Enable(CapabilityBlend)
	context.BlendFunc(BlendFactorSrcAlpha, BlendFactorOneMinusSrcAlpha)

}

func ExampleContext_GetError() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	context.Enable(CapabilityDepthTest)

	code := context.GetError()

	if code != ErrorCodeNoError {
		console.Error(code.String())
	}

}

func ExampleContext_GetParameterString() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	console.Log(context.GetParameterString(ParameterVersion))
	console.Log(context.GetParameterInt(ParameterMaxTextureSize))

}

func ExampleContext_ReadPixels() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	context.ClearColor(0.0, 0.0, 0.0, 1.0)
	context.Clear(BufferBitColor)

	// four components of one byte each, so a row of 64 pixels is already a
	// multiple of the default 4-byte pack alignment and needs no padding
	pixels := context.ReadPixels(0, 0, 64, 64, TextureFormatRGBA, DataTypeUnsignedByte)

	console.Log(len(pixels))

	// a floating point read needs a floating point framebuffer AND the matching
	// type; the bytes come back raw, four per component, to be decoded by the
	// caller
	floats := context.ReadPixels(0, 0, 64, 64, TextureFormatRGBA, DataTypeFloat)

	console.Log(len(floats))

}

func ExampleContext_TexImage2D() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()
	texture  := context.CreateTexture()

	// one 16-bit depth value per pixel, so the bytes are read through a
	// Uint16Array rather than a Uint8Array
	depth := make([]byte, 64*64*2)

	context.BindTexture(TextureTarget2D, texture)

	err := context.TexImage2D(
		TextureTarget2D,
		0,
		TextureFormatDepthComponent16,
		64,
		64,
		TextureFormatDepthComponent,
		DataTypeUnsignedShort,
		depth,
	)

	if err != nil {
		console.Error(err)
	}

}
