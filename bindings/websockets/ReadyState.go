//go:build wasm

package websockets

type ReadyState int

const (
	ReadyStateConnecting ReadyState = 0
	ReadyStateOpen       ReadyState = 1
	ReadyStateClosing    ReadyState = 2
	ReadyStateClosed     ReadyState = 3
)
