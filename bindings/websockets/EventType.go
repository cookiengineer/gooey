//go:build wasm

package websockets

type EventType string

const (
	EventTypeClose   EventType = "close"
	EventTypeError   EventType = "error"
	EventTypeMessage EventType = "message"
	EventTypeOpen    EventType = "open"
)
