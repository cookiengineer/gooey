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

	return &document

}

func (document *Document) Register(tagname string, wrapper constructor) {
	document.Registry[tagname] = wrapper
}

func (document *Document) Parse(body *dom.Element) {

	if body.TagName != "BODY" {

		tmp := body.QuerySelector("body")

		if tmp != nil {
			body = tmp
		}

	}

	if body.TagName == "BODY" {

		children := body.Children()

		for _, element := range children {

			// TODO: How to define identifier for query path?
			// TODO: if children contains another with same tagname, use nth-of-type(...)?

			tagname := strings.ToLower(element.TagName)

			wrapper, ok := document.Registry[tagname]

			if ok == true {

				component := wrapper(element)

				fmt.Println(component)

			} else {

				// TODO: Make it a default component

				// TODO: Find out children, iterate and do the same

			}

		}

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



	return nil

}
