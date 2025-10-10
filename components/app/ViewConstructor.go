//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/interfaces"

type ViewConstructor func(*dom.Element) interfaces.View
