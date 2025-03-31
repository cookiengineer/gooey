package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Textarea struct {
	Label     string                `json:"label"`
	Value     string                `json:"value"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewTextarea(label string, value string) Textarea {

	var textarea Textarea

	element   := bindings.Document.CreateElement("textarea")
	component := components.NewComponent(element)

	textarea.Component = &component
	textarea.Label     = label
	textarea.Value     = value

	textarea.Component.InitEvent("change")
	textarea.Render()

	return textarea

}

func ToTextarea(element *dom.Element) Textarea {

	var textarea Textarea

	tmp := element.Value.Get("value")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		textarea.Value = strings.TrimSpace(tmp.String())
	} else {
		textarea.Value = ""
	}

	component := components.NewComponent(element)

	textarea.Component = &component
	textarea.Label     = strings.TrimSpace(element.GetAttribute("placeholder"))
	textarea.Disabled  = element.HasAttribute("disabled")

	textarea.Component.InitEvent("change")

	return textarea

}

func (textarea *Textarea) Disable() bool {

	textarea.Disabled = true
	textarea.Render()

	return true

}

func (textarea *Textarea) Enable() bool {

	textarea.Disabled = false
	textarea.Render()

	return true

}

func (textarea *Textarea) Render() *dom.Element {

	if textarea.Component.Element != nil {

		if textarea.Label != "" {
			textarea.Component.Element.SetAttribute("placeholder", textarea.Label)
		} else {
			textarea.Component.Element.RemoveAttribute("placeholder")
		}

		if textarea.Value != "" {
			textarea.Component.Element.Value.Set("value", textarea.Value)
		} else {
			textarea.Component.Element.Value.Set("value", "")
		}

		if textarea.Disabled == true {
			textarea.Component.Element.SetAttribute("disabled", "")
		} else {
			textarea.Component.Element.RemoveAttribute("disabled")
		}

	}

	return textarea.Component.Element

}

func (textarea *Textarea) String() string {

	html := "<textarea"

	if textarea.Label != "" {
		html += " placeholder=\"" + textarea.Label + "\""
	}

	if textarea.Disabled == true {
		html += " disabled"
	}

	html += ">"

	if textarea.Value != "" {
		html += textarea.Value
	}

	html += "</textarea>"

	return html

}
