package aescbc

import "errors"
import "syscall/js"

type export_state struct {
	buffer []byte
	err    error
}

func ExportKey(format string, key CryptoKey) ([]byte, error) {

	channel := make(chan *export_state)

	wrapped_format := js.ValueOf(format)
	wrapped_key := *key.Value

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		if format == "jwk" {

			// Promise fulfills with JSON Object
			value := js.Global().Get("JSON").Call("stringify", args[0])
			buffer := []byte(value.String())

			channel <- &export_state{
				buffer: buffer,
				err:    nil,
			}

		} else {

			// Promise fulfills with ArrayBuffer
			array := js.Global().Get("Uint8Array").New(args[0])
			buffer := make([]byte, array.Get("byteLength").Int())
			js.CopyBytesToGo(buffer, array)

			channel <- &export_state{
				buffer: buffer,
				err:    nil,
			}

		}

		return nil

	})

	defer on_success.Release()

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		channel <- &export_state{
			buffer: []byte{},
			err:    errors.New(message),
		}

		return nil

	})

	defer on_failure.Release()

	subtle := js.Global().Get("crypto").Get("subtle")

	if subtle.IsNull() || subtle.IsUndefined() {
		err := errors.New("Error: Unsecure WebPage Context.")
		return []byte{}, err
	}

	go subtle.Call("exportKey", wrapped_format, wrapped_key).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.buffer, state.err

}
