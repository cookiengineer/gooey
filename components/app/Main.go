//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/location"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/interfaces"

type Main struct {
	Element *dom.Element               `json:"element"`
	Client  *Client                    `json:"client"`
	Header  *layout.Header             `json:"header"`
	Footer  *layout.Footer             `json:"footer"`
	Dialog  *layout.Dialog             `json:"dialog"`
	Storage *Storage                   `json:"storage"`
	View    interfaces.View            `json:"view"`
	views   map[string]interfaces.View `json:"-"`
}

func NewMain() *Main {

	var main Main

	main.Element = dom.Document.QuerySelector("main")
	main.Client  = NewClient()
	main.Storage = NewStorage()
	main.View    = nil
	main.views   = make(map[string]interfaces.View)

	if main.Element != nil {
		main.Parse()
	}

	return &main

}

func ToMain(element *dom.Element) *Main {

	var main Main

	main.Element = element
	main.Client  = NewClient()
	main.Storage = NewStorage()
	main.View    = nil
	main.views   = make(map[string]interfaces.View)

	if main.Element != nil {
		main.Parse()
	}

	return &main

}

func (main *Main) Parse() {

	header_element := dom.Document.QuerySelector("body > header")

	if header_element != nil {

		main.Header = layout.ToHeader(header_element)
		main.Header.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]any) {

			name, ok1 := attributes["name"].(string)
			path, ok2 := attributes["path"].(string)

			if ok1 == true && ok2 == true {

				_, ok3 := main.views[name]

				if ok3 == true {

					// Single-page web app
					main.ChangeView(name)

				} else {

					// TODO: History API integration?
					// Multi-page web app
					location.Location.Replace(path)

				}

			}

		}, false))

	} else {
		main.Header = nil
	}

	footer_element := dom.Document.QuerySelector("body > footer")

	if footer_element != nil {
		main.Footer = layout.ToFooter(footer_element)
	} else {
		main.Footer = nil
	}

	dialog_element := dom.Document.QuerySelector("body > dialog")

	if dialog_element != nil {
		main.Dialog = layout.ToDialog(dialog_element)
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

	var result bool

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
