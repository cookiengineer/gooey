package console

import "syscall/js"

func GroupEnd(title string) {
	js.Global().Get("console").Call("groupEnd", js.ValueOf(title))
}
