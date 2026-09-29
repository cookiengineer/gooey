//go:build wasm

package webgl

type Capability uint

const (
	CapabilityBlend                 Capability = 0x0BE2
	CapabilityCullFace              Capability = 0x0B44
	CapabilityDepthTest             Capability = 0x0B71
	CapabilityDither                Capability = 0x0BD0
	CapabilityPolygonOffsetFill     Capability = 0x8037
	CapabilityRasterizerDiscard     Capability = 0x8C89
	CapabilitySampleAlphaToCoverage Capability = 0x809E
	CapabilitySampleCoverage        Capability = 0x80A0
	CapabilityScissorTest           Capability = 0x0C11
	CapabilityStencilTest           Capability = 0x0B90
)
