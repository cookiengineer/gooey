//go:build wasm

package dom

import "github.com/cookiengineer/gooey/bindings/console"

func Example() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console  := console.GetConsole()
	document := GetDocument()
	element  := document.CreateElement("hello-world")

	element.SetAttribute("data-example", "my-value")
	element.SetInnerHTML("Hello, world!<br>Please click me!")

	element.AddEventListener(EventTypeClick, ToEventListener(func(event *Event) {
		console.Log("Click Event")
		console.Log(event)
	}))

	document.Body.Append(element)

}

func ExampleDocument_AddEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console  := console.GetConsole()
	document := GetDocument()
	document.AddEventListener("click", ToEventListener(func(event *Event) {
		console.Log(event.Target)
	}))

}

func ExampleDocument_QuerySelector() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console  := console.GetConsole()
	document := GetDocument()
	element  := document.QuerySelector("body")

	console.Log(element.TagName)

}

func ExampleDocument_QuerySelectorAll() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console  := console.GetConsole()
	document := GetDocument()
	elements := document.QuerySelectorAll("div")

	for _, element := range elements {
		console.Log(element.TagName)
	}

}

func ExampleDocument_RemoveEventListener() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console  := console.GetConsole()
	document := GetDocument()

	event_listener := ToEventListener(func(event *Event) {
		console.Log(event.Target)
	})

	document.AddEventListener("click", event_listener)

	// Remove specified event listener
	document.RemoveEventListener("click", event_listener) // returns true
	document.RemoveEventListener("click", event_listener) // returns false, already removed

	// Remove all event listeners
	document.RemoveEventListener("click", nil) // returns true

}

