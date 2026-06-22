//go:build wasm

package canvas2d

import "syscall/js"

type CanvasPattern struct {
	Image      *Image     `json:"image"`
	Repetition Repetition `json:"repetition"`
	Value      *js.Value  `json:"value"`
}
