//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/interfaces"

type Main struct {
	Element *dom.Element               `json:"element"`
	Client  *Client                    `json:"client"`
	Storage *Storage                   `json:"storage"`
	View    interfaces.View            `json:"view"`
	Views   map[string]interfaces.View `json:"views"`
}

func (main *Main) Init(element *dom.Element) {

	client := NewClient()
	storage := NewStorage()

	main.Element = element
	main.Client  = &client
	main.Storage = &storage
	main.View    = nil
	main.Views   = make(map[string]interfaces.View)

}

func (main *Main) SetView(name string, view interfaces.View) {

	main.Views[name] = view

}

func (main *Main) ChangeView(name string) bool {

	var result bool = false

	view, ok := main.Views[name]

	if ok == true {

		if main.View != nil {
			main.View.Leave()
			main.View = nil
		}

		main.Element.SetAttribute("data-view", name)

		main.View = view
		main.View.Enter()

		result = true

	}

	return result

}

