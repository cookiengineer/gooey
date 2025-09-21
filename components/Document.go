package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/interfaces"

type Document struct {
	Components []interfaces.Component `json:"components"`
	Registry   map[string]constructor `json:"registry"`
}

func NewDocument() *Document {

	var document Document

	document.Components = make([]interfaces.Component, 0)
	document.Registry   = make(map[string]constructor)

	return &document

}

func (document *Document) Register(tagname string, wrapper constructor) {
	document.Registry[tagname] = wrapper
}

func (document *Document) Parse(html string) {

	// TODO: Parse HTML content

}

func (document *Document) QueryElement(query string) *dom.Element {

	// TODO

	return nil

}

func (document *Document) QueryComponent(query string) interfaces.Component {

	// TODO: Components interface needs a QueryComponent() method,
	//       where query selector is split up and delegated, so that
	//       each Component decides how to query e.g. "Footer" or "Header" or "Content[0]"

	// TODO: Support "main > section > table" syntax
	// TODO: Support "main > section > table:nth-of-type(1)" syntax
	// TODO: Support "main > section > table:nth-child(1)" syntax
	// TODO: Support "main > section > table[data-name=\"foobar\"]" syntax



	return nil

}
