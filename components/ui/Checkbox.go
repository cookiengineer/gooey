package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Checkbox struct {
	Checked   bool                  `json:"checked"`
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

	checkbox.Checked   = element.Value.Get("checked").Bool()
	checkbox.Component = &component
	checkbox.Label     = label
	checkbox.Value     = strings.ToLower(value)

	checkbox.Component.InitEvent("change")

	element.AddEventListener("change", dom.ToEventListener(func(_ dom.Event) {

		checkbox.Checked = element.Value.Get("checked").Bool()
		checkbox.Value   = element.Value.Get("value").String()

		checked := "false"

		if checkbox.Checked == true {
			checked = "true"
		}

		checkbox.Component.FireEventListeners("change", map[string]string{
			"checked": checked,
			"value":   checkbox.Value,
		})

	}))

	checkbox.Render()

	return checkbox

}

func ToCheckbox(element *dom.Element) Checkbox {

	var checkbox Checkbox

	component := components.NewComponent(element)

	checkbox.Checked   = element.Value.Get("checked").Bool()
	checkbox.Label     = strings.TrimSpace(element.GetAttribute("title"))
	checkbox.Value     = strings.ToLower(element.GetAttribute("value"))
	checkbox.Disabled  = element.HasAttribute("disabled")
	checkbox.Component = &component

	checkbox.Component.InitEvent("change")

	element.AddEventListener("change", dom.ToEventListener(func(_ dom.Event) {

		checkbox.Checked = element.Value.Get("checked").Bool()
		checkbox.Value   = element.Value.Get("value").String()

		checked := "false"

		if checkbox.Checked == true {
			checked = "true"
		}

		checkbox.Component.FireEventListeners("change", map[string]string{
			"checked": checked,
			"value":   checkbox.Value,
		})

	}))

	return checkbox

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

	if checkbox.Checked == true {
		html += " checked"
	}

	if checkbox.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}
