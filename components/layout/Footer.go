package components

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
	components.Component
}

func ToFooter(element *dom.Element) Footer {

	var footer Footer

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

	footer.Init(element)

	return footer

}

func (footer *Footer) SetContent(left []interfaces.Component, center []interfaces.Component, right []interfaces.Component) {

	footer.Content.Left   = left
	footer.Content.Center = center
	footer.Content.Right  = right

}

func (footer *Footer) Render() {

	// TODO: Render into dom Element

}

func (footer *Footer) String() string {

	// TODO: Render as string

	return ""

}
