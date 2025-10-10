//go:build wasm

package app

import "github.com/cookiengineer/gooey/components/interfaces"

func UnwrapController[Type any](controller interfaces.Controller) (Type, bool) {

	instance, ok := controller.(Type)

	return instance, ok

}
