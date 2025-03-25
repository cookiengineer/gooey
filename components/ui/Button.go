package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Button struct {
	Label    string `json:"label"`
	Action   string `json:"action"`
	Disabled bool   `json:"disabled"`
	components.Component
}

func NewButton(label string, action string) Button {

	var component Button

	element := bindings.Document.CreateElement("button")

	component.Label  = label
	component.Action = strings.ToLower(action)

	component.Init(element)
	component.Render()

	return component

}

func ToButton(element *dom.Element) Button {

	var component Button

	component.Label    = strings.TrimSpace(element.TextContent)
	component.Action   = strings.ToLower(element.GetAttribute("data-action"))
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

func (component *Button) Render() {

	if component.Element != nil {

		if component.Label != "" {
			component.Element.SetInnerHTML(component.Label)
		} else {
			component.Element.SetInnerHTML("")
		}

		if component.Action != "" {
			component.Element.SetAttribute("data-action", component.Action)
		} else {
			component.Element.RemoveAttribute("data-action")
		}

		if component.Disabled == true {
			component.Element.SetAttribute("disabled", "")
		} else {
			component.Element.RemoveAttribute("disabled")
		}

	}

}

func (component *Button) String() string {

	html := "<button"

	if component.Action != "" {
		html += " data-action=\"" + component.Action + "\""
	}

	if component.Disabled == true {
		html += " disabled"
	}

	html += ">"

	if component.Label != "" {
		html += component.Label
	}

	html += "</button>"

	return html

}
