//go:build wasm

package webgl

import "syscall/js"

type UniformLocation struct {
	Name  string    `json:"name"`
	Value *js.Value `json:"value"`
}

func ToUniformLocation(value js.Value) *UniformLocation {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var location UniformLocation

	location.Name = ""
	location.Value = &value

	return &location

}
