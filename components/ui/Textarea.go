package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Textarea struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
	components.Component
}

func NewTextarea(label string, value string) Textarea {

	var component Textarea

	element := bindings.Document.CreateElement("textarea")

	component.Label = label
	component.Value = value

	component.Init(element)
	component.Render()

	return component

}

func ToTextarea(element *dom.Element) Textarea {

	var component Textarea

	tmp := element.Value.Get("value")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		component.Value = strings.TrimSpace(tmp.String())
	} else {
		component.Value = ""
	}

	component.Label    = strings.TrimSpace(element.GetAttribute("placeholder"))
	component.Disabled = element.HasAttribute("disabled")

	component.Init(element)

	return component

}

func (component *Textarea) Disable() {
	component.Disabled = true
	component.Render()
}

func (component *Textarea) Enable() {
	component.Disabled = false
	component.Render()
}

func (component *Textarea) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetAttribute("placeholder", component.Label)
		} else {
			component.Element.RemoveAttribute("placeholder")
		}

		if component.Value != "" {
			component.Element.Value.Set("value", component.Value)
		} else {
			component.Element.Value.Set("value", "")
		}

		if component.Disabled == true {
			component.Element.SetAttribute("disabled", "")
		} else {
			component.Element.RemoveAttribute("disabled")
		}

	}

}

func (component *Textarea) String() string {

	html := "<textarea"

	if component.Label != "" {
		html += " placeholder=\"" + component.Label + "\""
	}

	if component.Disabled == true {
		html += " disabled"
	}

	html += ">"

	if component.Value != "" {
		html += component.Value
	}

	html += "</textarea>"

	return html

}
