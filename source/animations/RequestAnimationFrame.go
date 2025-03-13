//go:build wasm

package animations

import "syscall/js"

type AnimationCallback func(timestamp float64)

func RequestAnimationFrame(callback AnimationCallback) uint {

	var result uint = 0

	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) == 1 {
			timestamp := args[0].Float()
			callback(timestamp)
		}

	})

	tmp := js.Global().Call("requestAnimationFrame", wrapped_callback)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint(tmp.Int())
	}

	return result

}
