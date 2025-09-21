//go:build wasm

package components

import "github.com/cookiengineer/gooey/interfaces"

func Unwrap[Type any](component interfaces.Component) (Type, bool) {

    instance, ok := component.(Type)

    return instance, ok

}

