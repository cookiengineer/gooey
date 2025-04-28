//go:build wasm

package cookiestore

import "errors"
import "syscall/js"

func Set(options SetOptions) error {

	cookiestore := js.Global().Get("cookieStore")

	if cookiestore.IsNull() || cookiestore.IsUndefined() {
		return errors.New("Error: CookieStore API not supported.")
	}

	channel := make(chan error)
	wrapped_options := js.ValueOf(options.MapToJS())

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		channel <- nil

		return nil

	})

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		channel <- errors.New(message)

		return nil

	})

	go cookiestore.Call("set", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	err := <-channel

	return err

}
