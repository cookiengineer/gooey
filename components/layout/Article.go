package layout

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import content_components "github.com/cookiengineer/gooey/components/content"
import ui_components "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/types"

type Article struct {
	Layout    types.Layout           `json:"layout"`
	Content   []interfaces.Component `json:"content"`
	Component *components.Component  `json:"component"`
}

func NewArticle() Article {

	var article Article

	element := dom.Document.CreateElement("article")
	component := components.NewComponent(element)

	article.Layout = types.LayoutFlow
	article.Component = &component
	article.Content = make([]interfaces.Component, 0)

	return article

}

func ToArticle(element *dom.Element) *Article {

	var article Article

	component := components.NewComponent(element)

	article.Layout = types.LayoutFlow
	article.Component = &component
	article.Content = make([]interfaces.Component, 0)

	return &article

}

func (article *Article) Disable() bool {
	return false
}

func (article *Article) Enable() bool {
	return false
}

func (article *Article) Mount() bool {

	if article.Component.Element != nil {

		layout := article.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			article.Layout = types.Layout(layout)
		}

		elements := article.Component.Element.Children()
		content := make([]interfaces.Component, 0)

		for _, element := range elements {

			switch element.TagName {
			case "BUTTON":
				content = append(content, ui_components.ToButton(element))
			case "LABEL":
				content = append(content, ui_components.ToLabel(element))
			case "INPUT":

				typ := element.GetAttribute("type")

				if typ == "checkbox" {
					content = append(content, ui_components.ToCheckbox(element))
				} else if typ == "number" {
					content = append(content, ui_components.ToNumber(element))
				} else if typ == "range" {
					content = append(content, ui_components.ToRange(element))
				} else {
					content = append(content, ui_components.ToInput(element))
				}

			case "FIELDSET":
				content = append(content, content_components.ToFieldset(element))
			case "SELECT":
				content = append(content, ui_components.ToSelect(element))
			case "TABLE":
				content = append(content, content_components.ToTable(element))
			case "TEXTAREA":
				content = append(content, ui_components.ToTextarea(element))
			default:
				component := components.NewComponent(element)
				content = append(content, &component)
			}

		}

		article.Content = content

		for _, component := range article.Content {
			component.Mount()
		}

		return true

	} else {
		return false
	}

}

func (article *Article) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if article.Component.Element != nil {

			if utils.MatchesQuery(article.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range article.Content {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if article.Component.Element != nil {

			if utils.MatchesQuery(article.Component.Element, selectors[0]) == true {
				return article
			}

		}

	}

	return nil

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

func (article *Article) SetContent(components []interfaces.Component) {

	article.Content = components
	article.Render()

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

func (article *Article) Unmount() bool {

	if len(article.Content) > 0 {

		for _, component := range article.Content {
			component.Unmount()
		}

	}

	return true
}
