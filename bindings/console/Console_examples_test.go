//go:build wasm

package console

func Example() {

	console1 := GetConsole()
	console1.Log(1337)
	console1.Debug("foo bar")

}
