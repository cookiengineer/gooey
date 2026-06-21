//go:build wasm

package bindings

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleWindow_AddEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console := console.GetConsole()
	window := GetWindow()

	window.AddEventListener("resize", dom.ToEventListener(func(event *dom.Event) {
		console.Log(event.Target)
	}))

}

func ExampleWindow_RemoveEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console := console.GetConsole()
	window := GetWindow()

	event_listener := dom.ToEventListener(func(event *dom.Event) {
		console.Log(event.Target)
	})

	window.AddEventListener("resize", event_listener)

	// Remove specified event listener
	window.RemoveEventListener("resize", event_listener) // returns true
	window.RemoveEventListener("resize", event_listener) // returns false, already removed

	// Remove all event listeners
	window.RemoveEventListener("resize", nil) // returns true

}

func ExampleWindow_Confirm() {

	window := GetWindow()

	if window.Confirm("Are you at least 5 years old?") == false {
		window.Close()
	}

}

func ExampleWindow_MoveBy() {

	window := GetWindow()
	window.MoveBy(13, 37)

}

func ExampleWindow_MoveTo() {

	window := GetWindow()
	window.MoveTo(13, 37)

}
