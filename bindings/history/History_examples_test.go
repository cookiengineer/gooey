//go:build wasm

package history

import "github.com/cookiengineer/gooey/bindings/console"
import "time"

func Example() {

	console := console.GetConsole()
	history := GetHistory()

	history.AddEventListener(ToEventListener(func(event *PopStateEvent) {
		console.Log("PoppState Event")
		console.Log(event)
	}))

	history.PushState(&map[string]any{"page": 1}, "first page",  "/page-1.html")
	history.PushState(&map[string]any{"page": 2}, "second page", "/page-2.html")
	history.PushState(&map[string]any{"page": 3}, "third page",  "/page-3.html")

	history.Back()
	history.Back()

	go func() {

		time.Sleep(3 * time.Second)
		history.Forward()

		time.Sleep(3 * time.Second)
		history.Forward()

	}()

}
