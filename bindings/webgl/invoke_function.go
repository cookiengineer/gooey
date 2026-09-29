//go:build wasm

package webgl

import "syscall/js"

// Invokes a method that bind_function() bound earlier. An unimplemented method
// returns js.Undefined() instead of throwing, so that a missing WebGL2 feature
// degrades into a null result rather than into a panic.
func invoke_function(method js.Value, args ...any) js.Value {

	if method.IsUndefined() {
		return js.Undefined()
	}

	return method.Invoke(args...)

}
