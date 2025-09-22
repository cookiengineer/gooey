package ui

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("button",                   components.Wrap(ToButton))
	document.Register("input[type=\"checkbox\"]", components.Wrap(ToCheckbox))
	document.Register("input[type=\"number\"]",   components.Wrap(ToNumber))
	document.Register("input[type=\"range\"]",    components.Wrap(ToRange))
	document.Register("input",                    components.Wrap(ToInput))
	document.Register("label",                    components.Wrap(ToLabel))
	document.Register("select",                   components.Wrap(ToSelect))
	document.Register("textarea",                 components.Wrap(ToTextarea))

}
