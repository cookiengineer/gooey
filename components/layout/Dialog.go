package layout

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"

type Dialog struct {
	Layout    types.Layout          `json:"layout"`
	Title     string                `json:"title"`
	Content   interfaces.Component  `json:"content"`
	Footer    *Footer               `json:"footer"`
	Component *components.Component `json:"component"`
}

func NewDialog() Dialog {

	var dialog Dialog

	element   := bindings.Document.CreateElement("dialog")
	component := components.NewComponent(element)

	element.SetAttribute("data-layout", types.LayoutFlow.String())

	dialog.Component = &component
	dialog.Layout    = types.LayoutFlow
	dialog.Title     = "Dialog"
	dialog.Content   = nil
	dialog.Footer    = nil

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
	dialog.Content   = nil
	dialog.Footer    = nil

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

	dialog.Footer.Component.AddEventListener("action", components.ToComponentListener(func(event string, attributes map[string]string) {

		dialog.Component.FireEventListeners("action", map[string]string{
			"action": attributes["data-action"],
		})

	}, false))

	return dialog

}

func (dialog *Dialog) Disable() bool {

	var result bool

	if dialog.Footer != nil {

		dialog.Footer.Disable()
		result = true

	}

	return result

}

func (dialog *Dialog) Enable() bool {

	var result bool

	if dialog.Footer != nil {

		dialog.Footer.Enable()
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

			tmp1 := article.QuerySelector("h3")

			if tmp1 != nil {
				dialog.Title = tmp1.TextContent
			}

			tmp2 := article.QuerySelector("fieldset, table")

			if tmp2 != nil {

				if tmp2.TagName == "FIELDSET" {

					component := content.ToFieldset(tmp2)
					dialog.Content = &component

				} else if tmp2.TagName == "TABLE" {

					component := content.ToTable(tmp2)
					dialog.Content = &component

				}

			}

			tmp3 := article.QuerySelector("footer")

			if tmp3 != nil {

				component := ToFooter(tmp3)
				dialog.Footer = &component

			}

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

			tmp := article.QuerySelectorAll("button[data-action=\"close\"], h3, footer")

			if len(tmp) == 3 && tmp[0].TagName == "BUTTON" && tmp[1].TagName == "H3" && tmp[2].TagName == "FOOTER" {

				elements := make([]*dom.Element, 0)

				tmp[1].SetInnerHTML(dialog.Title)

				elements = append(elements, tmp[0])
				elements = append(elements, tmp[1])

				if dialog.Content != nil {
					elements = append(elements, dialog.Content.Render())
				}

				elements = append(elements, tmp[2])

				article.ReplaceChildren(elements)

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
