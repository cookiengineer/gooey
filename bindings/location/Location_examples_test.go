//go:build wasm

package location

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleLocation_Assign() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	location := GetLocation()

	err := location.Assign("https://example.com/path/to/file.html")

	if err == nil {
		console.Info("Window URL assigned!")
	} else {
		// SecurityError or SyntaxError
		console.Error(err)
	}

}

func ExampleLocation_Reload() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	location := GetLocation()

	err := location.Reload()

	if err == nil {
		console.Info("Window reloaded!")
	} else {
		// SecurityError
		console.Error(err)
	}

}

func ExampleLocation_Replace() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	location := GetLocation()

	err := location.Replace("https://example.com/path/to/file.html")

	if err == nil {
		console.Info("Window URL replaced!")
	} else {
		// SecurityError or SyntaxError
		console.Error(err)
	}

}

