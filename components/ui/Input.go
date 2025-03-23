package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

type Input struct {
	Label    string    `json:"label"`
	Type     InputType `json:"type"`
	Value    string    `json:"value"`
	Disabled bool      `json:"disabled"`
	Component
}

func NewInput(label string, value string, typ InputType) Input {

	var component Input

	element := bindings.Document.CreateElement("input")
	element.SetAttribute("type", typ.String())

	component.Label = label
	component.Type  = typ
	component.Value = strings.ToLower(value)

	component.Init(element)
	component.Render()

	return component

}

func ToInput(element *dom.Element) Input {

	var component Input

	component.Label    = strings.TrimSpace(element.GetAttribute("placeholder"))
	component.Value    = strings.ToLower(element.GetAttribute("value"))
	component.Disabled = element.HasAttribute("disabled")

	component.Init(element)

	return component

}

func (component *Input) Disable() {
	component.Disabled = true
	component.Render()
}

func (component *Input) Enable() {
	component.Disabled = false
	component.Render()
}

func (component *Input) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetAttribute("placeholder", component.Label)
		} else {
			component.Element.RemoveAttribute("placeholder")
		}

		component.Element.SetAttribute("type", component.Type.String())

		if component.Value != "" {
			component.Element.SetAttribute("value", component.Value)
		} else {
			component.Element.RemoveAttribute("value")
		}

		if component.Disabled == true {
			component.Element.SetAttribute("disabled", "")
		} else {
			component.Element.RemoveAttribute("disabled")
		}

	}

}

func (component *Input) String() string {

	html := "<input type=\"" + component.Type.String() + "\""

	if component.Label != "" {
		html += " placeholder=\"" + component.Label + "\""
	}

	if component.Value != "" {
		html += " value=\"" + component.Value + "\""
	}

	if component.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}
