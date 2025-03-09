package console

import "syscall/js"

func Clear() {
	js.Global().Get("console").Call("clear")
}
