//go:build wasm

// Package history provides bindings for the History interface.
//
// Specification: https://html.spec.whatwg.org/multipage/nav-history-apis.html#the-history-interface
//
// Example usage:
//
//	import "github.com/cookiengineer/gooey/bindings/console"
//	import "github.com/cookiengineer/gooey/bindings/history"
//	import "time"
//
//	console := console.GetConsole()
//	history := history.GetHistory()
//
//	history.AddEventListener(history.ToEventListener(func(event *history.PopStateEvent) {
//		console.Log("PoppState Event")
//		console.Log(event)
//	}))
//
//	history.PushState(&map[string]any{"page": 1}, "first page",  "/page-1.html")
//	history.PushState(&map[string]any{"page": 2}, "second page", "/page-2.html")
//	history.PushState(&map[string]any{"page": 3}, "third page",  "/page-3.html")
//
//	history.Back()
//	history.Back()
//
//	go func() {
//
//		time.Sleep(3 * time.Second)
//		history.Forward()
//
//		time.Sleep(3 * time.Second)
//		history.Forward()
//
//	}()
//
package history
