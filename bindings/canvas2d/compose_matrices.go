//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"

func compose_matrices(a *dom.Matrix, b *dom.Matrix) *dom.Matrix {

	if b.IsIdentity {
		return a
	}

	if a.IsIdentity {
		return b
	}

	result := &dom.Matrix{
		M11: a.M11*b.M11 + a.M21*b.M12 + a.M31*b.M13 + a.M41*b.M14,
		M12: a.M12*b.M11 + a.M22*b.M12 + a.M32*b.M13 + a.M42*b.M14,
		M13: a.M13*b.M11 + a.M23*b.M12 + a.M33*b.M13 + a.M43*b.M14,
		M14: a.M14*b.M11 + a.M24*b.M12 + a.M34*b.M13 + a.M44*b.M14,
		M21: a.M11*b.M21 + a.M21*b.M22 + a.M31*b.M23 + a.M41*b.M24,
		M22: a.M12*b.M21 + a.M22*b.M22 + a.M32*b.M23 + a.M42*b.M24,
		M23: a.M13*b.M21 + a.M23*b.M22 + a.M33*b.M23 + a.M43*b.M24,
		M24: a.M14*b.M21 + a.M24*b.M22 + a.M34*b.M23 + a.M44*b.M24,
		M31: a.M11*b.M31 + a.M21*b.M32 + a.M31*b.M33 + a.M41*b.M34,
		M32: a.M12*b.M31 + a.M22*b.M32 + a.M32*b.M33 + a.M42*b.M34,
		M33: a.M13*b.M31 + a.M23*b.M32 + a.M33*b.M33 + a.M43*b.M34,
		M34: a.M14*b.M31 + a.M24*b.M32 + a.M34*b.M33 + a.M44*b.M34,
		M41: a.M11*b.M41 + a.M21*b.M42 + a.M31*b.M43 + a.M41*b.M44,
		M42: a.M12*b.M41 + a.M22*b.M42 + a.M32*b.M43 + a.M42*b.M44,
		M43: a.M13*b.M41 + a.M23*b.M42 + a.M33*b.M43 + a.M43*b.M44,
		M44: a.M14*b.M41 + a.M24*b.M42 + a.M34*b.M43 + a.M44*b.M44,
	}

	result.A = result.M11
	result.B = result.M12
	result.C = result.M21
	result.D = result.M22
	result.E = result.M41
	result.F = result.M42

	result.Is2D = result.M13 == 0 && result.M14 == 0 &&
		result.M23 == 0 && result.M24 == 0 &&
		result.M31 == 0 && result.M32 == 0 &&
		result.M33 == 1 && result.M34 == 0 &&
		result.M43 == 0 && result.M44 == 1

	result.IsIdentity = result.M11 == 1 && result.M12 == 0 && result.M13 == 0 && result.M14 == 0 &&
		result.M21 == 0 && result.M22 == 1 && result.M23 == 0 && result.M24 == 0 &&
		result.M31 == 0 && result.M32 == 0 && result.M33 == 1 && result.M34 == 0 &&
		result.M41 == 0 && result.M42 == 0 && result.M43 == 0 && result.M44 == 1

	return result

}

