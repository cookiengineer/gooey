package app

import "gooey/dom"

type BaseView struct {
	Elements map[string]*dom.Element `json:"elements"`
}

func (view BaseView) Init() {
	view.Elements = make(map[string]*dom.Element)
}

func (view BaseView) Enter() bool {
	return true
}

func (view BaseView) Leave() bool {
	return true
}

func (view BaseView) GetElement(id string) *dom.Element {

	var result *dom.Element = nil

	if id != "" {

		tmp, ok := view.Elements[id]

		if ok == true {
			result = tmp
		}
		
	}

	return result

}

func (view BaseView) SetElement(id string, element *dom.Element) {

	if id != "" && element != nil {
		view.Elements[id] = element
	}

}

func (view BaseView) RemoveElement(id string) bool {

	var result bool = false

	_, ok := view.Elements[id]

	if ok == true {
		delete(view.Elements, id)
		result = true
	}

	return result

}
