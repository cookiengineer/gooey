//go:build wasm

package timers

import "syscall/js"

func ClearInterval(identifier uint) {

	wrapped_identifier := js.ValueOf(identifier)

	js.Global().Call("clearInterval", wrapped_identifier)

}
