//go:build wasm

package timers

import "syscall/js"

func SetInterval(callback func(), milliseconds uint) uint {

	var result uint = 0

	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		callback()
		return nil
	})
	wrapped_delay := js.ValueOf(milliseconds)

	tmp := js.Global().Call("setInterval", wrapped_callback, wrapped_delay)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint(tmp.Int())
	}

	return result

}
