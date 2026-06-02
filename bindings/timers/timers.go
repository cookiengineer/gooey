//go:build wasm

package timers

type Duration uint64

const (
	Millisecond Duration = 1
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
