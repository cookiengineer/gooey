//go:build wasm

package timers

import "syscall/js"

func ClearTimeout(identifier uint) {

	wrapped_identifier := js.ValueOf(identifier)

	js.Global().Call("clearTimeout", wrapped_identifier)

}
