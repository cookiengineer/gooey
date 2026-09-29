//go:build wasm

package webgl

import "encoding/binary"
import "math"
import "syscall/js"

// syscall/js can only copy bytes across the boundary, so a TypedArray is built
// by copying into a Uint8Array and then viewing its ArrayBuffer. Both the Go
// scratch slice and the Uint8Array are kept on the Context and only ever grow,
// so that uploading a mesh or a matrix does not allocate on either heap once
// the render loop has settled.
type typed_array_staging struct {
	bytes []byte
	array js.Value
	size  int
}

func (staging *typed_array_staging) resize(size int) {

	if staging.size < size {

		staging.array = js.Global().Get("Uint8Array").New(size)
		staging.size = size

	}

	if cap(staging.bytes) < size {
		staging.bytes = make([]byte, size)
	}

}

func slice_to_typedarray(staging *typed_array_staging, constructor string, size int, length int) js.Value {

	js.CopyBytesToJS(staging.array, staging.bytes[0:size])

	return js.Global().Get(constructor).New(staging.array.Get("buffer"), 0, length)

}

func uint8_to_typedarray(staging *typed_array_staging, data []byte) js.Value {

	staging.resize(len(data))
	copy(staging.bytes[0:len(data)], data)

	return slice_to_typedarray(staging, "Uint8Array", len(data), len(data))

}

func uint16_to_typedarray(staging *typed_array_staging, data []uint16) js.Value {

	size := len(data) * 2

	staging.resize(size)

	for d := 0; d < len(data); d++ {
		binary.LittleEndian.PutUint16(staging.bytes[d*2:], data[d])
	}

	return slice_to_typedarray(staging, "Uint16Array", size, len(data))

}

func uint32_to_typedarray(staging *typed_array_staging, data []uint32) js.Value {

	size := len(data) * 4

	staging.resize(size)

	for d := 0; d < len(data); d++ {
		binary.LittleEndian.PutUint32(staging.bytes[d*4:], data[d])
	}

	return slice_to_typedarray(staging, "Uint32Array", size, len(data))

}

func int32_to_typedarray(staging *typed_array_staging, data []int32) js.Value {

	size := len(data) * 4

	staging.resize(size)

	for d := 0; d < len(data); d++ {
		binary.LittleEndian.PutUint32(staging.bytes[d*4:], uint32(data[d]))
	}

	return slice_to_typedarray(staging, "Int32Array", size, len(data))

}

func float32_to_typedarray(staging *typed_array_staging, data []float32) js.Value {

	size := len(data) * 4

	staging.resize(size)

	for d := 0; d < len(data); d++ {
		binary.LittleEndian.PutUint32(staging.bytes[d*4:], math.Float32bits(data[d]))
	}

	return slice_to_typedarray(staging, "Float32Array", size, len(data))

}
