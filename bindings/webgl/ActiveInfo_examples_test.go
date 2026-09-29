//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleActiveInfo() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()
	program  := context.CreateProgram()

	amount := context.GetProgramParameterInt(program, ProgramParameterActiveUniforms)

	for a := 0; a < amount; a++ {

		info := context.GetActiveUniform(program, uint(a))

		if info != nil {
			console.Log(info.Name)
		}

	}

}
