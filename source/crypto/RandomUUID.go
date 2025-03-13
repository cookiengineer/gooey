//go:build wasm

package crypto

import "syscall/js"

func RandomUUID() string {

	var result string

	tmp := js.Global().Get("crypto").Call("randomUUID")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = tmp.String()
	}

	return result

}
