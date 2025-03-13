//go:build wasm

package fetch

type Mode string

const (
	ModeSameOrigin Mode = "same-origin"
	ModeNoCORS     Mode = "no-cors"
	ModeCORS       Mode = "cors"
	ModeNavigate   Mode = "navigate"
)

func (mode Mode) String() string {
	return string(mode)
}
