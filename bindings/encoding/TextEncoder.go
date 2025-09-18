//go:build wasm

package encoding

import "syscall/js"

type TextEncoder struct {
	Encoding Encoding `json:"encoding"`
	Value    *js.Value
}

func NewTextEncoder(encoding Encoding) *TextEncoder {

	if encoding == EncodingUTF_8 {

		var encoder TextEncoder

		encoder.Encoding = encoding

		constructor := js.Global().Get("TextEncoder")

		if !constructor.IsNull() && !constructor.IsUndefined() {
			value := constructor.New()
			encoder.Value = &value
		}

		return &encoder

	}

	return nil

}

func ToTextEncoder(value js.Value) *TextEncoder {

	var encoder TextEncoder

	encoder.Encoding = EncodingUTF_8
	encoder.Value = &value

	return &encoder

}

func (encoder *TextEncoder) Encode(value string) []byte {

	var result []byte

	if encoder.Value != nil {

		wrapped_string := js.ValueOf(value)

		value := encoder.Value.Get("encode").Call(wrapped_string)

		if !value.IsNull() && !value.IsUndefined() {
			result = make([]byte, value.Get("byteLength").Int())
			js.CopyBytesToGo(result, value)
		}

	}

	return result

}
