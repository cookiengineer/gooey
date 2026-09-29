//go:build wasm

package webgl

type FrontFace uint

const (
	FrontFaceClockwise        FrontFace = 0x0900
	FrontFaceCounterClockwise FrontFace = 0x0901
)
