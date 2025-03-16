//go:build wasm

package fetch

type Method string

const (
	MethodConnect Method = "CONNECT"
	MethodDelete  Method = "DELETE"
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodOptions Method = "OPTIONS"
	MethodPatch   Method = "PATCH"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodTrace   Method = "TRACE"
)

func (method Method) String() string {
	return string(method)
}
