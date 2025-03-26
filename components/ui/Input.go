package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Input struct {
	Label     string                `json:"label"`
	Type      InputType             `json:"type"`
	Value     string                `json:"value"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewInput(label string, value string, typ InputType) Input {

	var input Input

	element   := bindings.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", typ.String())

	input.Component = &component
	input.Label     = strings.TrimSpace(label)
	input.Type      = typ
	input.Value     = strings.TrimSpace(value)

	input.Component.InitEvent("change")
	input.Render()

	return input

}

func ToInput(element *dom.Element) Input {

	var input Input

	tmp := element.Value.Get("value")

	if !tmp.IsNull() && !tmp.IsUndefined() {
		input.Value = strings.TrimSpace(tmp.String())
	} else {
		input.Value = ""
	}

	component := components.NewComponent(element)

	input.Component = &component
	input.Label     = strings.TrimSpace(element.GetAttribute("placeholder"))
	input.Disabled  = element.HasAttribute("disabled")

	input.Component.InitEvent("change")

	return input

}

func (input *Input) Disable() {
	input.Disabled = true
	input.Render()
}

func (input *Input) Enable() {
	input.Disabled = false
	input.Render()
}

func (input *Input) Render() {

	if input.Component.Element != nil {

		if input.Label != "" {
			input.Component.Element.SetAttribute("placeholder", input.Label)
		} else {
			input.Component.Element.RemoveAttribute("placeholder")
		}

		input.Component.Element.SetAttribute("type", input.Type.String())

		if input.Value != "" {
			input.Component.Element.Value.Set("value", input.Value)
		} else {
			input.Component.Element.Value.Set("value", "")
		}

		if input.Disabled == true {
			input.Component.Element.SetAttribute("disabled", "")
		} else {
			input.Component.Element.RemoveAttribute("disabled")
		}

	}

}

func (input *Input) String() string {

	html := "<input type=\"" + input.Type.String() + "\""

	if input.Label != "" {
		html += " placeholder=\"" + input.Label + "\""
	}

	if input.Value != "" {
		html += " value=\"" + input.Value + "\""
	}

	if input.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}
