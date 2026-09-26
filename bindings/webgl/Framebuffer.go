//go:build wasm

package webgl

import "syscall/js"

type Framebuffer struct {
	Value *js.Value `json:"value"`
}

func ToFramebuffer(value js.Value) *Framebuffer {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var framebuffer Framebuffer

	framebuffer.Value = &value

	return &framebuffer

}
