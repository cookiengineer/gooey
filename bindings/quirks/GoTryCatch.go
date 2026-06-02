//go:build wasm

package quirks

import "fmt"
import "syscall/js"

var js_wrapper_function js.Value

// Compiles a reusable JS try/catch trampoline.
func init() {
	js_wrapper_function = js.Global().Get("Function").New(`
		return function(fn) {
			try {
				fn();
				return null;
			} catch (e) {
				return e;
			}
		}
	`).Invoke()
}

// GoTryCatch executes function and converts any JS exception into a Go error.
//
// Example:
//
//	err := helpers.GoTryCatch(func() {
//		location.JSValue.Call("assign", js.ValueOf(url))
//	})

func GoTryCatch(function func()) error {

	wrapped := js.FuncOf(func(this js.Value, args []js.Value) any {
		function()
		return nil
	})

	defer wrapped.Release()

	err := js_wrapper_function.Invoke(wrapped)

	if err.IsNull() || err.IsUndefined() {

		return nil

	} else {

		message := err.Get("message")

		if message.Type() == js.TypeString {
			return fmt.Errorf(message.String())
		} else {
			return fmt.Errorf("%v", err)
		}

	}

}
