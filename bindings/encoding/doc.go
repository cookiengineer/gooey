//go:build wasm

// Package encoding provides bindings for the WHATWG Encoding Living Standard
//
// Specification: https://encoding.spec.whatwg.org/
//
// Example usage:
//
//	import "github.com/cookiengineer/gooey/bindings/console"
//	import "github.com/cookiengineer/gooey/bindings/encoding"
//
//	console := console.GetConsole()
//	encoder := encoding.NewTextEncoder(encoding.EncodingUTF8)
//	decoder := encoding.NewTextDecoder(encoding.EncodingUTF8, encoding.TextDecoderOptions{
//		Fatal:     false,
//		IgnoreBOM: false,
//	})
//
//	encoded_bytes := encoder.Encode("€")
//	console.Log("Encoded Bytes: ")
//	console.Log(encoded_bytes)
//
//	decoded_string := decoder.Decode(encoded_bytes)
//	console.Log("Decoded String: ")
//	console.Log(decoded_string)
//
package encoding
