package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/interfaces"
import "strings"
import "fmt"

type Document struct {
	Components map[string]interfaces.Component `json:"components"`
	Registry   map[string]constructor          `json:"registry"`
	Body       *dom.Element                    `json:"body"`
}

func NewDocument() *Document {

	var document Document

	document.Components = make(map[string]interfaces.Component, 0)
	document.Registry   = make(map[string]constructor)
	document.Body       = dom.Document.QuerySelector("body")

	document.Mount()

	return &document

}

func ToDocument(element *dom.Element) *Document {

	var document Document

	document.Components = make(map[string]interfaces.Component, 0)
	document.Registry   = make(map[string]constructor)
	document.Body       = element

	document.Mount()

	return &document

}

func (document *Document) Register(tagname string, wrapper constructor) {
	document.Registry[tagname] = wrapper
}

func (document *Document) Mount() bool {

	if document.Body != nil {

		if document.Body.TagName != "BODY" {

			tmp := document.Body.QuerySelector("body")

			if tmp != nil {
				document.Body = tmp
			}

		}

		children := document.Body.Children()

		for _, element := range children {

			// TODO: How to define identifier for query path?
			// TODO: if children contains another with same tagname, use nth-of-type(...)?

			tagname := strings.ToLower(element.TagName)

			wrapper, ok := document.Registry[tagname]

			if ok == true {

				component := wrapper(element)

				fmt.Println("component", component)

			} else {

				fmt.Println("element", element)

				// TODO: Make it a default component

				// TODO: Find out children, iterate and do the same

			}

		}

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
