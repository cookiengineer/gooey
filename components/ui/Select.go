//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "slices"
import "strings"
import "syscall/js"

type Select struct {
	Label         string                `json:"label"`
	Type          types.Input           `json:"type"`
	Value         string                `json:"value"`
	Values        []string              `json:"values"`
	Disabled      bool                  `json:"disabled"`
	Component     *components.Component `json:"component"`
	default_value string
}

func NewSelect(label string, value string, values []string) Select {

	var self Select

	element   := dom.Document.CreateElement("select")
	component := components.NewComponent(element)

	self.Component     = &component
	self.Label         = strings.TrimSpace(label)
	self.Type          = types.InputText
	self.Value         = strings.TrimSpace(value)
	self.Values        = make([]string, 0)
	self.default_value = self.Value

	for _, val := range values {

		tmp := strings.TrimSpace(val)

		if tmp != "" {
			self.Values = append(self.Values, tmp)
		}

	}

	self.Mount()
	self.Render()

	return self

}

func ToSelect(element *dom.Element) *Select {

	var self Select

	component := components.NewComponent(element)

	self.Component = &component
	self.Type      = types.InputText
	self.Disabled  = element.HasAttribute("disabled")

	self.Mount()

	return &self

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

func (self *Select) Mount() bool {

	if self.Component != nil {
		self.Component.InitEvent("change-value")
	}

	if self.Component.Element != nil {

		if len(self.Values) == 0 || self.default_value == "" {

			tmp := self.Component.Element.QuerySelector("option")

			// First option element is the placeholder
			if tmp != nil && tmp.GetAttribute("value") == "" {
				self.Label = tmp.TextContent
			} else {
				self.Label = ""
			}

			elements := self.Component.Element.QuerySelectorAll("option")

			value := ""
			values := make([]string, 0)

			for _, element := range elements {

				tmp := element.GetAttribute("value")

				if tmp != "" {

					if element.HasAttribute("selected") {
						value = tmp
					}

					values = append(values, tmp)

				}

			}

			self.default_value = value
			self.Value         = value
			self.Values        = values

		}

		self.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

			index   := self.Component.Element.Value.Get("selectedIndex").Int()
			options := self.Component.Element.Value.Get("options")

			if !options.IsNull() && !options.IsUndefined() && index != -1 {

				value := options.Index(index).Get("value")

				if !value.IsNull() && !value.IsUndefined() {

					if slices.Contains(self.Values, value.String()) {
						self.Value = value.String()
					}

					self.Component.FireEventListeners("change-value", map[string]any{
						"value": self.Value,
					})

				}

			} else if index == -1 {
				self.Value = ""
			}

		}))

		return true

	} else {
		return false
	}

}

func (self *Select) Query(query string) interfaces.Component {

	if self.Component.Element != nil {

		if utils.MatchesQuery(self.Component.Element, query) == true {
			return self.Component
		}

	}

	return nil

}

func (self *Select) Render() *dom.Element {

	if self.Component.Element != nil {

		if self.Disabled == true {
			self.Component.Element.SetAttribute("disabled", "")
		} else {
			self.Component.Element.RemoveAttribute("disabled")
		}

		elements := make(map[string]*dom.Element)

		tmp := self.Component.Element.QuerySelectorAll("option")

		for _, element := range tmp {

			value := element.GetAttribute("value")

			if slices.Contains(self.Values, value) {
				elements[value] = element
			}

		}

		children := make([]*dom.Element, 0)

		if self.Label != "" {

			placeholder := dom.Document.CreateElement("option")
			placeholder.SetAttribute("value", "")
			placeholder.SetInnerHTML(self.Label)

			children = append(children, placeholder)

		}

		for _, value := range self.Values {

			element, ok := elements[value]

			if ok == false {

				element = dom.Document.CreateElement("option")
				element.SetAttribute("value", value)
				element.SetInnerHTML(value)

				if self.Value != "" {

					if self.Value == value {
						element.SetAttribute("selected", "")
					} else {
						element.RemoveAttribute("selected")
					}

				}

				children = append(children, element)

			} else {

				if self.Value != "" {

					if self.Value == value {
						element.SetAttribute("selected", "")
					} else {
						element.RemoveAttribute("selected")
					}

				}

				children = append(children, element)

			}

		}

		self.Component.Element.ReplaceChildren(children)

	}

	return self.Component.Element

}

func (self *Select) Reset() bool {

	self.Value = self.default_value
	self.Render()

	return true

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

func (self *Select) ToValue() js.Value {

	var result js.Value

	if self.Component.Element != nil {

		index   := self.Component.Element.Value.Get("selectedIndex").Int()
		options := self.Component.Element.Value.Get("options")

		if !options.IsNull() && !options.IsUndefined() && index != -1 {

			value := options.Index(index).Get("value")

			if !value.IsNull() && !value.IsUndefined() {

				if slices.Contains(self.Values, value.String()) {
					result = value
				}

			}

		} else if index == -1 {
			result = js.ValueOf("")
		}

	}

	return result

}

func (self *Select) Unmount() bool {

	if self.Component.Element != nil {
		self.Component.Element.RemoveEventListener("change", nil)
	}

	return true

}
