//go:build wasm

package crypto

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleGetRandomValues() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	random  := GetRandomValues(1337)

	console.Log(random)

}

