//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"

func transform_point(matrix *dom.Matrix, x float64, y float64) (float64, float64) {

	nx := matrix.A*x + matrix.C*y + matrix.E
	ny := matrix.B*x + matrix.D*y + matrix.F

	return nx, ny

}

