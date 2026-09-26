//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleRenderbuffer() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	framebuffer  := context.CreateFramebuffer()
	renderbuffer := context.CreateRenderbuffer()

	context.BindRenderbuffer(renderbuffer)
	context.RenderbufferStorage(RenderbufferFormatDepthComponent16, 512, 512)

	context.BindFramebuffer(FramebufferTargetFramebuffer, framebuffer)
	context.FramebufferRenderbuffer(FramebufferTargetFramebuffer, FramebufferAttachmentDepth, renderbuffer)

	status := context.CheckFramebufferStatus(FramebufferTargetFramebuffer)

	if status != FramebufferStatusComplete {
		console.Error(status.String())
	}

	context.BindFramebuffer(FramebufferTargetFramebuffer, nil)

}
