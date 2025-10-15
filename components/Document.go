//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "strings"

type Document struct {
	Content  []interfaces.Component          `json:"components"`
	Registry map[string]ComponentConstructor `json:"registry"`
	body     *dom.Element                    `json:"-"`
}

func NewDocument() *Document {

	var document Document

	document.Content = make([]interfaces.Component, 0)
	document.Registry = make(map[string]ComponentConstructor)
	document.body = dom.Document.QuerySelector("body")

	return &document

}

func ToDocument(element *dom.Element) *Document {

	var document Document

	document.Content = make([]interfaces.Component, 0)
	document.Registry = make(map[string]ComponentConstructor)
	document.body = element

	return &document

}

func (document *Document) Register(tagname string, wrapper ComponentConstructor) {
	document.Registry[strings.ToLower(tagname)] = wrapper
}

func (document *Document) Mount() bool {

	if document.body != nil {

		if document.body.TagName != "BODY" {

			tmp := document.body.QuerySelector("body")

			if tmp != nil {
				document.body = tmp
			}

		}

		content := make([]interfaces.Component, 0)
		children := document.body.Children()

		for _, element := range children {

			if element.TagName != "SCRIPT" && element.TagName != "STYLE" {

				wrapper, ok := document.Registry[strings.ToLower(element.TagName)]

				if ok == true {

					component := wrapper(element)

					if component != nil {
						content = append(content, component)
					}

				} else {

					component := NewComponent(element)
					traverseComponent(document, &component)
					content = append(content, &component)

				}

			}

		}

		document.Content = content

		for _, component := range document.Content {
			component.Mount()
		}

		return true

	} else {
		return false
	}

}

func (document *Document) QueryComponent(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 && selectors[0] == "body" {

		if len(document.Content) > 0 {

			tmp_query := utils.JoinQuery(selectors[1:])

			for _, component := range document.Content {

				tmp_component := component.Query(tmp_query)

				if tmp_component != nil {
					return tmp_component
				}

			}

		}

	} else if len(selectors) >= 1 {

		if len(document.Content) > 0 {

			tmp_query := utils.JoinQuery(selectors)

			for _, content := range document.Content {

				tmp_component := content.Query(tmp_query)

				if tmp_component != nil {
					return tmp_component
				}

			}

		}

	}

	return nil

}

func (document *Document) QuerySelector(query string) *dom.Element {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 && selectors[0] == "body" {

		if document.body != nil {

			children := document.body.Children()

			for _, element := range children {

				if utils.MatchesQuery(element, selectors[2]) == true {

					tmp_query := utils.JoinQuery(selectors[2:])
					tmp_found := element.QuerySelector(tmp_query)

					if tmp_found != nil {
						return tmp_found
					}

				}

			}

		}

	} else if len(selectors) >= 1 {

		if document.body != nil {
			return document.body.QuerySelector(utils.JoinQuery(selectors))
		}

	}

	return nil

}

func (document *Document) QuerySelectorAll(query string) []*dom.Element {

	result := make([]*dom.Element, 0)
	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 && selectors[0] == "body" {

		if document.body != nil {

			children := document.body.Children()

			for _, element := range children {

				if utils.MatchesQuery(element, selectors[2]) == true {

					tmp_query := utils.JoinQuery(selectors[2:])
					tmp_found := element.QuerySelectorAll(tmp_query)

					for _, tmp_element := range tmp_found {
						result = append(result, tmp_element)
					}

				}

			}

		}

	} else if len(selectors) >= 1 {
		result = document.body.QuerySelectorAll(utils.JoinQuery(selectors))
	}

	return result

}

func (document *Document) String() string {

	html := "<!DOCTYPE html>"
	html += "<html>"
	html += "<body>"

	for _, content := range document.Content {
		html += content.String()
	}

	html += "</body>"
	html += "</html>"

	return html

}

func (document *Document) Unmount() bool {

	for _, component := range document.Content {
		component.Unmount()
	}

	return true

}
