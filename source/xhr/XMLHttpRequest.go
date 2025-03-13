//go:build wasm

package xhr

import "syscall/js"
import "time"

type XMLHttpRequest struct {
	OnLoad      func(int, []byte) `json:"onload"`
	OnError     func()            `json:"onerror"`
	OnTimeout   func()            `json:"ontimeout"`
	ReadyState  int               `json:"readyState"`
	Response    []byte            `json:"response"`
	ResponseURL string            `json:"responseURL"`
	Status      int               `json:"status"`
	StatusText  string            `json:"statusText"`
	Timeout     time.Duration     `json:"timeout"`
	// Upload not supportable without reimplementing Event API
	WithCredentials bool      `json:"withCredentials"`
	Value           *js.Value `json:"value"`
}

func NewXMLHttpRequest() XMLHttpRequest {

	var xhr XMLHttpRequest

	value := js.Global().Get("XMLHttpRequest").New()

	xhr.OnLoad = func(int, []byte) {}
	xhr.OnError = func() {}
	xhr.OnTimeout = func() {}
	xhr.ReadyState = 0
	xhr.Status = 0
	xhr.WithCredentials = false
	xhr.Value = &value

	return xhr

}

func (xhr *XMLHttpRequest) Abort() {
	xhr.Value.Call("abort")
}

func (xhr *XMLHttpRequest) GetAllResponseHeaders() string {

	var result string = ""

	tmp := xhr.Value.Call("getAllResponseHeaders")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = tmp.String()
	}

	return result

}

func (xhr *XMLHttpRequest) GetResponseHeader(key string) string {

	var result string = ""

	tmp := xhr.Value.Call("getResponseHeader", js.ValueOf(key))

	if !tmp.IsNull() && !tmp.IsUndefined() {
		result = tmp.String()
	}

	return result

}

func (xhr *XMLHttpRequest) OpenWithUserAndPassword(method Method, url string, user string, password string) {

	if method.String() != "" && url != "" {
		xhr.Value.Call("open", method.String(), url, true, user, password)
	}

}

func (xhr *XMLHttpRequest) Open(method Method, url string) {

	if method.String() != "" && url != "" {
		xhr.Value.Call("open", method.String(), url, true)
	}

}

func (xhr *XMLHttpRequest) Send() {

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

		response := xhr.Value.Get("response")
		array := js.Global().Get("Uint8Array").New(response)
		xhr.Response = make([]byte, array.Get("byteLength").Int())
		js.CopyBytesToGo(xhr.Response, array)

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

	xhr.Value.Set("responseType", js.ValueOf("arraybuffer"))

	xhr.Value.Call("send", js.ValueOf(nil))

}

func (xhr *XMLHttpRequest) SendBytes(body []byte) {

	// TODO

}

func (xhr *XMLHttpRequest) SetRequestHeader(header string, value string) {
	xhr.Value.Call("setRequestHeader", header, value)
}
