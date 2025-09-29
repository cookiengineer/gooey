//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
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

	checkbox.Mount()
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

	checkbox.Mount()

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

func (checkbox *Checkbox) Mount() bool {

	if checkbox.Component != nil {
		checkbox.Component.InitEvent("change-value")
	}

	if checkbox.Component.Element != nil {

		checkbox.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

			checkbox.Value = checkbox.Component.Element.Value.Get("checked").Bool()

			checked := "false"

			if checkbox.Value == true {
				checked = "true"
			}

			checkbox.Component.FireEventListeners("change-value", map[string]any{
				"value": checked,
			})

		}))

		return true

	} else {
		return false
	}

}

func (checkbox *Checkbox) Query(query string) interfaces.Component {

	if checkbox.Component.Element != nil {

		if utils.MatchesQuery(checkbox.Component.Element, query) == true {
			return checkbox.Component
		}

	}

	return nil

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

func (checkbox *Checkbox) Reset() bool {

	checkbox.Value = false
	checkbox.Render()

	return true

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

func (checkbox *Checkbox) Unmount() bool {

	if checkbox.Component.Element != nil {
		checkbox.Component.Element.RemoveEventListener("change", nil)
	}

	return true

}
