//go:build wasm

package history

type EventType string

const (
	EventTypePopstate EventType = "popstate"
)
