//go:build wasm

package canvas2d

import "syscall/js"

type GradientType string

const (
	GradientTypeLinear GradientType = "linear"
	GradientTypeConic  GradientType = "conic"
	GradientTypeRadial GradientType = "radial"
)

type CanvasGradient struct {
	Type  GradientType `json:"type"`
	Value *js.Value    `json:"value"`
}

func (gradient *CanvasGradient) AddColorStop(offset float64, color *Color) {

	if offset >= 0.0 && offset <= 1.0 && color != nil {
		gradient.Value.Call("addColorStop", offset, color.String())
	}

}
