//go:build wasm

package app

import "github.com/cookiengineer/gooey/components/interfaces"

func WrapController[Type interfaces.Controller](wrapped func(*Main, interfaces.View) Type) ControllerConstructor {

	return func(main *Main, view interfaces.View) interfaces.Controller {
		return wrapped(main, view)
	}

}
