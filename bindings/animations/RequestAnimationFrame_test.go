//go:build wasm

package animations

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleRequestAnimationFrame() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()

	identifier := RequestAnimationFrame(func(timestamp float64) {
		console.Log(timestamp)
	})

	console.Log(identifier)

}
