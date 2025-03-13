//go:build wasm

package gooey

type ScreenOrientation struct {
	Angle uint   `json:"angle"`
	Type  string `json:"type"`
}
