//go:build wasm

package canvas2d

type Repetition string

const (
	RepetitionRepeat   Repetition = "repeat"
	RepetitionRepeatX  Repetition = "repeat-x"
	RepetitionRepeatY  Repetition = "repeat-y"
	RepetitionNoRepeat Repetition = "no-repeat"
)

