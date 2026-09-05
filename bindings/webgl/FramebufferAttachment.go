//go:build wasm

package webgl

type FramebufferAttachment uint

const (
	FramebufferAttachmentColor0       FramebufferAttachment = 0x8CE0
	FramebufferAttachmentColor1       FramebufferAttachment = 0x8CE1
	FramebufferAttachmentColor2       FramebufferAttachment = 0x8CE2
	FramebufferAttachmentColor3       FramebufferAttachment = 0x8CE3
	FramebufferAttachmentDepth        FramebufferAttachment = 0x8D00
	FramebufferAttachmentDepthStencil FramebufferAttachment = 0x821A
	FramebufferAttachmentStencil      FramebufferAttachment = 0x8D20
)
