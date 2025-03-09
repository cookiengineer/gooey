package fetch

import "bytes"
import "errors"
import "io"
import "syscall/js"

type fetch_state struct {
	response *Response
	err      error
}

func Fetch(url string, request *Request) (*Response, error) {

	channel := make(chan *fetch_state)
	done := make(chan struct{}, 1)

	wrapped_options := make(map[string]interface{})

	if request != nil {

		wrapped_options = request.MapToJS()

		if request.Signal != nil {

			controller := js.Global().Get("AbortController").New()
			signal := controller.Get("signal")
			wrapped_options["signal"] = signal

			go func() {
				select {
				case <-request.Signal.Done():
					controller.Call("abort")
				case <-done:
				}
			}()

		}

		if request.Body != nil {

			buffer := make([]byte, 0)

			switch tmp1 := request.Body.(type) {
			case *bytes.Buffer:

				buffer = tmp1.Bytes()

			default:

				tmp2, err := io.ReadAll(request.Body)

				if err == nil {
					buffer = tmp2
				}

			}

			array := js.Global().Get("Uint8Array").New(len(buffer))
			js.CopyBytesToJS(array, buffer)

			wrapped_options["body"] = array

		}

	}

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		response := ToResponse(args[0])

		channel <- &fetch_state{
			response: &response,
			err:      nil,
		}

		return nil

	})

	defer on_success.Release()

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()

		done <- struct{}{}
		channel <- &fetch_state{
			response: nil,
			err:      errors.New(message),
		}

		return nil

	})

	defer on_failure.Release()

	go js.Global().Call("fetch", url, wrapped_options).Call("then", on_success).Call("catch", on_failure)

	state := <-channel

	if state.err != nil {
		return nil, state.err
	}

	on_arraybuffer_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		array := js.Global().Get("Uint8Array").New(args[0])
		state.response.Body = make([]byte, array.Get("byteLength").Int())
		js.CopyBytesToGo(state.response.Body, array)

		channel <- state

		return nil

	})

	defer on_arraybuffer_success.Release()

	on_arraybuffer_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		value := args[0]
		message := value.Get("message").String()
		channel <- &fetch_state{
			response: nil,
			err:      errors.New(message),
		}

		return nil

	})

	defer on_arraybuffer_failure.Release()

	go state.response.Value.Call("arrayBuffer").Call("then", on_arraybuffer_success, on_arraybuffer_failure)

	state = <- channel

	return state.response, state.err

}
