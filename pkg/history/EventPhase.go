//go:build wasm

package history

type EventPhase int

const (
	NONE            EventPhase = 0
	CAPTURING_PHASE EventPhase = 1
	AT_TARGET       EventPhase = 2
	BUBBLING_PHASE  EventPhase = 3
)

