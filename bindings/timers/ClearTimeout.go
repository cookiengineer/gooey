//go:build wasm

package timers

import "syscall/js"

func ClearTimeout(handler_id uint) {

	wrapped_identifier := js.ValueOf(handler_id)

	js.Global().Call("clearTimeout", wrapped_identifier)

}
