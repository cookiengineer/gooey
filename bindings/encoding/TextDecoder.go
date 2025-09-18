//go:build wasm

package encoding

import "syscall/js"

type TextDecoder struct {
	Encoding Encoding           `json:"encoding"`
	Options  TextDecoderOptions `json:"options"`
	Value    *js.Value
}

func NewTextDecoder(encoding Encoding, options TextDecoderOptions) *TextDecoder {

	var decoder TextDecoder

	decoder.Encoding = encoding
	decoder.Options  = options
	decoder.Value    = nil

	constructor := js.Global().Get("TextDecoder")

	if !constructor.IsNull() && !constructor.IsUndefined() {

		wrapped_encoding = js.ValueOf(encoding.String())
		wrapped_options  = js.ValueOf(options.MapToJS())

		value := constructor.New(wrapped_encoding, wrapped_options)
		decoder.Value = &value

	}

	return &decoder

}

func ToTextDecoder(value js.Value) *TextDecoder {

	var decoder TextDecoder

	decoder.Encoding = Encoding(value.Get("encoding").String())
	decoder.Options  = ToTextDecoderOptions(value)
	decoder.Value = &value

	return &decoder

}

func (decoder *TextDecoder) Decode(buffer []byte) string {

	var result string

	if decoder.Value != nil {

		wrapped_buffer := js.Global().Get("Uint8Array").New(len(buffer))
		js.CopyBytesToJS(wrapped_buffer, buffer)

		value := decoder.Value.Get("decode").Call(wrapped_buffer)

		if !value.IsNull() && !value.IsUndefined() {
			result = value.String()
		}

	}

	return result

}
