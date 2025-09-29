package layout

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Footer struct {
	Layout  types.Layout `json:"layout"`
	Content struct {
		Left   []interfaces.Component `json:"left"`
		Center []interfaces.Component `json:"center"`
		Right  []interfaces.Component `json:"right"`
	} `json:"content"`
	Component *components.Component `json:"component"`
}

func NewFooter() Footer {

	var footer Footer

	element   := dom.Document.CreateElement("footer")
	component := components.NewComponent(element)

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Mount()
	footer.Render()

	return footer

}

func ToFooter(element *dom.Element) *Footer {

	var footer Footer

	component := components.NewComponent(element)

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Mount()

	return &footer

}

func (footer *Footer) Disable() bool {

	var result bool

	if len(footer.Content.Left) > 0 || len(footer.Content.Right) > 0 {

		for _, component := range footer.Content.Left {
			component.Disable()
		}

		for _, component := range footer.Content.Right {
			component.Disable()
		}

		result = true

	}

	return result

}

func (footer *Footer) Enable() bool {

	var result bool

	if len(footer.Content.Left) > 0 || len(footer.Content.Right) > 0 {

		for _, component := range footer.Content.Left {
			component.Enable()
		}

		for _, component := range footer.Content.Right {
			component.Enable()
		}

		result = true

	}

	return result

}

func (footer *Footer) Mount() bool {

	if footer.Component != nil {

		footer.Component.InitEvent("click")
		footer.Component.InitEvent("action")

		footer.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]any) {

			_, ok1 := attributes["data-action"]

			if ok1 == true {

				footer.Component.FireEventListeners("action", map[string]any{
					"action": attributes["data-action"],
				})

			}

		}, false))

	}

	if footer.Component.Element != nil {

		layout := footer.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			footer.Layout = types.Layout(layout)
		}

		tmp := footer.Component.Element.QuerySelectorAll("div")

		if len(tmp) == 3 && tmp[0].TagName == "DIV" && tmp[1].TagName == "DIV" && tmp[2].TagName == "DIV" {

			buttons_left := tmp[0].QuerySelectorAll("button")

			for _, button := range buttons_left {
				footer.Content.Left = append(footer.Content.Left, ui.ToButton(button))
			}

			elements_center := tmp[1].QuerySelectorAll("button, label, input")

			for _, element := range elements_center {

				if element.TagName == "BUTTON" {
					footer.Content.Center = append(footer.Content.Center, ui.ToButton(element))
				} else if element.TagName == "LABEL" {
					footer.Content.Center = append(footer.Content.Center, ui.ToLabel(element))
				} else if element.TagName == "INPUT" {
					footer.Content.Center = append(footer.Content.Center, ui.ToInput(element))
				}

			}

			buttons_right := tmp[2].QuerySelectorAll("button")

			for _, button := range buttons_right {
				footer.Content.Right = append(footer.Content.Right, ui.ToButton(button))
			}

			return true

		} else {

			console.Group("Footer: Invalid Markup")
			console.Error("Expected <div></div><div></div><div></div>")
			console.Error(footer.Component.Element.InnerHTML)
			console.GroupEnd("Footer: Invalid Markup")

			return false

		}

	} else {
		return false
	}

}

func (footer *Footer) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if footer.Component.Element != nil {

			if utils.MatchesQuery(footer.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range footer.Content.Left {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

				for _, content := range footer.Content.Center {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

				for _, content := range footer.Content.Right {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if footer.Component.Element != nil {

			if utils.MatchesQuery(footer.Component.Element, selectors[0]) == true {
				return footer
			}

		}

	}

	return nil

}

func (footer *Footer) Render() *dom.Element {

	if footer.Component.Element != nil {

		tmp := footer.Component.Element.QuerySelectorAll("div")

		if len(tmp) == 0 {
			footer.Component.Element.SetInnerHTML("<div></div><div></div><div></div>")
			tmp = footer.Component.Element.QuerySelectorAll("div")
		}

		if len(tmp) == 3 {

			footer.Component.Element.SetAttribute("data-layout", footer.Layout.String())

			elements_left   := make([]*dom.Element, 0)
			elements_center := make([]*dom.Element, 0)
			elements_right  := make([]*dom.Element, 0)

			for _, component := range footer.Content.Left {
				elements_left = append(elements_left, component.Render())
			}

			for _, component := range footer.Content.Center {
				elements_center = append(elements_center, component.Render())
			}

			for _, component := range footer.Content.Right {
				elements_right = append(elements_right, component.Render())
			}

			tmp[0].ReplaceChildren(elements_left)
			tmp[1].ReplaceChildren(elements_center)
			tmp[2].ReplaceChildren(elements_left)

		}

	}

	return footer.Component.Element

}

func (footer *Footer) SetCenter(components []interfaces.Component) {
	footer.Content.Center = components
}

func (footer *Footer) SetLeft(components []interfaces.Component) {
	footer.Content.Left = components
}

func (footer *Footer) SetRight(components []interfaces.Component) {
	footer.Content.Right = components
}

func (footer *Footer) String() string {

	html := "<footer"

	if footer.Layout != types.LayoutFlex {
		html += " data-layout=\"" + footer.Layout.String() + "\""
	}

	html += ">"
	html += "<div>"

	if len(footer.Content.Left) > 0 {

		for _, component := range footer.Content.Left {
			html += component.String()
		}

	}

	html += "</div>"
	html += "<div>"

	if len(footer.Content.Center) > 0 {

		for _, component := range footer.Content.Center {
			html += component.String()
		}

	}

	html += "</div>"
	html += "<div>"

	if len(footer.Content.Right) > 0 {

		for _, component := range footer.Content.Right {
			html += component.String()
		}

	}

	html += "</div>"
	html += "</footer>"

	return html

}

func (footer *Footer) Unmount() bool {

	if footer.Component != nil {
		footer.Component.RemoveEventListener("click", nil)
	}

	return true

}
