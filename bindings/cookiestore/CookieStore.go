//go:build wasm

package cookiestore

import "errors"
import "syscall/js"

var global_cookiestore *CookieStore

func init() {
	global_cookiestore = GetCookieStore()
}

type get_state struct {
	cookie *Cookie
	err    error
}

type getall_state struct {
	cookies []*Cookie
	err     error
}

type CookieStore struct {
	Value *js.Value `json:"value"`
}

// Returns the global CookieStore instance.
func GetCookieStore() *CookieStore {

	if global_cookiestore != nil {

		return global_cookiestore

	} else {

		value := js.Global().Get("cookieStore")
		cookiestore := CookieStore{
			Value: &store,
		}

		return &cookiestore

	}

}

// Deletes a matching Cookie from the CookieStore.
func (cookiestore *CookieStore) Delete(options DeleteOptions) error {

	if cookiestore.Value.IsNull() || cookiestore.Value.IsUndefined() {
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

	go cookiestore.Value.Call("delete", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	err := <-channel

	return err

}

// Returns a matching Cookie from the CookieStore.
func (cookiestore *CookieStore) Get(options GetOptions) (*Cookie, error) {

	if cookiestore.Value.IsNull() || cookiestore.Value.IsUndefined() {
		return errors.New("Error: CookieStore API not supported.")
	}

	channel := make(chan *get_state)
	wrapped_options := js.ValueOf(options.MapToJS())

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]

		if !value.IsNull() && !value.IsUndefined() {

			cookie := ToCookie(value)

			channel <- &get_state{
				cookie: cookie,
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

		channel <- &get_state{
			cookie: nil,
			err:    errors.New(message),
		}

		return nil

	})

	go cookiestore.Value.Call("get", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.cookie, state.err

}

// Returns all matching Cookies from the CookieStore.
func (cookiestore *CookieStore) GetAll(options GetOptions) ([]*Cookie, error) {

	if cookiestore.Value.IsNull() || cookiestore.Value.IsUndefined() {
		return errors.New("Error: CookieStore API not supported.")
	}

	channel := make(chan *getall_state)
	wrapped_options := js.Value{}

	if options != nil {
		wrapped_options = js.ValueOf(options.MapToJS())
	}

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		cookies := make([]*Cookie, 0)
		values := args[0]

		if !values.IsNull() && !values.IsUndefined() && values.Length() > 0 {

			for v := 0; v < values.Length(); v++ {

				value := values.Index(v)
				cookie := ToCookie(value)

				if cookie.Name != "" && cookie.Value != "" {
					cookies = append(cookies, cookie)
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

	go cookiestore.Value.Call("getAll", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	return state.cookies, state.err

}

// Stores a Cookie to the CookieStore.
func (cookiestore *CookieStore) Set(options SetOptions) error {

	if cookiestore.Value.IsNull() || cookiestore.Value.IsUndefined() {
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

	go cookiestore.Value.Call("set", wrapped_options).Call("then", on_success).Call("catch", on_failure)

	err := <-channel

	return err

}
