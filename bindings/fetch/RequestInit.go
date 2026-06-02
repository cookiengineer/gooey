//go:build wasm

package fetch

import "context"
import "io"

type RequestInit struct {
	Method         Method            `json:"method"`
	Headers        map[string]string `json:"headers"`
	Body           io.Reader         `json:"body"`
	Mode           Mode              `json:"mode"`
	Credentials    Credentials       `json:"credentials"`
	Cache          Cache             `json:"cache"`
	Redirect       Redirect          `json:"redirect"`
	Referrer       string            `json:"referrer"`
	ReferrerPolicy ReferrerPolicy    `json:"referrerPolicy"`
	Integrity      string            `json:"integrity"`
	KeepAlive      bool              `json:"keepalive"`
	Signal         context.Context
}

func (options *RequestInit) MapToJS() map[string]any {

	result := make(map[string]any)

	if tmp := options.Method.String(); tmp != "" {
		result["method"] = options.Method.String()
	}

	if len(options.Headers) > 0 {

		result_headers := make(map[string]any)

		for key, val := range options.Headers {
			result_headers[key] = val
		}

	}

	if tmp := options.Mode.String(); tmp != "" {
		result["mode"] = options.Mode.String()
	}

	if tmp := options.Credentials.String(); tmp != "" {
		result["credentials"] = options.Credentials.String()
	}

	if tmp := options.Cache.String(); tmp != "default" && tmp != "" {
		result["cache"] = options.Cache.String()
	}

	if tmp := options.Redirect.String(); tmp != "follow" && tmp != "" {
		result["redirect"] = options.Redirect.String()
	}

	if options.Referrer != "" {
		result["referrer"] = options.Referrer
	}

	if tmp := options.ReferrerPolicy.String(); tmp != "" {
		result["referrerPolicy"] = options.ReferrerPolicy.String()
	}

	if options.Integrity != "" {
		result["integrity"] = options.Integrity
	}

	if options.KeepAlive != false {
		result["keepalive"] = options.KeepAlive
	}

	return result

}
