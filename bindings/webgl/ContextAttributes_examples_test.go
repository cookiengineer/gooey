//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleContextAttributes_ToValue() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")

	attributes := NewContextAttributes()
	attributes.Antialias = false
	attributes.Depth = true
	attributes.PowerPreference = PowerPreferenceHighPerformance
	attributes.PreserveDrawingBuffer = true

	canvas := ToCanvasWithAttributes(element, &attributes)

	if canvas != nil {
		canvas.GetContext().Viewport(0, 0, int(canvas.Width), int(canvas.Height))
	}

}
