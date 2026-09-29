//go:build wasm

package webgl

import "syscall/js"

// Reads a boolean property, answering the fallback when the Web Browser Engine
// does not report it. getContextAttributes() is the reason this exists: not
// every engine returns every attribute, and js.Value.Bool() panics on
// undefined rather than answering false.
func bool_or(value js.Value, name string, fallback bool) bool {

	property := value.Get(name)

	if property.IsNull() || property.IsUndefined() {
		return fallback
	}

	return property.Bool()

}
