//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/interfaces"

type Main struct {
	Element *dom.Element    `json:"element"`
	Client  *Client         `json:"client"`
	Header  *layout.Header  `json:"header"`
	Footer  *layout.Footer  `json:"footer"`
	Dialog  *layout.Dialog  `json:"dialog"`
	Storage *Storage        `json:"storage"`
	View    interfaces.View `json:"view"`
	views   map[string]interfaces.View
}

func (main *Main) Init(element *dom.Element) {

	client := NewClient()
	storage := NewStorage()

	main.Element = element
	main.Client  = &client
	main.Storage = &storage
	main.View    = nil
	main.views   = make(map[string]interfaces.View)

	header_element := bindings.Document.QuerySelector("body > header")
	footer_element := bindings.Document.QuerySelector("body > footer")
	dialog_element := bindings.Document.QuerySelector("body > dialog")

	if header_element != nil {

		header := layout.ToHeader(header_element)
		main.Header = &header

		main.Header.Component.AddEventListener("change-view", components.ToComponentListener(func(event string, attributes map[string]string) {

			name, ok := attributes["name"]

			if ok == true {
				main.ChangeView(name)
			}

		}, false))

	} else {
		main.Header = nil
	}

	if footer_element != nil {

		footer := layout.ToFooter(footer_element)
		main.Footer = &footer

	} else {
		main.Footer = nil
	}

	if dialog_element != nil {
		dialog := layout.ToDialog(dialog_element)
		main.Dialog = &dialog
	} else {
		main.Dialog = nil
	}

}

func (main *Main) Render() {

	if main.Header != nil {
		main.Header.Render()
	}

	if main.View != nil {
		main.View.Render()
	}

	if main.Footer != nil {
		main.Footer.Render()
	}

	if main.Dialog != nil {
		main.Dialog.Render()
	}

}

func (main *Main) ChangeView(name string) bool {

	var result bool = false

	view, ok := main.views[name]

	if ok == true {

		if main.View != nil {
			main.View.Leave()
			main.View = nil
		}

		main.Element.SetAttribute("data-view", name)

		if main.Header != nil {
			main.Header.ChangeView(name)
		}

		main.View = view
		main.View.Enter()

		result = true

	}

	return result

}

func (main *Main) SetView(view interfaces.View) {

	name := view.GetProperty("Name")

	if name != "" {

		main.views[name] = view

		if main.Header != nil {
			main.Header.SetView(view)
		}

	}

}
