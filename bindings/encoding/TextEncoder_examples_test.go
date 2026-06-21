//go:build wasm

package encoding

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleTextEncoder_Encode() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.NewConsole()
	encoder := NewTextEncoder(EncodingUTF8)

	data := encoder.Encode("Hello, 世界")

	console.Log(data)

	// Output:
	// [72 101 108 108 111 44 32 228 184 150 231 149 140]

}
