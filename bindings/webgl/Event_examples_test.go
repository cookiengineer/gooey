//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleEvent_PreventDefault() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)

	canvas.AddEventListener(EventTypeContextLost, ToEventListener(func(event *Event) {

		// without this, the Web Browser Engine never restores the Context
		event.PreventDefault()

		console.Log("Context Lost Event")

	}))

	canvas.AddEventListener(EventTypeContextRestored, ToEventListener(func(event *Event) {
		console.Log("Context Restored Event")
	}))

}
