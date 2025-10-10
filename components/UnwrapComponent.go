//go:build wasm

package components

import "github.com/cookiengineer/gooey/components/interfaces"

func UnwrapComponent[Type any](component interfaces.Component) (Type, bool) {

	instance, ok := component.(Type)

	return instance, ok

}
