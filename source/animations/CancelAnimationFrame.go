//go:build wasm

package animations

import "syscall/js"

func CancelAnimationFrame(identifier uint) {

	wrapped_identifier := js.ValueOf(identifier)

	js.Global().Call("cancelAnimationFrame", wrapped_identifier)

}
