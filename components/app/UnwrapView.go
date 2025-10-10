//go:build wasm

package app

import "github.com/cookiengineer/gooey/components/interfaces"

func UnwrapView[Type any](view interfaces.View) (Type, bool) {

	instance, ok := view.(Type)

	return instance, ok

}
