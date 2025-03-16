//go:build wasm

package console

import "syscall/js"

func Group(title string) {
	js.Global().Get("console").Call("group", js.ValueOf(title))
}
