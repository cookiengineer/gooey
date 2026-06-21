//go:build wasm

package timers

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleSetTimeout() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()

	SetTimeout(func() {
		console.Log("13 milliseconds have passed!")
	}, 13 * Millisecond)

	SetTimeout(func() {
		console.Log("37 seconds have passed!")
	}, 37 * Second)

	clearable_timeout_id := SetTimeout(func() {
		console.Log("1 minute has passed!")
	}, 1 * Minute)

	if clearable_timeout_id != 0 {

		SetTimeout(func() {
			// This will cancel the timeout callback from being fired.
			console.Warn("Cleared last timeout callback!")
			ClearTimeout(clearable_timeout_id)
		}, 20 * Second)

	}

}
