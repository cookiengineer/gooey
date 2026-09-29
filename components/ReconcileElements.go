//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/virtual"

func ReconcileElements(element *dom.Element, elements []*dom.Element) {

	if element == nil {
		return
	}

	virtual.Reconcile(virtual.ToElement(element), virtual.ToElements(elements))

}
