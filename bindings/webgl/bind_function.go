//go:build wasm

package webgl

import "syscall/js"

// Binds a method of the given value, so that it can be invoked later without
// another property lookup. A method that the Web Browser Engine does not
// implement resolves to js.Undefined(), which is what invoke_function() checks.
func bind_function(value js.Value, name string) js.Value {

	method := value.Get(name)

	if method.IsNull() || method.IsUndefined() {
		return js.Undefined()
	}

	return method.Call("bind", value)

}
