package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Example struct {
	Layout    types.Layout           `json:"layout"`
	Content   []interfaces.Component `json:"content"`
	Component *components.Component  `json:"component"`
}

func NewExample() Example {

	var self Example

	element   := dom.Document.CreateElement("app-example")
	component := components.NewComponent(element)

	self.Component = &component
	article.Layout = types.LayoutFlow

	return self

}

func ToExample(element *dom.Element) *Example {

	var self Example

	component := components.NewComponent(element)

	self.Component = &component
	article.Layout = types.LayoutFlow

	return &self

}

func (self *Example) Disable() bool {
	return false
}

func (self *Example) Enable() bool {
	return false
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
