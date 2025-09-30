//go:build wasm

package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "strings"

type Document struct {
	Content  []interfaces.Component `json:"components"`
	Registry map[string]constructor `json:"registry"`
	Body     *dom.Element           `json:"body"`
}

func NewDocument() *Document {

	var document Document

	document.Content  = make([]interfaces.Component, 0)
	document.Registry = make(map[string]constructor)
	document.Body     = dom.Document.QuerySelector("body")

	return &document

}

func ToDocument(element *dom.Element) *Document {

	var document Document

	document.Content  = make([]interfaces.Component, 0)
	document.Registry = make(map[string]constructor)
	document.Body     = element

	return &document

}

func (document *Document) Register(tagname string, wrapper constructor) {
	document.Registry[strings.ToLower(tagname)] = wrapper
}

func (document *Document) Mount() bool {

	if document.Body != nil {

		if document.Body.TagName != "BODY" {

			tmp := document.Body.QuerySelector("body")

			if tmp != nil {
				document.Body = tmp
			}

		}

		content  := make([]interfaces.Component, 0)
		children := document.Body.Children()

		for _, element := range children {

			if element.TagName != "SCRIPT" && element.TagName != "STYLE" {

				wrapper, ok := document.Registry[strings.ToLower(element.TagName)]

				if ok == true {

					component := wrapper(element)

					if component != nil {
						component.Mount()
						content = append(content, component)
					}

				} else {

					component := NewComponent(element)
					nested_content  := make([]interfaces.Component, 0)
					nested_children := element.Children()

					for _, nested_element := range nested_children {

						wrapper, ok := document.Registry[strings.ToLower(nested_element.TagName)]

						if ok == true {

							nested_component := wrapper(nested_element)

							if nested_component != nil {
								nested_component.Mount()
								nested_content = append(nested_content, nested_component)
							}

						} else {

							nested_component := NewComponent(nested_element)
							nested_component.Mount()
							nested_content = append(nested_content, &nested_component)

						}

					}

					component.SetContent(nested_content)
					component.Mount()

					content = append(content, &component)

				}

			}

		}

		document.Content = content

		return true

	} else {
		return false
	}

}

func (document *Document) Query(query string) interfaces.Component {

	if document.Body != nil {

		selectors := utils.SplitQuery(query)

		if len(selectors) >= 1 {

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
	return false
}

