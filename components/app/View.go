//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"

type View struct {
	Element    *dom.Element                    `json:"element"`
	Layout     types.Layout                    `json:"layout"`
	Name       string                          `json:"name"`
	Label      string                          `json:"label"`
	Path       string                          `json:"path"`
	Components map[string]interfaces.Component `json:"components"`
}

func NewView(name string, label string, path string) View {

	var view View

	element := bindings.Document.CreateElement("section")
	element.SetAttribute("data-layout", types.LayoutFlow.String())

	view.Element = element
	view.Name    = strings.ToLower(name)
	view.Label   = label
	view.Layout  = types.LayoutFlow
	view.Path    = strings.ToLower(path)

	view.Components = make(map[string]interfaces.Component)

	return view

}

func (view *View) Enter() bool {
	return true
}

func (view *View) Leave() bool {
	return true
}

func (view *View) GetComponent(name string) interfaces.Component {

	var result interfaces.Component = nil

	if name != "" {

		tmp, ok := view.Components[name]

		if ok == true {
			result = tmp
		}

	}

	return result

}

func (view *View) GetProperty(name string) string {

	var result string

	switch name {
	case "Name":
		result = view.Name
	case "Label":
		result = view.Label
	case "Path":
		result = view.Path
	}

	return result

}

func (view *View) RemoveComponent(name string) bool {

	var result bool = false

	_, ok := view.Components[name]

	if ok == true {
		delete(view.Components, name)
		result = true
	}

	return result

}

func (view *View) Render() {

	for _, component := range view.Components {
		component.Render()
	}

}

func (view *View) SetComponent(name string, component interfaces.Component) {

	if name != "" {
		view.Components[name] = component
	}

}

func (view *View) SetProperty(name string, value string) bool {

	var result bool

	switch name {
	case "Name":
		view.Name = value
		result    = true
	case "Label":
		view.Label = value
		result     = true
	case "Path":
		view.Path = value
		result    = true
	}

	return result

}

