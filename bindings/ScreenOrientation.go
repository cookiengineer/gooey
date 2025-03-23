//go:build wasm

package bindings

type ScreenOrientation struct {
	Angle uint   `json:"angle"`
	Type  string `json:"type"`
}
