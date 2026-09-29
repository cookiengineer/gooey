//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/interfaces"

func ReconcileComponents(element *dom.Element, content []interfaces.Component) {

	elements := make([]*dom.Element, 0)

	for _, component := range content {

		if component == nil {
			continue
		}

		rendered := component.Render()

		if rendered != nil {
			elements = append(elements, rendered)
		}

	}

	ReconcileElements(element, elements)

}

