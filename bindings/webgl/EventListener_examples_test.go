//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)

	listener := ToEventListener(func(event *Event) {
		console.Log(event.StatusMessage)
	})

	canvas.AddEventListener(EventTypeContextCreationError, listener)
	canvas.RemoveEventListener(EventTypeContextCreationError, listener)

}
