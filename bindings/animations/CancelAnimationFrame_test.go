//go:build wasm

package animations

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleCancelAnimationFrame() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	identifier := RequestAnimationFrame(func(timestamp float64) {
		console.Log(timestamp)
	})

	CancelAnimationFrame(identifier)

}
