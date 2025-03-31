package layout

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"

type Dialog struct {
	Layout  types.Layout           `json:"layout"`
	Title   string                 `json:"title"`
	Content []interfaces.Component `json:"content"`
	// TODO: Make Footer a *layout.Footer
	Footer  struct {
		Left  []interfaces.Component `json:"left"`
		Right []interfaces.Component `json:"right"`
	} `json:"footer"`
	Component *components.Component `json:"component"`
}

func NewDialog() Dialog {

	var dialog Dialog

	element   := bindings.Document.CreateElement("dialog")
	component := components.NewComponent(element)

	element.SetAttribute("data-layout", types.LayoutFlow.String())

	dialog.Component    = &component
	dialog.Layout       = types.LayoutFlow
	dialog.Title        = "Dialog"
	dialog.Content      = make([]interfaces.Component, 0)
	dialog.Footer.Left  = make([]interfaces.Component, 0)
	dialog.Footer.Right = make([]interfaces.Component, 0)

	dialog.Component.InitEvent("click")
	dialog.Component.InitEvent("action")

	dialog.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

		action, ok1 := attributes["data-action"]

		if ok1 == true {

			if action == "close" {

				dialog.Hide()

			} else {

				dialog.Component.FireEventListeners("action", map[string]string{
					"action": attributes["data-action"],
				})

			}

		}

	}, false))

	dialog.Render()

	return dialog

}

func ToDialog(element *dom.Element) Dialog {

	var dialog Dialog

	component := components.NewComponent(element)

	dialog.Component = &component
	dialog.Layout    = types.Layout(element.GetAttribute("data-layout"))
	dialog.Content      = make([]interfaces.Component, 0)
	dialog.Footer.Left  = make([]interfaces.Component, 0)
	dialog.Footer.Right = make([]interfaces.Component, 0)

	dialog.Parse()

	dialog.Component.InitEvent("click")
	dialog.Component.InitEvent("action")

	dialog.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

		action, ok1 := attributes["data-action"]

		if ok1 == true {

			if action == "close" {

				dialog.Hide()

			} else {

				dialog.Component.FireEventListeners("action", map[string]string{
					"action": attributes["data-action"],
				})

			}

		}

	}, false))

	return dialog

}

func (dialog *Dialog) Disable() bool {

	var result bool

	if len(dialog.Footer.Left) > 0 || len(dialog.Footer.Right) > 0 {

		for _, component := range dialog.Footer.Left {
			component.Disable()
		}

		for _, component := range dialog.Footer.Right {
			component.Disable()
		}

		result = true

	}

	return result

}

func (dialog *Dialog) Enable() bool {

	var result bool

	if len(dialog.Footer.Left) > 0 || len(dialog.Footer.Right) > 0 {

		for _, component := range dialog.Footer.Left {
			component.Enable()
		}

		for _, component := range dialog.Footer.Right {
			component.Enable()
		}

		result = true

	}

	return result

}

func (dialog *Dialog) Hide() bool {

	var result bool

	if dialog.Component.Element != nil {

		dialog.Component.Element.RemoveAttribute("open")
		result = true

	}

	return result

}

func (dialog *Dialog) Parse() {

	if dialog.Component.Element != nil {

		article := dialog.Component.Element.QuerySelector("article")

		if article != nil {

			h3 := dialog.Component.Element.QuerySelector("h3")

			if h3 != nil {
				dialog.Title = h3.TextContent
			}

			// TODO: Parse article into Content
			// TODO: Parse footer into actions []string?

		} else {

			console.Group("Dialog: Invalid Markup")
			console.Error("Expected <article><h3></h3><footer></footer></article>")
			console.Error(dialog.Component.Element.InnerHTML)
			console.GroupEnd("Dialog: Invalid Markup")

		}

	}

}

func (dialog *Dialog) Render() *dom.Element {

	if dialog.Component.Element != nil {

		article := dialog.Component.Element.QuerySelector("article")

		if article == nil {
			dialog.Component.Element.SetInnerHTML("<article><button data-action=\"close\"></button><h3>Dialog</h3><footer></footer></article>")
			article = dialog.Component.Element.QuerySelector("article")
		}

		if article != nil {

			tmp := article.QuerySelectorAll("h3, footer")

			if len(tmp) == 2 && tmp[0].TagName == "H3" && tmp[1].TagName == "FOOTER" {

				// TODO: Render title into h3
				// TODO: Render Content into article
				// TODO: Render actions into footer

			}

		}

	}

	return dialog.Component.Element

}

func (dialog *Dialog) SetTitle(value string) {
	dialog.Title = strings.TrimSpace(value)
}

func (dialog *Dialog) Show() bool {

	var result bool

	if dialog.Component.Element != nil {

		dialog.Component.Element.SetAttribute("open", "")
		result = true

	}

	return result

}
