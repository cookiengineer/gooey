//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleVertexArray() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()

	vertices := []float32{
		-0.5, -0.5, 0.0,
		 0.5, -0.5, 0.0,
		 0.5,  0.5, 0.0,
		-0.5,  0.5, 0.0,
	}

	indices := []uint16{
		0, 1, 2,
		0, 2, 3,
	}

	// One instance record per copy, read once per instance instead of once per vertex
	offsets := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
	}

	array := context.CreateVertexArray()
	context.BindVertexArray(array)

	geometry := context.CreateBuffer()
	context.BindBuffer(BufferTargetArray, geometry)
	context.BufferDataFloat32(BufferTargetArray, vertices, BufferUsageStaticDraw)
	context.VertexAttribPointer(0, 3, DataTypeFloat, false, 0, 0)
	context.EnableVertexAttribArray(0)

	instances := context.CreateBuffer()
	context.BindBuffer(BufferTargetArray, instances)
	context.BufferDataFloat32(BufferTargetArray, offsets, BufferUsageDynamicDraw)
	context.VertexAttribPointer(1, 2, DataTypeFloat, false, 0, 0)
	context.EnableVertexAttribArray(1)
	context.VertexAttribDivisor(1, 1)

	elements := context.CreateBuffer()
	context.BindBuffer(BufferTargetElementArray, elements)
	context.BufferDataUint16(BufferTargetElementArray, indices, BufferUsageStaticDraw)

	context.DrawElementsInstanced(DrawModeTriangles, len(indices), DataTypeUnsignedShort, 0, 3)

	context.BindVertexArray(nil)

}
