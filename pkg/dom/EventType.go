//go:build wasm

package dom

type EventType string

const (
	EventTypeChange EventType = "change"
	EventTypeClick  EventType = "click"
)
