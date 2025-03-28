package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Checkbox struct {
	Label     string                `json:"label"`
	Value     string                `json:"value"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewCheckbox(label string, value string) Checkbox {

	var checkbox Checkbox

	element   := bindings.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", "checkbox")

	checkbox.Component = &component
	checkbox.Label     = label
	checkbox.Value     = strings.ToLower(value)

	checkbox.Component.InitEvent("change")
	checkbox.Render()

	return checkbox

}

func ToCheckbox(element *dom.Element) Checkbox {

	var checkbox Checkbox

	component := components.NewComponent(element)

	checkbox.Label     = strings.TrimSpace(element.GetAttribute("title"))
	checkbox.Value     = strings.ToLower(element.GetAttribute("value"))
	checkbox.Disabled  = element.HasAttribute("disabled")
	checkbox.Component = &component

	checkbox.Component.InitEvent("change")

	return checkbox

}

func (checkbox *Checkbox) Disable() {
	checkbox.Disabled = true
	checkbox.Render()
}

func (checkbox *Checkbox) Enable() {
	checkbox.Disabled = false
	checkbox.Render()
}

func (checkbox *Checkbox) Render() *dom.Element {

	if checkbox.Component.Element != nil {

		if checkbox.Label != "" {
			checkbox.Component.Element.SetAttribute("title", checkbox.Label)
		} else {
			checkbox.Component.Element.RemoveAttribute("title")
		}

		if checkbox.Value != "" {
			checkbox.Component.Element.SetAttribute("value", checkbox.Value)
		} else {
			checkbox.Component.Element.RemoveAttribute("value")
		}

		if checkbox.Disabled == true {
			checkbox.Component.Element.SetAttribute("disabled", "")
		} else {
			checkbox.Component.Element.RemoveAttribute("disabled")
		}

	}

	return checkbox.Component.Element

}

func (checkbox *Checkbox) String() string {

	html := "<input type=\"checkbox\""

	if checkbox.Label != "" {
		html += " title=\"" + checkbox.Label + "\""
	}

	if checkbox.Value != "" {
		html += " value=\"" + checkbox.Value + "\""
	}

	if checkbox.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}
