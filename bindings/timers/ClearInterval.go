//go:build wasm

package timers

import "syscall/js"

func ClearInterval(handler_id uint64) {

	wrapped_identifier := js.ValueOf(handler_id)

	js.Global().Call("clearInterval", wrapped_identifier)

}
