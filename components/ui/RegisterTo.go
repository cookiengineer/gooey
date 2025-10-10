package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"

func RegisterTo(document *components.Document) {

	document.Register("button", components.WrapComponent(ToButton))
	document.Register("input", func(element *dom.Element) interfaces.Component {

		typ := element.GetAttribute("type")

		if typ == "checkbox" {
			return ToCheckbox(element)
		} else if typ == "number" {
			return ToNumber(element)
		} else if typ == "range" {
			return ToRange(element)
		} else if typ == "text" {
			return ToInput(element)
		} else {
			return ToInput(element)
		}

	})
	document.Register("label", components.WrapComponent(ToLabel))
	document.Register("select", components.WrapComponent(ToSelect))
	document.Register("textarea", components.WrapComponent(ToTextarea))

}
