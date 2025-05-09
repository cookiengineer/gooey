package aesctr

import "errors"
import "syscall/js"

type decrypt_state struct {
	buffer []byte
	err    error
}

func Decrypt(counter []byte, length uint, key *CryptoKey, buffer []byte) ([]byte, error) {

	if key != nil && key.Value != nil {

		if len(counter) == 16 && length >= 64 && length <= 128 {

			channel := make(chan *decrypt_state)

			algorithm := make(map[string]any)
			wrapped_algorithm := js.ValueOf(algorithm)
			wrapped_counter_array := js.Global().Get("Uint8Array").New(len(counter))
			js.CopyBytesToJS(wrapped_counter_array, counter)
			wrapped_algorithm.Set("name", "AES-CTR")
			wrapped_algorithm.Set("counter", wrapped_counter_array)
			wrapped_algorithm.Set("length", js.ValueOf(length))

			wrapped_key := *key.Value
			wrapped_buffer := js.Global().Get("Uint8Array").New(len(buffer))

			js.CopyBytesToJS(wrapped_buffer, buffer)

			on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

				array := js.Global().Get("Uint8Array").New(args[0])
				buffer := make([]byte, array.Get("byteLength").Int())
				js.CopyBytesToGo(buffer, array)

				channel <- &decrypt_state{
					buffer: buffer,
					err:    nil,
				}

				return nil

			})

			defer on_success.Release()

			on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

				value := args[0]
				message := value.Get("message").String()

				channel <- &decrypt_state{
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

			go subtle.Call("decrypt", wrapped_algorithm, wrapped_key, wrapped_buffer).Call("then", on_success).Call("catch", on_failure)

			state := <-channel

			return state.buffer, state.err

		} else if len(counter) != 16 {
			return []byte{}, errors.New("AES-CTR expects counter to be 16 bytes long")
		} else {
			return []byte{}, errors.New("AES-CTR expects nonce/counter length to be between 64 and 128 bits")
		}

	} else {
		return []byte{}, errors.New("Error: Invalid CryptoKey")
	}

}
