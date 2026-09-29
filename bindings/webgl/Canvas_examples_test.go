//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleCanvas_GetContext() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	context.ClearColor(0.0, 0.0, 0.0, 1.0)
	context.Clear(BufferBitColor | BufferBitDepth)

}

func ExampleCanvas_SetWidth() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)

	canvas.SetWidth(800)
	canvas.SetHeight(600)

}
