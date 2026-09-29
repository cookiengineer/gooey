//go:build wasm

package webgl

type CullFace uint

const (
	CullFaceBack         CullFace = 0x0405
	CullFaceFront        CullFace = 0x0404
	CullFaceFrontAndBack CullFace = 0x0408
)
