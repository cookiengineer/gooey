//go:build wasm

package bindings

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleScreen_AddEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console := console.GetConsole()
	window := GetWindow()

	window.Screen.AddEventListener("change", dom.ToEventListener(func(event *dom.Event) {
		console.Log(event)
	}))

}

func ExampleScreen_RemoveEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console := console.GetConsole()
	window := GetWindow()

	event_listener := dom.ToEventListener(func(event *dom.Event) {
		console.Log(event)
	})

	window.Screen.AddEventListener("change", event_listener)

	// Remove specified event listener
	window.Screen.RemoveEventListener("change", event_listener) // returns true
	window.Screen.RemoveEventListener("change", event_listener) // returns false, already removed

	// Remove all event listeners
	window.Screen.RemoveEventListener("change", nil) // returns true

}
