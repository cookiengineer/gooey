//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleUniformLocation() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()
	program  := context.CreateProgram()

	// column major, as the specification requires
	projection := []float32{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	context.UseProgram(program)

	tint := context.GetUniformLocation(program, "tint")
	matrix := context.GetUniformLocation(program, "projection")

	context.Uniform4f(tint, 1.0, 0.0, 0.0, 1.0)
	context.UniformMatrix4fv(matrix, projection)

}
