//go:build wasm

package app

import "gooey/dom"

type Main struct {
	Element *dom.Element     `json:"element"`
	Client  *Client          `json:"client"`
	Storage *Storage         `json:"storage"`
	View    View             `json:"view"`
	Views   map[string]View  `json:"views"`
}

func (main *Main) Init(element *dom.Element) {

	client := NewClient()
	storage := NewStorage()

	main.Element = element
	main.Client  = &client
	main.Storage = &storage
	main.View    = nil
	main.Views   = make(map[string]View)

}

func (main *Main) SetView(name string, view View) {

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

