//go:build wasm

package webgl

type FramebufferStatus uint

const (
	FramebufferStatusComplete                    FramebufferStatus = 0x8CD5
	FramebufferStatusIncompleteAttachment        FramebufferStatus = 0x8CD6
	FramebufferStatusIncompleteDimensions        FramebufferStatus = 0x8CD9
	FramebufferStatusIncompleteMissingAttachment FramebufferStatus = 0x8CD7
	FramebufferStatusIncompleteMultisample       FramebufferStatus = 0x8D56
	FramebufferStatusUnsupported                 FramebufferStatus = 0x8CDD
)

// String returns the name the specification gives to the status.
func (status FramebufferStatus) String() string {

	switch status {
	case FramebufferStatusComplete:
		return "FRAMEBUFFER_COMPLETE"
	case FramebufferStatusIncompleteAttachment:
		return "FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case FramebufferStatusIncompleteDimensions:
		return "FRAMEBUFFER_INCOMPLETE_DIMENSIONS"
	case FramebufferStatusIncompleteMissingAttachment:
		return "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case FramebufferStatusIncompleteMultisample:
		return "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	case FramebufferStatusUnsupported:
		return "FRAMEBUFFER_UNSUPPORTED"
	}

	return "UNKNOWN"

}
