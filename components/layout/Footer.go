package layout

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
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

	element   := bindings.Document.CreateElement("footer")
	component := components.NewComponent(element)

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Component.InitEvent("click")
	footer.Component.InitEvent("action")

	footer.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]string) {

		_, ok1 := attributes["data-action"]

		if ok1 == true {

			footer.Component.FireEventListeners("action", map[string]string{
				"action": attributes["data-action"],
			})

		}

	}, false))

	footer.Render()

	return footer

}

func ToFooter(element *dom.Element) Footer {

	var footer Footer

	component := components.NewComponent(element)

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Parse()

	footer.Component.InitEvent("click")
	footer.Component.InitEvent("action")

	footer.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]string) {

		_, ok1 := attributes["data-action"]

		if ok1 == true {

			footer.Component.FireEventListeners("action", map[string]string{
				"action": attributes["data-action"],
			})

		}

	}, false))

	return footer

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

func (footer *Footer) Parse() {

	if footer.Component.Element != nil {

		layout := footer.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			footer.Layout = types.Layout(layout)
		}

		tmp := footer.Component.Element.QuerySelectorAll("div")

		if len(tmp) == 3 && tmp[0].TagName == "DIV" && tmp[1].TagName == "DIV" && tmp[2].TagName == "DIV" {

			buttons_left := tmp[0].QuerySelectorAll("button")

			for _, button := range buttons_left {
				component := ui.ToButton(button)
				footer.Content.Left = append(footer.Content.Left, &component)
			}

			elements_center := tmp[1].QuerySelectorAll("button, label, input")

			for _, element := range elements_center {

				if element.TagName == "BUTTON" {

					component := ui.ToButton(element)
					footer.Content.Center = append(footer.Content.Center, &component)

				} else if element.TagName == "LABEL" {

					component := ui.ToLabel(element)
					footer.Content.Center = append(footer.Content.Center, &component)

				} else if element.TagName == "INPUT" {

					component := ui.ToInput(element)
					footer.Content.Center = append(footer.Content.Center, &component)

				}

			}

			buttons_right := tmp[2].QuerySelectorAll("button")

			for _, button := range buttons_right {
				component := ui.ToButton(button)
				footer.Content.Right = append(footer.Content.Right, &component)
			}

		} else {

			console.Group("Footer: Invalid Markup")
			console.Error("Expected <div></div><div></div><div></div>")
			console.Error(footer.Component.Element.InnerHTML)
			console.GroupEnd("Footer: Invalid Markup")

		}

	}

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
