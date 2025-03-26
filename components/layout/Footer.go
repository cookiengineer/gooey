package layout

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

// TODO: NewFooter()

func ToFooter(element *dom.Element) Footer {

	var footer Footer

	component := components.NewComponent(element)

	footer.Component      = &component
	footer.Layout         = types.LayoutFlex
	footer.Content.Left   = make([]interfaces.Component, 0)
	footer.Content.Center = make([]interfaces.Component, 0)
	footer.Content.Right  = make([]interfaces.Component, 0)

	tmp := element.QuerySelectorAll("div")

	if len(tmp) == 3 {
		// left, center and right
	} else if len(tmp) == 2 {
		// center is empty
	} else if len(tmp) == 1 {
		// center only
	}

	// TODO: Parse left, center and right <div> elements
	// TODO: Parse children into components

	footer.Component.InitEvent("click")
	footer.Component.InitEvent("action")



	return footer

}

func (footer *Footer) SetLeft(components []interfaces.Component) {
	footer.Content.Left = components
}

func (footer *Footer) SetRight(components []interfaces.Component) {
	footer.Content.Right = components
}

func (footer *Footer) Render() {

	// TODO: Render into dom Element

}

func (footer *Footer) String() string {

	// TODO: Render as string

	return ""

}
