package layout

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"

type Article struct {
	Layout    types.Layout           `json:"layout"`
	Content   []interfaces.Component `json:"content"`
	Component *components.Component  `json:"component"`
}

func NewArticle() Article {

	var article Article

	element   := dom.Document.CreateElement("article")
	component := components.NewComponent(element)

	article.Layout    = types.LayoutFlow
	article.Component = &component
	article.Content   = make([]interfaces.Component, 0)

	article.Render()

	return article

}

func ToArticle(element *dom.Element) *Article {

	var article Article

	component := components.NewComponent(element)

	article.Layout    = types.LayoutFlow
	article.Component = &component
	article.Content   = make([]interfaces.Component, 0)

	article.Parse()

	return &article

}

func (article *Article) Disable() bool {
	return false
}

func (article *Article) Enable() bool {
	return false
}

func (article *Article) Parse() {

	if article.Component.Element != nil {

		layout := article.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			article.Layout = types.Layout(layout)
		}

		elements := article.Component.Element.Children()
		mapped   := make([]interfaces.Component, 0)

		for _, element := range elements {

			switch element.TagName {
			case "BUTTON":
				mapped = append(mapped, ui.ToButton(element))
			case "LABEL":
				mapped = append(mapped, ui.ToLabel(element))
			case "INPUT":

				typ := element.GetAttribute("type")

				if typ == "checkbox" {
					mapped = append(mapped, ui.ToCheckbox(element))
				} else if typ == "radio" {

					// TODO: Radio support

				} else {
					mapped = append(mapped, ui.ToInput(element))
				}

			case "FIELDSET":
				mapped = append(mapped, content.ToFieldset(element))
			case "SELECT":
				mapped = append(mapped, ui.ToSelect(element))
			case "TABLE":
				mapped = append(mapped, content.ToTable(element))
			case "TEXTAREA":
				mapped = append(mapped, ui.ToTextarea(element))
			default:
				component := components.NewComponent(element)
				mapped = append(mapped, &component)
			}

		}

		article.Content = mapped

	}

}

func (article *Article) Render() *dom.Element {

	if article.Component.Element != nil {

		if article.Layout != types.LayoutFlow {
			article.Component.Element.SetAttribute("data-layout", article.Layout.String())
		}

		if len(article.Content) > 0 {

			elements := make([]*dom.Element, 0)

			for _, component := range article.Content {
				elements = append(elements, component.Render())
			}

			article.Component.Element.ReplaceChildren(elements)

		}

	}

	return article.Component.Element

}

func (article *Article) String() string {

	html := "<article"

	if article.Layout != types.LayoutFlow {
		html += " data-layout=\"" + article.Layout.String() + "\""
	}

	html += ">"

	if len(article.Content) > 0 {

		for _, component := range article.Content {
			html += component.String()
		}

	}

	html += "</article>"

	return html

}
