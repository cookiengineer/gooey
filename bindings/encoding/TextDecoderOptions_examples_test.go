//go:build wasm

package encoding

import "github.com/cookiengineer/gooey/bindings/console"
import "fmt"

func ExampleTextDecoderOptions() {

	// import "github.com/cookiengineer/gooey/bindings/console"
	// import "fmt"

	console := console.GetConsole()
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

