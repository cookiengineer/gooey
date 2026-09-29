//go:build wasm

package webgl

import "syscall/js"

type Texture struct {
	Value *js.Value `json:"value"`
}

func ToTexture(value js.Value) *Texture {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var texture Texture

	texture.Value = &value

	return &texture

}
