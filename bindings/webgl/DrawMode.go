//go:build wasm

package webgl

type DrawMode uint

const (
	DrawModeLineLoop      DrawMode = 0x0002
	DrawModeLineStrip     DrawMode = 0x0003
	DrawModeLines         DrawMode = 0x0001
	DrawModePoints        DrawMode = 0x0000
	DrawModeTriangleFan   DrawMode = 0x0006
	DrawModeTriangleStrip DrawMode = 0x0005
	DrawModeTriangles     DrawMode = 0x0004
)
