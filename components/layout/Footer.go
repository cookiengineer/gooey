package layout

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Footer struct {
	Layout  types.Layout `json:"layout"`
	Content struct {
		Left   []interfaces.Component `json:"left"`
		Center []interfaces.Component `json:"center"`
		Right  []interfaces.Component `json:"right"`
	} `json:"layout"`
	Component *components.Component `json:"component"`
}

func NewFooter() Footer {

	var footer Footer

	element   := bindings.Document.CreateElement("footer")
	component := components.NewComponent(element)

	element.SetAttribute("data-layout", types.LayoutFlex.String())

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Component.InitEvent("click")
	footer.Component.InitEvent("action")

	footer.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

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
	footer.Layout         = types.Layout(element.GetAttribute("data-layout"))
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	footer.Parse()

	footer.Component.InitEvent("click")
	footer.Component.InitEvent("action")

	footer.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

		_, ok1 := attributes["data-action"]

		if ok1 == true {

			footer.Component.FireEventListeners("action", map[string]string{
				"action": attributes["data-action"],
			})

		}

	}, false))

	return footer

}

func (footer *Footer) Parse() {

	// TODO

}

func (footer *Footer) Render() *dom.Element {

	if footer.Component.Element != nil {

		// TODO: Render first <div></div> for footer.Content.Left
		// TODO: Render second <div></div> for footer.Content.Center
		// TODO: Render third <div></div> for footer.Content.Right

		console.Warn(footer.Component.Element)

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
