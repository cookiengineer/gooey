//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

type View struct {
	Name     string                  `json:"name"`
	Label    string                  `json:"label"`
	Path     string                  `json:"path"`
	Elements map[string]*dom.Element `json:"elements"`
}

func (view View) Init(name string, label string, path string) {

	view.Name = strings.ToLower(name)
	view.Label = label

	if strings.HasPrefix(path, "/") && strings.HasSuffix(path, ".html") {
		view.Path = strings.ToLower(path)
	} else {
		view.Path = "/" + view.Name + ".html"
	}

	view.Elements = make(map[string]*dom.Element)

}

func (view View) Properties() (string, string, string) {
	return view.Name, view.Label, view.Path
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

func (view View) Render() {
	// Render into dom Element
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
