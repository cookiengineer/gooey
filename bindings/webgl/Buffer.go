//go:build wasm

package webgl

import "syscall/js"

type Buffer struct {
	Value *js.Value `json:"value"`
}

func ToBuffer(value js.Value) *Buffer {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var buffer Buffer

	buffer.Value = &value

	return &buffer

}
