//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleBuffer() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	vertices := []float32{
		-0.5, -0.5, 0.0,
		 0.5, -0.5, 0.0,
		 0.0,  0.5, 0.0,
	}

	buffer := context.CreateBuffer()

	context.BindBuffer(BufferTargetArray, buffer)
	context.BufferDataFloat32(BufferTargetArray, vertices, BufferUsageStaticDraw)

	context.VertexAttribPointer(0, 3, DataTypeFloat, false, 0, 0)
	context.EnableVertexAttribArray(0)

	context.DrawArrays(DrawModeTriangles, 0, 3)

	context.BindBuffer(BufferTargetArray, nil)
	context.DeleteBuffer(buffer)

}
