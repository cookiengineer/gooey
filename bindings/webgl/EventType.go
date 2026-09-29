//go:build wasm

package webgl

type EventType string

const (
	EventTypeContextCreationError EventType = "webglcontextcreationerror"
	EventTypeContextLost          EventType = "webglcontextlost"
	EventTypeContextRestored      EventType = "webglcontextrestored"
)
