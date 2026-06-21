//go:build wasm

package encoding

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleTextEncoder_Encode() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	encoder := NewTextEncoder(EncodingUTF8)

	encoded_bytes := encoder.Encode("€")
	console.Log("Encoded Bytes: ")
	console.Log(encoded_bytes)

}
