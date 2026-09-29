//go:build wasm

package webgl

import "syscall/js"

type VertexArray struct {
	Value *js.Value `json:"value"`
}

func ToVertexArray(value js.Value) *VertexArray {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var array VertexArray

	array.Value = &value

	return &array

}
