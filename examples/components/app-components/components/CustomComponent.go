package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type CustomComponent struct {
	Name      string                 `json:"name"`
	Layout    types.Layout           `json:"layout"`
	Content   []interfaces.Component `json:"content"`
	Component *components.Component  `json:"component"`
}

func NewCustomComponent() CustomComponent {

	var self CustomComponent

	element   := dom.Document.CreateElement("app-custom-component")
	component := components.NewComponent(element)

	self.Component = &component
	self.Layout    = types.LayoutFlow

	self.Mount()
	self.Render()

	return self

}

func ToCustomComponent(element *dom.Element) *CustomComponent {

	var self CustomComponent

	component := components.NewComponent(element)

	self.Component = &component
	self.Layout    = types.LayoutFlow

	self.Mount()

	return &self

}

func (component *CustomComponent) Disable() bool {
	return false
}

func (component *CustomComponent) Enable() bool {
	return false
}

func (component *CustomComponent) Mount() bool {

	if component.Component.Element != nil {

		name := component.Component.Element.GetAttribute("data-name")

		if name != "" {
			component.Name = name
		}

		layout := component.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			component.Layout = types.Layout(layout)
		}

		elements := component.Component.Element.Children()
		content  := make([]interfaces.Component, 0)

		for _, element := range elements {

			switch element.TagName {
			case "BUTTON":
				content = append(content, ui.ToButton(element))
			case "INPUT":

				typ := element.GetAttribute("type")

				if typ == "checkbox" {
					content = append(content, ui.ToCheckbox(element))
				} else if typ == "number" {
					content = append(content, ui.ToNumber(element))
				} else if typ == "range" {
					content = append(content, ui.ToRange(element))
				} else {
					content = append(content, ui.ToInput(element))
				}
			case "SELECT":
				content = append(content, ui.ToSelect(element))
			case "TEXTAREA":
				content = append(content, ui.ToTextarea(element))
			default:
				component := components.NewComponent(element)
				content = append(content, &component)
			}

		}

		component.Content = content

		for _, content := range component.Content {
			content.Mount()
		}

		return true

	}

	return false

}

func (component *CustomComponent) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if component.Component.Element != nil {

			if utils.MatchesQuery(component.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range component.Content {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if component.Component.Element != nil {

			if utils.MatchesQuery(component.Component.Element, selectors[0]) == true {
				return component
			}

		}

	}

	return nil

}

func (component *CustomComponent) Render() *dom.Element {

	if component.Component.Element != nil {

		if len(component.Content) > 0 {

			elements := make([]*dom.Element, 0)

			for _, component := range component.Content {
				elements = append(elements, component.Render())
			}

			component.Component.Element.ReplaceChildren(elements)

		}

	}

	return component.Component.Element

}

func (component *CustomComponent) String() string {

	html := "<app-custom-component"

	if component.Name != "" {
		html += " data-name=\"" + component.Name + "\""
	}

	if component.Layout != types.LayoutFlow {
		html += " data-layout=\"" + component.Layout.String() + "\""
	}

	html += ">"

	if len(component.Content) > 0 {

		for _, content := range component.Content {
			html += content.String()
		}

	}

	html += "</app-custom-component>"

	return html

}

func (component *CustomComponent) Unmount() bool {

	for _, content := range component.Content {
		content.Unmount()
	}

	// Nothing To Do, has no Event Listeners

	return true

}
