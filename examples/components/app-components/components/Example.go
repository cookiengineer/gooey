package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Example struct {
	Name      string                 `json:"name"`
	Layout    types.Layout           `json:"layout"`
	Content   []interfaces.Component `json:"content"`
	Component *components.Component  `json:"component"`
}

func NewExample() Example {

	var self Example

	element   := dom.Document.CreateElement("app-example")
	component := components.NewComponent(element)

	self.Component = &component
	self.Layout    = types.LayoutFlow

	return self

}

func ToExample(element *dom.Element) *Example {

	var self Example

	component := components.NewComponent(element)

	self.Component = &component
	self.Layout    = types.LayoutFlow

	return &self

}

func (self *Example) Disable() bool {
	return false
}

func (self *Example) Enable() bool {
	return false
}

func (self *Example) Mount() bool {

	if self.Component.Element != nil {

		name := self.Component.Element.GetAttribute("data-name")

		if name != "" {
			self.Name = name
		}

		layout := self.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			self.Layout = types.Layout(layout)
		}

		elements := self.Component.Element.Children()
		mapped   := make([]interfaces.Component, 0)

		for _, element := range elements {

			switch element.TagName {
			case "BUTTON":
				mapped = append(mapped, ui.ToButton(element))
			case "INPUT":

				typ := element.GetAttribute("type")

				if typ == "checkbox" {
					mapped = append(mapped, ui.ToCheckbox(element))
				} else if typ == "number" {
					mapped = append(mapped, ui.ToNumber(element))
				} else if typ == "range" {
					mapped = append(mapped, ui.ToRange(element))
				} else {
					mapped = append(mapped, ui.ToInput(element))
				}
			case "SELECT":
				mapped = append(mapped, ui.ToSelect(element))
			case "TEXTAREA":
				mapped = append(mapped, ui.ToTextarea(element))
			default:
				component := components.NewComponent(element)
				mapped = append(mapped, &component)
			}

		}

		self.Content = mapped

		return true

	}

	return false

}

func (self *Example) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if self.Component.Element != nil {

			if utils.MatchesQuery(self.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range self.Content {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if self.Component.Element != nil {

			if utils.MatchesQuery(self.Component.Element, selectors[0]) == true {
				return self
			}

		}

	}

	return nil

}

func (self *Example) Render() *dom.Element {

	if self.Component.Element != nil {

		if len(self.Content) > 0 {

			elements := make([]*dom.Element, 0)

			for _, component := range self.Content {
				elements = append(elements, component.Render())
			}

			self.Component.Element.ReplaceChildren(elements)

		}

	}

	return self.Component.Element

}

func (self *Example) String() string {

	html := "<app-example"

	if self.Name != "" {
		html += " data-name=\"" + self.Name + "\""
	}

	html += ">"

	if len(self.Content) > 0 {

		for _, component := range self.Content {
			html += component.String()
		}

	}

	html += "</app-example>"

	return html

}

func (self *Example) Unmount() bool {

	// Nothing To Do, has no Event Listeners
	return true

}
