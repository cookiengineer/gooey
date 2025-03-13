//go:build wasm

package fetch

import "context"
import "io"

type Request struct {
	Method         Method            `json:"method"`
	Headers        map[string]string `json:"headers"`
	Body           io.Reader         `json:"body"`
	Mode           Mode              `json:"mode"`
	Credentials    Credentials       `json:"credentials"`
	Cache          Cache             `json:"cache"`
	Redirect       Redirect          `json:"redirect"`
	Referrer       string            `json:"referrer"`
	ReferrerPolicy string            `json:"referrerPolicy"`
	Integrity      string            `json:"integrity"`
	KeepAlive      bool              `json:"keepalive"`
	Signal         context.Context
}

func (request *Request) MapToJS() map[string]any {

	result := make(map[string]any)

	if tmp := request.Method.String(); tmp != "" {
		result["method"] = request.Method.String()
	}

	if len(request.Headers) > 0 {

		result_headers := make(map[string]any)

		for key, val := range request.Headers {
			result_headers[key] = val
		}

	}

	if tmp := request.Mode.String(); tmp != "" {
		result["mode"] = request.Mode.String()
	}

	if tmp := request.Credentials.String(); tmp != "" {
		result["credentials"] = request.Credentials.String()
	}

	if tmp := request.Cache.String(); tmp != "default" && tmp != "" {
		result["cache"] = request.Cache.String()
	}

	if tmp := request.Redirect.String(); tmp != "follow" && tmp != "" {
		result["redirect"] = request.Redirect.String()
	}

	if request.Referrer != "" {
		result["referrer"] = request.Referrer
	}

	if request.ReferrerPolicy != "" {
		result["referrerPolicy"] = request.ReferrerPolicy
	}

	if request.Integrity != "" {
		result["integrity"] = request.Integrity
	}

	if request.KeepAlive != false {
		result["keepalive"] = request.KeepAlive
	}

	return result

}
