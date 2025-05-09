//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/types"
import "strings"
import "syscall/js"

type Checkbox struct {
	Label     string                `json:"label"`
	Type      types.Input           `json:"type"`
	Value     bool                  `json:"value"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewCheckbox(label string, value string) Checkbox {

	var checkbox Checkbox

	element   := dom.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", "checkbox")

	checkbox.Component = &component
	checkbox.Label     = label
	checkbox.Type      = types.InputCheckbox
	checkbox.Value     = false
	checkbox.Disabled  = false

	checkbox.Component.InitEvent("change-value")

	checkbox.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

		checkbox.Value = element.Value.Get("checked").Bool()

		checked := "false"

		if checkbox.Value == true {
			checked = "true"
		}

		checkbox.Component.FireEventListeners("change-value", map[string]string{
			"value": checked,
		})

	}))

	checkbox.Render()

	return checkbox

}

func ToCheckbox(element *dom.Element) *Checkbox {

	var checkbox Checkbox

	component := components.NewComponent(element)

	checkbox.Component = &component
	checkbox.Label     = strings.TrimSpace(element.GetAttribute("title"))
	checkbox.Type      = types.InputCheckbox
	checkbox.Value     = element.Value.Get("checked").Bool()
	checkbox.Disabled  = element.HasAttribute("disabled")

	checkbox.Component.InitEvent("change-value")

	checkbox.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

		checkbox.Value = element.Value.Get("checked").Bool()

		checked := "false"

		if checkbox.Value == true {
			checked = "true"
		}

		checkbox.Component.FireEventListeners("change-value", map[string]string{
			"value": checked,
		})

	}))

	return &checkbox

}

func (checkbox *Checkbox) Disable() bool {

	checkbox.Disabled = true
	checkbox.Render()

	return true

}

func (checkbox *Checkbox) Enable() bool {

	checkbox.Disabled = false
	checkbox.Render()

	return true

}

func (checkbox *Checkbox) Render() *dom.Element {

	if checkbox.Component.Element != nil {

		if checkbox.Label != "" {
			checkbox.Component.Element.SetAttribute("title", checkbox.Label)
		} else {
			checkbox.Component.Element.RemoveAttribute("title")
		}

		if checkbox.Value == true {
			checkbox.Component.Element.Value.Set("checked", true)
		} else {
			checkbox.Component.Element.Value.Set("checked", false)
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

	if checkbox.Value == true {
		html += " checked"
	}

	if checkbox.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}

func (checkbox *Checkbox) ToValue() js.Value {

	var result js.Value

	if checkbox.Component.Element != nil {

		tmp := checkbox.Component.Element.Value.Get("checked")

		if !tmp.IsNull() && !tmp.IsUndefined() {

			if tmp.Bool() == true {
				result = js.ValueOf(true)
			} else {
				result = js.ValueOf(false)
			}

		}

	}

	return result

}
