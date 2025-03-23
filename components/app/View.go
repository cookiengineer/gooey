//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"

type View struct {
	Elements map[string]*dom.Element `json:"elements"`
}

func (view View) Init() {
	view.Elements = make(map[string]*dom.Element)
}

func (view View) Enter() bool {
	return true
}

func (view View) Leave() bool {
	return true
}

func (view View) GetElement(id string) *dom.Element {

	var result *dom.Element = nil

	if id != "" {

		tmp, ok := view.Elements[id]

		if ok == true {
			result = tmp
		}
		
	}

	return result

}

func (view View) SetElement(id string, element *dom.Element) {

	if id != "" && element != nil {
		view.Elements[id] = element
	}

}

func (view View) RemoveElement(id string) bool {

	var result bool = false

	_, ok := view.Elements[id]

	if ok == true {
		delete(view.Elements, id)
		result = true
	}

	return result

}
