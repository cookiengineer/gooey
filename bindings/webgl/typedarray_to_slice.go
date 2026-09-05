//go:build wasm

package webgl

import "encoding/binary"
import "math"
import "syscall/js"

// The read-back direction of slice_to_typedarray(). The TypedArray is viewed as
// a Uint8Array over the same ArrayBuffer, because js.CopyBytesToGo() only
// accepts bytes.
func typedarray_to_slice(value js.Value) []byte {

	bytes := make([]byte, 0)

	if !value.IsNull() && !value.IsUndefined() {

		buffer := value.Get("buffer")

		if !buffer.IsNull() && !buffer.IsUndefined() {

			array := js.Global().Get("Uint8Array").New(buffer, value.Get("byteOffset"), value.Get("byteLength"))
			bytes = make([]byte, array.Length())

			js.CopyBytesToGo(bytes, array)

		}

	}

	return bytes

}

func typedarray_to_uint8(value js.Value) []byte {
	return typedarray_to_slice(value)
}

func typedarray_to_float32(value js.Value) []float32 {

	bytes := typedarray_to_slice(value)
	data := make([]float32, len(bytes)/4)

	for d := 0; d < len(data); d++ {
		data[d] = math.Float32frombits(binary.LittleEndian.Uint32(bytes[d*4:]))
	}

	return data

}
