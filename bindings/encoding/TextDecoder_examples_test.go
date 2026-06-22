//go:build wasm

package encoding

import "github.com/cookiengineer/gooey/bindings/console"

func ExampleNewTextDecoder() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	decoder := NewTextDecoder(
		EncodingKOI8R,
		TextDecoderOptions{},
	)

	data := []byte{
		0xF0, 0xD2, 0xC9, 0xD7, 0xC5, 0xD4,
		0x2C, 0x20,
		0xCD, 0xC9, 0xD2,
		0x21,
	}

	text := decoder.Decode(data)

	console.Log(text)

	// Output:
	// Привет, мир!

}

func ExampleTextDecoder_Decode() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.GetConsole()
	decoder := NewTextDecoder(
		EncodingUTF8,
		TextDecoderOptions{},
	)

	data := []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F,
		0x2C, 0x20,
		0xE4, 0xB8, 0x96,
		0xE7, 0x95, 0x8C,
	}

	text := decoder.Decode(data)

	console.Log(text)

	// Output:
	// Hello, 世界

}

