//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/interfaces"

func WrapComponent[Type interfaces.Component](wrapped func(*dom.Element) Type) ComponentConstructor {

	return func(element *dom.Element) interfaces.Component {
		return wrapped(element)
	}

}
