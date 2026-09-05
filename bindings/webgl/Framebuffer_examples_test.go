//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleFramebuffer() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	target := context.CreateTexture()

	context.BindTexture(TextureTarget2D, target)

	// a nil pixel slice allocates the storage without filling it
	context.TexImage2D(TextureTarget2D, 0, TextureFormatRGBA, 512, 512, TextureFormatRGBA, DataTypeUnsignedByte, nil)

	context.TexParameterFilter(TextureTarget2D, TextureParameterMagFilter, TextureFilterNearest)
	context.TexParameterFilter(TextureTarget2D, TextureParameterMinFilter, TextureFilterNearest)

	framebuffer := context.CreateFramebuffer()

	context.BindFramebuffer(FramebufferTargetFramebuffer, framebuffer)
	context.FramebufferTexture2D(FramebufferTargetFramebuffer, FramebufferAttachmentColor0, TextureTarget2D, target, 0)

	status := context.CheckFramebufferStatus(FramebufferTargetFramebuffer)

	if status == FramebufferStatusComplete {

		context.Viewport(0, 0, 512, 512)
		context.ClearColor(0.0, 0.0, 0.0, 1.0)
		context.Clear(BufferBitColor | BufferBitDepth)

		pixels := context.ReadPixels(0, 0, 512, 512, TextureFormatRGBA, DataTypeUnsignedByte)

		console.Log(len(pixels))

	}

	context.BindFramebuffer(FramebufferTargetFramebuffer, nil)

}
