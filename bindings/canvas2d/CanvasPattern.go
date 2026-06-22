//go:build wasm

package canvas2d

type CanvasPattern struct {
	Image      *Image     `json:"image"`
	Repetition Repetition `json:"repetition"`
}
