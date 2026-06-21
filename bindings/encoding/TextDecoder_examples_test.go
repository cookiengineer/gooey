package encoding

import "github.com/cookiengineer/gooey/bindings/console"
import "fmt"

func ExampleNewTextDecoder() {

	// import "github.com/cookiengineer/gooey/bindings/console"

	console := console.NewConsole()
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

	console := console.NewConsole()
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

func ExampleTextDecoderOptions() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "fmt"

	console := console.NewConsole()
	decoder := NewTextDecoder(
		EncodingUTF8,
		TextDecoderOptions{
			Fatal:     false,
			IgnoreBOM: true,
		},
	)

	data := []byte{
		0xEF, 0xBB, 0xBF, // UTF-8 BOM
		'H', 'e', 'l', 'l', 'o',
	}

	text := decoder.Decode(data)

	console.Log(fmt.Sprintf("%q\n", text))

	// Output:
	// "\ufeffHello"

}
