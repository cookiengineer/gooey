//go:build wasm

package webgl

import "syscall/js"

// The specification defines exactly one target for a Renderbuffer, so it is a
// constant here rather than a parameter that only ever takes one value.
const renderbuffer_target uint = 0x8D41

type Renderbuffer struct {
	Value *js.Value `json:"value"`
}

func ToRenderbuffer(value js.Value) *Renderbuffer {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var renderbuffer Renderbuffer

	renderbuffer.Value = &value

	return &renderbuffer

}
