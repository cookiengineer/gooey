package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

type Checkbox struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
	Component
}

func NewCheckbox(label string, value string) Checkbox {

	var component Checkbox

	element := bindings.Document.CreateElement("input")
	element.SetAttribute("type", "checkbox")

	component.Label = label
	component.Value = strings.ToLower(value)

	component.Init(element)
	component.Render()

	return component

}

func ToCheckbox(element *dom.Element) Checkbox {

	var component Checkbox

	component.Label    = strings.TrimSpace(element.GetAttribute("title"))
	component.Value    = strings.ToLower(element.GetAttribute("value"))
	component.Disabled = element.HasAttribute("disabled")

	component.Init(element)

	return component

}

func (component *Button) Disable() {
	component.Disabled = true
	component.Render()
}

func (component *Button) Enable() {
	component.Disabled = false
	component.Render()
}

func (component *Checkbox) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetAttribute("title", component.Label)
		} else {
			component.Element.RemoveAttribute("title")
		}

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

func (component *Checkbox) String() string {

	html := "<input type=\"checkbox\""

	if component.Label != "" {
		html += " title=\"" + component.Label + "\""
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
