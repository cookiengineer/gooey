//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleShader() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	source := `#version 300 es

in vec3 position;

void main() {
	gl_Position = vec4(position, 1.0);
}
`

	shader := context.CreateShader(ShaderTypeVertex)

	context.ShaderSource(shader, source)

	err := context.CompileShader(shader)

	if err != nil {
		console.Error(err)
	}

}
