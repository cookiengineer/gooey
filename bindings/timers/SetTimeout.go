//go:build wasm

package timers

import "syscall/js"

// Returns the handler identifier used with ClearTimeout(handler_id)
func SetTimeout(callback func(), delay Duration) uint64 {

	var result uint64 = 0

	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		callback()
		return nil
	})
	wrapped_delay := js.ValueOf(delay)

	tmp := js.Global().Call("setTimeout", wrapped_callback, wrapped_delay)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint64(tmp.Int())
	}

	return result

}
