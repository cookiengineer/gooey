package layout

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Dialog struct {
	Layout  types.Layout           `json:"layout"`
	Title   string                 `json:"title"`
	Content []interfaces.Component `json:"content"`
	Footer  struct {
		Left  []interfaces.Component `json:"left"`
		Right []interfaces.Component `json:"right"`
	} `json:"footer"`
	Component *components.Component `json:"component"`
}

func NewDialog() Dialog {

	var dialog Dialog

	// TODO

	return dialog

}

func ToDialog(element *dom.Element) Dialog {

	var dialog Dialog

	// TODO

	return dialog

}

func (dialog *Dialog) Hide() bool {

	var result bool

	if dialog.Component.Element != nil {

		dialog.Component.Element.RemoveAttribute("open")
		result = true

	}

	return result

}

func (dialog *Dialog) Parse() {

	// TODO

}

func (dialog *Dialog) Render() *dom.Element {

	if dialog.Component.Element != nil {

		// TODO

	}

	return dialog.Component.Element

}

func (dialog *Dialog) SetTitle(value string) {

	// TODO: set title

}

func (dialog *Dialog) Show() bool {

	var result bool

	if dialog.Component.Element != nil {

		dialog.Component.Element.SetAttribute("open", "")
		result = true

	}

	return result

}
