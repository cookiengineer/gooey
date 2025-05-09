package aesgcm

import "errors"
import "syscall/js"

type generatekey_state struct {
	key *CryptoKey
	err error
}

func GenerateKey(length int, extractable bool, usages []string) (*CryptoKey, error) {

	channel := make(chan *generatekey_state)

	wrapped_algorithm := js.ValueOf(map[string]any{
		"name":   "AES-GCM",
		"length": length,
	})
	wrapped_extractable := js.ValueOf(extractable)
	wrapped_usages := js.Global().Get("Array").New(len(usages))

	for u := 0; u < len(usages); u++ {
		wrapped_usages.SetIndex(u, usages[u])
	}

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		key := ToCryptoKey(args[0])

		channel <- &generatekey_state{
			key: key,
			err: nil,
		}

		return nil

	})

	defer on_success.Release()

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		channel <- &generatekey_state{
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

	go subtle.Call("generateKey", wrapped_algorithm, wrapped_extractable, wrapped_usages).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.key, state.err

}
