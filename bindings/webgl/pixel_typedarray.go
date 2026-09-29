//go:build wasm

package webgl

import "syscall/js"

// The specification requires the TypedArray a pixel call reads or writes to
// match the pixel type it was given: readPixels(), texImage2D() and
// texSubImage2D() all raise INVALID_OPERATION when it does not, and readPixels
// then leaves the view untouched — a slice of zeroes rather than an error. So
// the mapping from a DataType to its view lives in one place, and the calls
// that need it refuse a type they cannot honour instead of guessing.
//
// The returned name is the JavaScript constructor and the size is one
// element in bytes. A DataType that is not a pixel type answers ("", 0).
func pixel_typedarray(kind DataType) (string, int) {

	switch kind {
	case DataTypeByte:
		return "Int8Array", 1
	case DataTypeUnsignedByte:
		return "Uint8Array", 1
	case DataTypeShort:
		return "Int16Array", 2
	case DataTypeUnsignedShort, DataTypeHalfFloat:
		return "Uint16Array", 2
	case DataTypeInt:
		return "Int32Array", 4
	case DataTypeUnsignedInt:
		return "Uint32Array", 4
	case DataTypeFloat:
		return "Float32Array", 4
	}

	return "", 0

}

// The number of components one pixel of a format carries, which is what turns
// a width and a height into a length. A format the package does not know
// answers 0.
func format_components(format TextureFormat) int {

	switch format {
	case TextureFormatAlpha, TextureFormatDepthComponent, TextureFormatDepthComponent16,
		TextureFormatDepthComponent24, TextureFormatDepthComponent32F, TextureFormatLuminance,
		TextureFormatR8, TextureFormatR32F, TextureFormatRed, TextureFormatRedInteger:
		return 1
	case TextureFormatDepthStencil, TextureFormatDepth24Stencil8, TextureFormatLuminanceAlpha,
		TextureFormatRG, TextureFormatRG8:
		return 2
	case TextureFormatRGB, TextureFormatRGB8, TextureFormatRGB16F, TextureFormatRGB32F,
		TextureFormatRGBInteger, TextureFormatSRGB8:
		return 3
	case TextureFormatRGBA, TextureFormatRGBA8, TextureFormatRGBA16F, TextureFormatRGBA32F,
		TextureFormatRGBAInteger, TextureFormatSRGB8Alpha8:
		return 4
	}

	return 0

}

// Builds the view an upload needs: the caller's bytes, seen through the
// TypedArray the pixel type asks for. It answers js.Undefined() for a type
// this package does not map, or for a byte count that is not a whole number of
// elements.
func pixels_to_typedarray(staging *typed_array_staging, kind DataType, data []byte) js.Value {

	constructor, size := pixel_typedarray(kind)

	if constructor == "" || len(data) == 0 || len(data)%size != 0 {
		return js.Undefined()
	}

	staging.resize(len(data))
	copy(staging.bytes[0:len(data)], data)

	return slice_to_typedarray(staging, constructor, len(data), len(data)/size)

}

// Builds the view a read-back writes into. It allocates rather than borrowing
// the staging array, because the caller keeps the bytes and a read-back is not
// on the frame path anyway.
func new_pixel_typedarray(kind DataType, elements int) js.Value {

	constructor, _ := pixel_typedarray(kind)

	if constructor == "" || elements <= 0 {
		return js.Undefined()
	}

	return js.Global().Get(constructor).New(elements)

}
