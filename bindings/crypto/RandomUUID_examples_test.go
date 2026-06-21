//go:build wasm

package crypto

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleRandomUUID() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	uuid    := RandomUUID()

	console.Log(uuid)

}

