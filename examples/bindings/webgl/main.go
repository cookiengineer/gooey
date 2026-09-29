package main

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/webgl"
import "time"

var vertex_source = `#version 300 es

in vec2 position;
in vec3 color;

out vec3 varying_color;

void main() {
	varying_color = color;
	gl_Position = vec4(position, 0.0, 1.0);
}
`

var fragment_source = `#version 300 es

precision highp float;

in vec3 varying_color;

uniform float alpha;

out vec4 fragment_color;

void main() {
	fragment_color = vec4(varying_color, alpha);
}
`

func main() {

	the_console := console.GetConsole()
	document := dom.GetDocument()
	element := document.QuerySelector("canvas")

	attributes := webgl.NewContextAttributes()
	attributes.Antialias = false
	attributes.PowerPreference = webgl.PowerPreferenceHighPerformance

	// A demo convenience, not a recommendation: it keeps the drawn frame
	// readable after compositing, which is what toDataURL(), drawImage() and a
	// screenshot need, and it costs a copy per frame on a real GPU. A renderer
	// that does not screenshot itself should leave this off and read the frame
	// back inside the RequestAnimationFrame callback instead.
	attributes.PreserveDrawingBuffer = true

	canvas := webgl.ToCanvasWithAttributes(element, &attributes)

	if canvas == nil || canvas.GetContext() == nil {
		the_console.Error("WebGL2 is not available")
		return
	}

	context := canvas.GetContext()

	the_console.Log(context.Version)
	the_console.Log(context.Renderer)

	canvas.AddEventListener(webgl.EventTypeContextLost, webgl.ToEventListener(func(event *webgl.Event) {

		// without this, the Web Browser Engine never restores the Context
		event.PreventDefault()

		the_console.Warn("Context Lost Event")

	}))

	canvas.AddEventListener(webgl.EventTypeContextRestored, webgl.ToEventListener(func(event *webgl.Event) {
		the_console.Log("Context Restored Event")
	}))

	vertex_shader := context.CreateShader(webgl.ShaderTypeVertex)
	context.ShaderSource(vertex_shader, vertex_source)

	err1 := context.CompileShader(vertex_shader)

	if err1 != nil {
		the_console.Error(err1)
		return
	}

	fragment_shader := context.CreateShader(webgl.ShaderTypeFragment)
	context.ShaderSource(fragment_shader, fragment_source)

	err2 := context.CompileShader(fragment_shader)

	if err2 != nil {
		the_console.Error(err2)
		return
	}

	program := context.CreateProgram()

	context.AttachShader(program, vertex_shader)
	context.AttachShader(program, fragment_shader)
	context.BindAttribLocation(program, 0, "position")
	context.BindAttribLocation(program, 1, "color")

	err3 := context.LinkProgram(program)

	if err3 != nil {
		the_console.Error(err3)
		return
	}

	context.DeleteShader(vertex_shader)
	context.DeleteShader(fragment_shader)

	// x, y, r, g, b per vertex
	vertices := []float32{
		+0.0, +0.8, 1.0, 0.2, 0.2,
		-0.8, -0.6, 0.2, 1.0, 0.2,
		+0.8, -0.6, 0.2, 0.2, 1.0,
	}

	indices := []uint16{
		0, 1, 2,
	}

	array := context.CreateVertexArray()
	context.BindVertexArray(array)

	geometry := context.CreateBuffer()
	context.BindBuffer(webgl.BufferTargetArray, geometry)
	context.BufferDataFloat32(webgl.BufferTargetArray, vertices, webgl.BufferUsageStaticDraw)

	context.VertexAttribPointer(0, 2, webgl.DataTypeFloat, false, 20, 0)
	context.EnableVertexAttribArray(0)
	context.VertexAttribPointer(1, 3, webgl.DataTypeFloat, false, 20, 8)
	context.EnableVertexAttribArray(1)

	elements := context.CreateBuffer()
	context.BindBuffer(webgl.BufferTargetElementArray, elements)
	context.BufferDataUint16(webgl.BufferTargetElementArray, indices, webgl.BufferUsageStaticDraw)

	alpha := context.GetUniformLocation(program, "alpha")

	context.Viewport(0, 0, int(canvas.Width), int(canvas.Height))
	context.Enable(webgl.CapabilityDepthTest)
	context.DepthFunc(webgl.DepthFuncLEqual)

	animations.RequestAnimationFrame(func(timestamp float64) {

		context.ClearColor(0.13, 0.16, 0.23, 1.0)
		context.ClearDepth(1.0)
		context.Clear(webgl.BufferBitColor | webgl.BufferBitDepth)

		context.UseProgram(program)
		context.BindVertexArray(array)
		context.Uniform1f(alpha, 1.0)

		context.DrawElements(webgl.DrawModeTriangles, len(indices), webgl.DataTypeUnsignedShort, 0)

		code := context.GetError()

		if code != webgl.ErrorCodeNoError {
			the_console.Error(code.String())
		} else {
			the_console.Log("WebGL Example: frame drawn")
		}

	})

	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
