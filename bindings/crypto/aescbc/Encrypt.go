package aescbc

import "errors"
import "syscall/js"

type encrypt_state struct {
	buffer []byte
	err    error
}

func Encrypt(iv []byte, key *CryptoKey, buffer []byte) ([]byte, error) {

	if key != nil && key.Value != nil {

		channel := make(chan *encrypt_state)

		algorithm := make(map[string]any)
		wrapped_algorithm := js.ValueOf(algorithm)
		wrapped_iv_array := js.Global().Get("Uint8Array").New(len(iv))
		js.CopyBytesToJS(wrapped_iv_array, iv)
		wrapped_algorithm.Set("name", "AES-CBC")
		wrapped_algorithm.Set("iv", wrapped_iv_array)

		wrapped_key := *key.Value
		wrapped_buffer := js.Global().Get("Uint8Array").New(len(buffer))

		js.CopyBytesToJS(wrapped_buffer, buffer)

		on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

			array := js.Global().Get("Uint8Array").New(args[0])
			buffer := make([]byte, array.Get("byteLength").Int())
			js.CopyBytesToGo(buffer, array)

			channel <- &encrypt_state{
				buffer: buffer,
				err:    nil,
			}

			return nil

		})

		defer on_success.Release()

		on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

			value := args[0]
			message := value.Get("message").String()

			channel <- &encrypt_state{
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

		go subtle.Call("encrypt", wrapped_algorithm, wrapped_key, wrapped_buffer).Call("then", on_success).Call("catch", on_failure)

		state := <-channel

		return state.buffer, state.err

	} else {
		return []byte{}, errors.New("Error: Invalid CryptoKey")
	}

}
