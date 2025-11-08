//go:build wasm

// Package bindings provides high-level API bindings for the global objects in the Browser
// such as Window, Screen, and ScreenOrientation.
//
// Example usage:
//
//	import "github.com/cookiengineer/gooey/bindings"
//	import "github.com/cookiengineer/gooey/bindings/console"
//	import "github.com/cookiengineer/gooey/bindings/dom"
//
//	console := console.GetConsole()
//	window := bindings.GetWindow()
//
//	window.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {
//		console1.Log(event)
//	}))
//
//	console.Log(window1)
//	console.Log(window1.Screen)
//
// The global Window object is instantiated via init() and the properties are automatically
// integrated with their respective events.
//
//	import "github.com/cookiengineer/gooey/bindings"
//	import "github.com/cookiengineer/gooey/bindings/console"
//
//	console := console.GetConsole()
//	window := bindings.GetWindow()
//
//	// The Window's InnerWidth/InnerHeight properties are automatically updated
//	window.AddEventListener("resize", dom.ToEventListener(func(event *dom.Event) {
//		console.Log(window1.InnerWidth)
//		console.Log(window1.InnerHeight)
//	}))
//
//	// The Window's ScrollX/ScrollY properties are automatically updated
//	window.AddEventListener("scroll", dom.ToEventListener(func(event *dom.Event) {
//		console.Log(window1.ScrollX)
//		console.Log(window1.ScrollY)
//	}))
//
//	// The Screen.Orientation's Angle/Type properties are automatically updated
//	window.Screen.AddEventListener("change", dom.ToEventListener(func(event *dom.Event) {
//		console.Log(window1.Screen.Orientation.Angle)
//		console.Log(window1.Screen.Orientation.Type)
//	}))
package bindings
