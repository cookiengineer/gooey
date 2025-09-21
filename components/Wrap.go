//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/interfaces"

type constructor func(*dom.Element) interfaces.Component

func Wrap[Type interfaces.Component](wrapped func(*dom.Element) Type) constructor {

    return func(element *dom.Element) interfaces.Component {
        return wrapped(element)
    }

}

