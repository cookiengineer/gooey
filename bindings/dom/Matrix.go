//go:build wasm

package dom

import "syscall/js"

type Matrix struct {
	M11        float64  `json:"m11"`
	M12        float64  `json:"m12"`
	M13        float64  `json:"m13"`
	M14        float64  `json:"m14"`
	M21        float64  `json:"m21"`
	M22        float64  `json:"m22"`
	M23        float64  `json:"m23"`
	M24        float64  `json:"m24"`
	M31        float64  `json:"m31"`
	M32        float64  `json:"m32"`
	M33        float64  `json:"m33"`
	M34        float64  `json:"m34"`
	M41        float64  `json:"m41"`
	M42        float64  `json:"m42"`
	M43        float64  `json:"m43"`
	M44        float64  `json:"m44"`
	A          float64  `json:"a"`
	B          float64  `json:"b"`
	C          float64  `json:"c"`
	D          float64  `json:"d"`
	E          float64  `json:"e"`
	F          float64  `json:"f"`
	Is2D       bool     `json:"is2D"`
	IsIdentity bool     `json:"isIdentity"`
	Value      *js.Value `json:"value"`
}

func NewIdentityMatrix() *Matrix {

	var matrix Matrix

	value := js.Global().Get("DOMMatrix").New()
	matrix.Value = &value

	matrix.sync_from_js()

	return &matrix

}

func NewMatrix(values []float64) *Matrix {

	var matrix Matrix

	if len(values) == 6 {

		array := js.Global().Get("Array").New(6)
		array.SetIndex(0, values[0])
		array.SetIndex(1, values[1])
		array.SetIndex(2, values[2])
		array.SetIndex(3, values[3])
		array.SetIndex(4, values[4])
		array.SetIndex(5, values[5])

		value := js.Global().Get("DOMMatrix").New(array)
		matrix.Value = &value

	} else if len(values) == 16 {

		array := js.Global().Get("Array").New(16)

		for v := 0; v < 16; v++ {
			array.SetIndex(v, values[v])
		}

		value := js.Global().Get("DOMMatrix").New(array)
		matrix.Value = &value

	}

	matrix.sync_from_js()

	return &matrix

}

func (matrix *Matrix) clone() *Matrix {

	var clone Matrix

	value := js.Global().Get("DOMMatrix").Call("fromMatrix", matrix.Value)
	clone.Value = &value

	clone.sync_from_js()

	return &clone

}

func (matrix *Matrix) sync_from_js() {

	if matrix.Value == nil {
		return
	}

	val := *matrix.Value

	matrix.M11 = val.Get("m11").Float()
	matrix.M12 = val.Get("m12").Float()
	matrix.M13 = val.Get("m13").Float()
	matrix.M14 = val.Get("m14").Float()
	matrix.M21 = val.Get("m21").Float()
	matrix.M22 = val.Get("m22").Float()
	matrix.M23 = val.Get("m23").Float()
	matrix.M24 = val.Get("m24").Float()
	matrix.M31 = val.Get("m31").Float()
	matrix.M32 = val.Get("m32").Float()
	matrix.M33 = val.Get("m33").Float()
	matrix.M34 = val.Get("m34").Float()
	matrix.M41 = val.Get("m41").Float()
	matrix.M42 = val.Get("m42").Float()
	matrix.M43 = val.Get("m43").Float()
	matrix.M44 = val.Get("m44").Float()

	matrix.A = val.Get("a").Float()
	matrix.B = val.Get("b").Float()
	matrix.C = val.Get("c").Float()
	matrix.D = val.Get("d").Float()
	matrix.E = val.Get("e").Float()
	matrix.F = val.Get("f").Float()

	matrix.Is2D = val.Get("is2D").Bool()
	matrix.IsIdentity = val.Get("isIdentity").Bool()

}

func (matrix *Matrix) Inverse() *Matrix {

	clone := matrix.clone()
	clone.Value.Call("invertSelf")
	clone.sync_from_js()

	return clone

}

func (matrix *Matrix) Multiply(other *Matrix) *Matrix {

	clone := matrix.clone()
	clone.Value.Call("multiplySelf", other.Value)
	clone.sync_from_js()

	return clone

}

func (matrix *Matrix) Scale(scale_x float64, scale_y float64, scale_z float64, origin_x float64, origin_y float64, origin_z float64) *Matrix {

	clone := matrix.clone()
	clone.Value.Call("scaleSelf", scale_x, scale_y, scale_z, origin_x, origin_y, origin_z)
	clone.sync_from_js()

	return clone

}

func (matrix *Matrix) Translate(translate_x float64, translate_y float64, translate_z float64) *Matrix {

	clone := matrix.clone()
	clone.Value.Call("translateSelf", translate_x, translate_y, translate_z)
	clone.sync_from_js()

	return clone

}

func (matrix *Matrix) InvertSelf() *Matrix {

	matrix.Value.Call("invertSelf")
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) MultiplySelf(other *Matrix) *Matrix {

	matrix.Value.Call("multiplySelf", other.Value)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) PreMultiplySelf(other *Matrix) *Matrix {

	matrix.Value.Call("preMultiplySelf", other.Value)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) TranslateSelf(translate_x float64, translate_y float64, translate_z float64) *Matrix {

	matrix.Value.Call("translateSelf", translate_x, translate_y, translate_z)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) ScaleSelf(scale_x float64, scale_y float64, scale_z float64, origin_x float64, origin_y float64, origin_z float64) *Matrix {

	matrix.Value.Call("scaleSelf", scale_x, scale_y, scale_z, origin_x, origin_y, origin_z)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) Scale3dSelf(scale float64, origin_x float64, origin_y float64, origin_z float64) *Matrix {

	matrix.Value.Call("scale3dSelf", scale, origin_x, origin_y, origin_z)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) RotateSelf(rot_x float64, rot_y float64, rot_z float64) *Matrix {

	matrix.Value.Call("rotateSelf", rot_x, rot_y, rot_z)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) RotateAxisAngleSelf(x float64, y float64, z float64, angle float64) *Matrix {

	matrix.Value.Call("rotateAxisAngleSelf", x, y, z, angle)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) RotateFromVectorSelf(x float64, y float64) *Matrix {

	matrix.Value.Call("rotateFromVectorSelf", x, y)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) SkewXSelf(sx float64) *Matrix {

	matrix.Value.Call("skewXSelf", sx)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) SkewYSelf(sy float64) *Matrix {

	matrix.Value.Call("skewYSelf", sy)
	matrix.sync_from_js()

	return matrix

}

func (matrix *Matrix) SetMatrixValue(transform_list string) *Matrix {

	matrix.Value.Call("setMatrixValue", transform_list)
	matrix.sync_from_js()

	return matrix

}

