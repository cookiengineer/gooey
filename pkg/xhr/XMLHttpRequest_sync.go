//go:build wasm

package xhr

import "syscall/js"
import "time"

func (xhr *XMLHttpRequest) OpenSync(method Method, url string) {

	if method.String() != "" && url != "" {
		xhr.Value.Call("open", method.String(), url, false)
	}

}

func (xhr *XMLHttpRequest) SendSync() {

	on_load := js.FuncOf(func(this js.Value, args []js.Value) any {

		xhr.ReadyState = xhr.Value.Get("readyState").Int()
		xhr.ResponseURL = xhr.Value.Get("responseURL").String()
		xhr.Status = xhr.Value.Get("status").Int()
		xhr.StatusText = xhr.Value.Get("statusText").String()

		timeout := xhr.Value.Get("timeout").Int()

		if timeout >= 0 {
			xhr.Timeout = time.Duration(timeout) * time.Millisecond
		}

		xhr.WithCredentials = xhr.Value.Get("withCredentials").Bool()

		response_text := xhr.Value.Get("responseText")
		buffer := []byte(response_text.String())
		xhr.Response = buffer

		if xhr.OnLoad != nil {
			xhr.OnLoad(xhr.Status, xhr.Response)
		}

		return nil

	})

	on_error := js.FuncOf(func(this js.Value, args []js.Value) any {

		xhr.ReadyState = xhr.Value.Get("readyState").Int()
		xhr.ResponseURL = xhr.Value.Get("responseURL").String()
		xhr.Status = xhr.Value.Get("status").Int()
		xhr.StatusText = xhr.Value.Get("statusText").String()

		timeout := xhr.Value.Get("timeout").Int()

		if timeout >= 0 {
			xhr.Timeout = time.Duration(timeout) * time.Millisecond
		}

		xhr.WithCredentials = xhr.Value.Get("withCredentials").Bool()

		if xhr.OnError != nil {
			xhr.OnError()
		}

		return nil

	})

	xhr.Value.Set("onload", on_load)
	xhr.Value.Set("onerror", on_error)

	xhr.Value.Call("send", js.ValueOf(nil))

}
