//go:build wasm

package cookiestore

import "errors"
import "syscall/js"

type get_state struct {
	cookie *Cookie
	err    error
}

func Get(options GetOptions) (*Cookie, error) {

	cookiestore := js.Global().Get("cookieStore")

	if cookiestore.IsNull() || cookiestore.IsUndefined() {
		return nil, errors.New("Error: CookieStore API not supported.")
	}

	channel := make(chan *get_state)
	wrapped_options := js.ValueOf(options.MapToJS())

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]

		if !value.IsNull() && !value.IsUndefined() {

			cookie := ToCookie(value)

			channel <- &get_state{
				cookie: &cookie,
				err:    nil,
			}

		} else {

			channel <- &get_state{
				cookie: nil,
				err:    errors.New("Error: Unsecure WebPage Context."),
			}

		}

		return nil

	})

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		channel <= &get_state{
			cookie: nil,
			err:    errors.New(message),
		}

		return nil

	})

	go cookiestore.Call("get", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.cookie, state.err

}
