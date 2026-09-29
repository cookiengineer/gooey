//go:build wasm

package animations

import "syscall/js"

func RequestAnimationFrame(callback AnimationCallback) uint {

	var result uint = 0
	var wrapped js.Func

	wrapped = js.FuncOf(func(this js.Value, args []js.Value) any {

		defer func() {
			frames.Delete(result)
			wrapped.Release()
		}()

		if len(args) == 1 {
			timestamp := args[0].Float()
			callback(timestamp)
		}

		return nil

	})

	tmp := js.Global().Call("requestAnimationFrame", wrapped)

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = uint(tmp.Int())
		frames.Store(result, wrapped)
	} else {
		wrapped.Release()
	}

	return result

}
