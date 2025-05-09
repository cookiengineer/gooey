package aesgcm

import "errors"
import "syscall/js"

type importkey_state struct {
	key *CryptoKey
	err error
}

func ImportKey(format string, keydata []byte, extractable bool, usages []string) (*CryptoKey, error) {

	channel := make(chan *importkey_state)

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		key := ToCryptoKey(args[0])

		channel <- &importkey_state{
			key: key,
			err: nil,
		}

		return nil

	})

	defer on_success.Release()

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		channel <- &importkey_state{
			key: nil,
			err: errors.New(message),
		}

		return nil

	})

	defer on_failure.Release()

	subtle := js.Global().Get("crypto").Get("subtle")

	if subtle.IsNull() || subtle.IsUndefined() {
		err := errors.New("Error: Unsecure WebPage Context.")
		return nil, err
	}

	if format == "jwk" {

		wrapped_format := js.ValueOf(format)
		wrapped_keydata := js.Global().Get("JSON").Call("parse", string(keydata))
		wrapped_algorithm := js.ValueOf(map[string]any{
			"name": "AES-GCM",
		})
		wrapped_extractable := js.ValueOf(extractable)
		wrapped_usages := js.Global().Get("Array").New(len(usages))

		for u := 0; u < len(usages); u++ {
			wrapped_usages.SetIndex(u, usages[u])
		}

		go subtle.Call("importKey", wrapped_format, wrapped_keydata, wrapped_algorithm, wrapped_extractable, wrapped_usages).Call("then", on_success).Call("catch", on_failure)

		state := <-channel

		return state.key, state.err

	} else {

		wrapped_format := js.ValueOf(format)
		wrapped_keydata := js.Global().Get("Uint8Array").New(len(keydata))
		wrapped_algorithm := js.ValueOf(map[string]any{
			"name": "AES-GCM",
		})
		wrapped_extractable := js.ValueOf(extractable)
		wrapped_usages := js.Global().Get("Array").New(len(usages))

		js.CopyBytesToJS(wrapped_keydata, keydata)

		for u := 0; u < len(usages); u++ {
			wrapped_usages.SetIndex(u, usages[u])
		}

		go subtle.Call("importKey", wrapped_format, wrapped_keydata, wrapped_algorithm, wrapped_extractable, wrapped_usages).Call("then", on_success).Call("catch", on_failure)

		state := <-channel

		return state.key, state.err

	}

}
