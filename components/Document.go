package components

import "github.com/cookiengineer/gooey/bindings/dom"
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

	document.Mount()

	return &document

}

func ToDocument(element *dom.Element) *Document {

	var document Document

	document.Content  = make([]interfaces.Component, 0)
	document.Registry = make(map[string]constructor)
	document.Body     = element

	document.Mount()

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

				for _, element := range nested_children {

					wrapper, ok := document.Registry[strings.ToLower(element.TagName)]

					if ok == true {

						nested_component := wrapper(element)

						if nested_component != nil {
							nested_component.Mount()
							content = append(content, nested_component)
						}

					} else {

						nested_component := NewComponent(element)
						nested_component.Mount()
						nested_content = append(nested_content, &nested_component)

					}

				}

				component.SetContent(nested_content)
				component.Mount()

				content = append(content, &component)

			}

		}

		document.Content = content

		return true

	} else {
		return false
	}

}

func (document *Document) Query(query string) interfaces.Component {

	// TODO: Components interface needs a QueryComponent() method,
	//       where query selector is split up and delegated, so that
	//       each Component decides how to query e.g. "Footer" or "Header" or "Content[0]"

	// TODO: Support "main > section > table" syntax
	// TODO: Support "main > section > table:nth-of-type(1)" syntax
	// TODO: Support "main > section > table:nth-child(1)" syntax
	// TODO: Support "main > section > table[data-name=\"foobar\"]" syntax


	// TODO: If Query() returns nil, throw an error?


	return nil

}

func (document *Document) Unmount() bool {
	return false
}
