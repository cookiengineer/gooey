//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/types"
import "strings"
import "syscall/js"

type Select struct {
	Label     string                `json:"label"`
	Type      types.Input           `json:"type"`
	Value     string                `json:"value"`
	Values    []string              `json:"values"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewSelect(label string, value string, values []string) Select {

	var self Select

	element   := bindings.Document.CreateElement("select")
	component := components.NewComponent(element)

	self.Component = &component
	self.Label     = strings.TrimSpace(label)
	self.Type      = types.InputText
	self.Value     = strings.TrimSpace(value)
	self.Values    = make([]string, 0)

	for _, val := range values {

		tmp := strings.TrimSpace(val)

		if tmp != "" {
			self.Values = append(self.Values, tmp)
		}

	}

	self.Component.InitEvent("change")

	self.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ dom.Event) {

		// TODO: Get Value from select element and set self.Value

		self.Component.FireEventListeners("change", map[string]string{
			"value": self.Value,
		})

	}))

	self.Render()

	return self

}

func ToSelect(element *dom.Element) Select {

	var self Select

	component := components.NewComponent(element)

	self.Component = &component
	// TODO: self.Label
	self.Disabled  = element.HasAttribute("disabled")

	// TODO: Value
	// TODO: Values

	self.Parse()

	self.Component.InitEvent("change")

	self.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ dom.Event) {

		// TODO: Get Value from select element and set self.Value

		self.Component.FireEventListeners("change", map[string]string{
			"value": self.Value,
		})

	}))

	return self

}

func (self *Select) Disable() bool {

	self.Disabled = true
	self.Render()

	return true

}

func (self *Select) Enable() bool {

	self.Disabled = false
	self.Render()

	return true

}

func (self *Select) Parse() {


}

func (self *Select) Render() *dom.Element {

	if self.Component.Element != nil {

		// TODO: Label property into an option value=""

		// TODO: Set active value via Element.Value.Set("value", ...)

		if self.Disabled == true {
			self.Component.Element.SetAttribute("disabled", "")
		} else {
			self.Component.Element.RemoveAttribute("disabled")
		}

	}

	return self.Component.Element

}

func (self *Select) String() string {

	html := "<select"

	if self.Disabled == true {
		html += " disabled"
	}

	html += ">"

	if self.Label != "" && self.Value == "" {
		html += "<option value=\"\">" + self.Label + "</option>"
	}

	if self.Value != "" {

		for _, value := range self.Values {

			if value == self.Value {
				html += "<option value=\"" + value + "\" selected>" + value + "</option>"
			} else {
				html += "<option value=\"" + value + "\">" + value + "</option>"
			}

		}
		
	} else {

		for _, value := range self.Values {
			html += "<option value=\"" + value + "\">" + value + "</option>"
		}

	}

	html += "</select>"

	return html

}

