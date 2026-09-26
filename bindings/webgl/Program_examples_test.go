//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleProgram() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "github.com/cookiengineer/gooey/bindings/dom"

	console  := console.GetConsole()
	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	vertex_source := `#version 300 es

in vec3 position;

void main() {
	gl_Position = vec4(position, 1.0);
}
`

	fragment_source := `#version 300 es

precision highp float;

uniform vec4 tint;
out vec4 color;

void main() {
	color = tint;
}
`

	vertex_shader   := context.CreateShader(ShaderTypeVertex)
	fragment_shader := context.CreateShader(ShaderTypeFragment)

	context.ShaderSource(vertex_shader, vertex_source)
	context.ShaderSource(fragment_shader, fragment_source)

	err1 := context.CompileShader(vertex_shader)
	err2 := context.CompileShader(fragment_shader)

	if err1 != nil || err2 != nil {
		return
	}

	program := context.CreateProgram()

	context.AttachShader(program, vertex_shader)
	context.AttachShader(program, fragment_shader)

	err3 := context.LinkProgram(program)

	if err3 != nil {
		console.Error(err3)
		return
	}

	context.UseProgram(program)

	context.DeleteShader(vertex_shader)
	context.DeleteShader(fragment_shader)

}
