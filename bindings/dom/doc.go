//go:build wasm

// Package dom provides bindings for the WHATWG Document Object Model Living Standard
//
// Specification: https://dom.spec.whatwg.org/
//
// Example usage:
//
//	import "github.com/cookiengineer/gooey/bindings/console"
//	import "github.com/cookiengineer/gooey/bindings/dom"
//
//	document := dom.GetDocument()
//	console  := console.GetConsole()
//	element  := document.CreateElement("hello-world")
//
//	element.SetAttribute("data-example", "my-value")
//	element.SetInnerHTML("Hello, world!<br>Please click me!")
//
//	element.AddEventListener(dom.EventTypeClick, dom.ToEventListener(func(event *dom.Event) {
//		console.Log("Click Event")
//		console.Log(event)
//	}))
//
//	document.Body.Append(element)
//
package dom
