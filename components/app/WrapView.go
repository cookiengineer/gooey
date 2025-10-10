//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/interfaces"

func WrapView[Type interfaces.View](wrapped func(*dom.Element) Type) ViewConstructor {

	return func(element *dom.Element) interfaces.View {
		return wrapped(element)
	}

}
