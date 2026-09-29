//go:build wasm

package animations

import "syscall/js"

func CancelAnimationFrame(identifier uint) {

	if value, ok := frames.LoadAndDelete(identifier); ok == true {

		js.Global().Call("cancelAnimationFrame", js.ValueOf(identifier))

		if callback, ok2 := value.(js.Func); ok2 == true {
			callback.Release()
		}

	}

}
