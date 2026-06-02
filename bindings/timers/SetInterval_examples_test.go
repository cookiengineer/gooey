//go:build wasm

package timers

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleSetInterval() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	interval_id := SetInterval(func() {
		console.Log("1 second has passed... again!")
	}, 1000 * Millisecond)

	if interval_id != 0 {

		setTimeout(func() {
			ClearInterval(interval_id)
		}, 10 * Second)

	}

}

