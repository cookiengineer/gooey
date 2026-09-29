//go:build wasm

package webgl

type PowerPreference string

const (
	PowerPreferenceDefault         PowerPreference = "default"
	PowerPreferenceHighPerformance PowerPreference = "high-performance"
	PowerPreferenceLowPower        PowerPreference = "low-power"
)
