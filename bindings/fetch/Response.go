//go:build wasm

package fetch

import "syscall/js"

type Response struct {
	Headers    Headers `json:"headers"`
	OK         bool    `json:"ok"`
	Redirected bool    `json:"redirected"`
	Status     int     `json:"status"`
	StatusText string  `json:"statusText"`
	Type       string  `json:"type"`
	URL        string  `json:"url"`
	Body       []byte  `json:"body"`
	BodyUsed   bool    `json:"bodyUsed"`
	Value      *js.Value
}

func ToResponse(value js.Value) *Response {

	var response Response

	headers := Headers{}
	iterator := value.Get("headers").Call("entries")

	for {

		next := iterator.Call("next")

		if next.Get("done").Bool() {
			break
		} else {

			tmp := next.Get("value")
			headers.Append(tmp.Index(0).String(), tmp.Index(1).String())

		}

	}

	response.Headers = headers
	response.OK = value.Get("ok").Bool()
	response.Redirected = value.Get("redirected").Bool()
	response.Status = value.Get("status").Int()
	response.StatusText = value.Get("statusText").String()
	response.Type = value.Get("type").String()
	response.URL = value.Get("url").String()
	response.Body = []byte{}
	response.BodyUsed = value.Get("bodyUsed").Bool()
	response.Value = &value

	return &response

}
