//go:build wasm

package encoding

import "syscall/js"

type TextDecoderOptions struct {
	Fatal     bool `json:"fatal"`
	IgnoreBOM bool `json:"ignoreBOM"`
}

func ToTextDecoderOptions(value js.Value) TextDecoderOptions {

	var options TextDecoderOptions

	options.Fatal = value.Get("fatal").Bool()
	options.IgnoreBOM = value.Get("ignoreBOM").Bool()

	return options

}

func (options *TextDecoderOptions) MapToJS() map[string]any {

	result := make(map[string]any)

	result["fatal"] = options.Fatal
	result["ignoreBOM"] = options.IgnoreBOM

	return result

}
