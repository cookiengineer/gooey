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
//	bindings.Window.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {
//	  console.Log(event)
//	}))
//
//	console.Log(bindings.Window)
//	console.Log(bindings.Window.Screen)
//
// The global Window object is instantiated via init() and the properties are automatically
// integrated with their respective events.
//
//	import "github.com/cookiengineer/gooey/bindings"
//	import "github.com/cookiengineer/gooey/bindings/console"
//
//	// The Window's InnerWidth/InnerHeight properties are automatically updated
//	bindings.Window.AddEventListener("resize", dom.ToEventListener(func(event *dom.Event) {
//	  console.Log(bindings.Window.InnerWidth)
//	  console.Log(bindings.Window.InnerHeight)
//	}))
//
//	// The Window's ScrollX/ScrollY properties are automatically updated
//	bindings.Window.AddEventListener("scroll", dom.ToEventListener(func(event *dom.Event) {
//	  console.Log(bindings.Window.ScrollX)
//	  console.Log(bindings.Window.ScrollY)
//	}))
//
//	// The Screen.Orientation's Angle/Type properties are automatically updated
//	bindings.Window.Screen.AddEventListener("change", dom.ToEventListener(func(event *dom.Event) {
//	  console.Log(bindings.Window.Screen.Orientation.Angle)
//	  console.Log(bindings.Window.Screen.Orientation.Type)
//	}))
package bindings
