//go:build wasm

package fetch

type Method string

const (
	MethodCONNECT Method = "CONNECT"
	MethodDELETE  Method = "DELETE"
	MethodGET     Method = "GET"
	MethodHEAD    Method = "HEAD"
	MethodOPTIONS Method = "OPTIONS"
	MethodPATCH   Method = "PATCH"
	MethodPOST    Method = "POST"
	MethodPUT     Method = "PUT"
	MethodTRACE   Method = "TRACE"
)

func (method Method) String() string {
	return string(method)
}
