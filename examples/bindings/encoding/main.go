package main

import "github.com/cookiengineer/gooey/bindings/encoding"
import "fmt"
import "time"

func main() {

	encoder := encoding.NewTextEncoder(encoding.EncodingUTF8)
	decoder := encoding.NewTextDecoder(encoding.EncodingUTF8, encoding.TextDecoderOptions{
		Fatal:     false,
		IgnoreBOM: false,
	})

	encoded_bytes := encoder.Encode("€")
	fmt.Println("Encoded Bytes: ", encoded_bytes)

	decoded_string := decoder.Decode(encoded_bytes)
	fmt.Println("Decoded String: ", decoded_string)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
