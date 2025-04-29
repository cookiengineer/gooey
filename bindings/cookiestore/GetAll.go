//go:build wasm

package cookiestore

import "errors"
import "syscall/js"

type getall_state struct {
	cookies []*Cookie
	err     error
}

func GetAll(options *GetOptions) ([]*Cookie, error) {

	cookiestore := js.Global().Get("cookieStore")

	if cookiestore.IsNull() || cookiestore.IsUndefined() {
		return nil, errors.New("Error: CookieStore API not supported.")
	}

	channel := make(chan *getall_state)

	wrapped_options := js.Value{}

	if options != nil {
		wrapped_options = js.ValueOf(options.MapToJS())
	}

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		cookies := make([]*Cookie, 0)
		values  := args[0]

		if !values.IsNull() && !values.IsUndefined() && values.Length() > 0 {

			for v := 0; v < values.Length(); v++ {

				value  := values.Index(v)
				cookie := ToCookie(value)

				if cookie.Name != "" && cookie.Value != "" {
					cookies = append(cookies, &cookie)
				}

			}

		}

		channel <- &getall_state{
			cookies: cookies,
			err:     nil,
		}

		return nil

	})

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		cookies := make([]*Cookie, 0)
		value := args[0]
		message := value.Get("message").String()

		channel <- &getall_state{
			cookies: cookies,
			err:     errors.New(message),
		}

		return nil

	})

	go cookiestore.Call("getAll", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.cookies, state.err

}

