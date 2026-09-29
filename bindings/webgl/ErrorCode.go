//go:build wasm

package webgl

type ErrorCode uint

const (
	ErrorCodeContextLost                 ErrorCode = 0x9242
	ErrorCodeInvalidEnum                 ErrorCode = 0x0500
	ErrorCodeInvalidFramebufferOperation ErrorCode = 0x0506
	ErrorCodeInvalidOperation            ErrorCode = 0x0502
	ErrorCodeInvalidValue                ErrorCode = 0x0501
	ErrorCodeNoError                     ErrorCode = 0
	ErrorCodeOutOfMemory                 ErrorCode = 0x0505
)

// String returns the name the specification gives to the error code.
func (code ErrorCode) String() string {

	switch code {
	case ErrorCodeContextLost:
		return "CONTEXT_LOST_WEBGL"
	case ErrorCodeInvalidEnum:
		return "INVALID_ENUM"
	case ErrorCodeInvalidFramebufferOperation:
		return "INVALID_FRAMEBUFFER_OPERATION"
	case ErrorCodeInvalidOperation:
		return "INVALID_OPERATION"
	case ErrorCodeInvalidValue:
		return "INVALID_VALUE"
	case ErrorCodeNoError:
		return "NO_ERROR"
	case ErrorCodeOutOfMemory:
		return "OUT_OF_MEMORY"
	}

	return "UNKNOWN"

}
