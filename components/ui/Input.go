//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"
import "syscall/js"

type Input struct {
	Label     string                `json:"label"`
	Type      types.Input           `json:"type"`
	Value     string                `json:"value"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewInput(label string, value string, typ types.Input) Input {

	var input Input

	element   := dom.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", typ.String())

	input.Component = &component
	input.Label     = strings.TrimSpace(label)
	input.Type      = typ
	input.Value     = strings.TrimSpace(value)
	input.Disabled  = false

	input.Mount()
	input.Render()

	return input

}

func ToInput(element *dom.Element) *Input {

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
	input.Type      = types.Input(element.GetAttribute("type"))
	input.Disabled  = element.HasAttribute("disabled")

	input.Mount()

	return &input

}

func (input *Input) Disable() bool {

	input.Disabled = true
	input.Render()

	return true

}

func (input *Input) Enable() bool {

	input.Disabled = false
	input.Render()

	return true

}

func (input *Input) Mount() bool {

	input.Component.InitEvent("change-value")

	if input.Component.Element != nil {

		input.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

			input.Value = input.Component.Element.Value.Get("value").String()

			input.Component.FireEventListeners("change-value", map[string]any{
				"value": input.Value,
			})

		}))

		return true

	} else {
		return false
	}

}

func (input *Input) Query(query string) interfaces.Component {
	return nil
}

func (input *Input) Render() *dom.Element {

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

	return input.Component.Element

}

func (input *Input) Reset() bool {

	input.Value = ""
	input.Render()

	return true

}

func (input *Input) SetChildren(children []interfaces.Component) bool {
	return false
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

func (input *Input) ToValue() js.Value {

	var result js.Value

	if input.Component.Element != nil {

		tmp := input.Component.Element.Value.Get("value")

		if !tmp.IsNull() && !tmp.IsUndefined() {
			result = tmp
		}

	}

	return result

}

func (input *Input) Unmount() bool {

	if input.Component.Element != nil {
		input.Component.Element.RemoveEventListener("change", nil)
	}

	return true

}
