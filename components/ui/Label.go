package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "strings"

type Label struct {
	Label     string                `json:"label"`
	Type      string                `json:"type"`
	Component *components.Component `json:"component"`
}

func NewLabel(lbl string, typ string) Label {

	var label Label

	element   := dom.Document.CreateElement("label")
	component := components.NewComponent(element)

	label.Component = &component
	label.Label     = lbl
	label.Type      = strings.ToLower(typ)

	label.Render()

	return label

}

func ToLabel(element *dom.Element) *Label {

	var label Label

	component := components.NewComponent(element)

	label.Component = &component
	label.Label     = strings.TrimSpace(element.TextContent)
	label.Type      = element.GetAttribute("data-type")

	return &label

}

func (label *Label) Disable() bool {
	return false
}

func (label *Label) Enable() bool {
	return false
}

func (label *Label) Query(query string) interfaces.Component {
	return nil
}

func (label *Label) Render() *dom.Element {

	if label.Component.Element != nil {

		if label.Label != "" {
			label.Component.Element.SetInnerHTML(label.Label)
		}

		if label.Type != "" {
			label.Component.Element.SetAttribute("data-type", label.Type)
		}

	}

	return label.Component.Element

}

func (label *Label) SetChildren(children []interfaces.Component) bool {
	return false
}

func (label *Label) String() string {

	html := ""

	if label.Type != "" {
		html += "<label data-type=\"" + label.Type + "\">"
	} else {
		html += "<label>"
	}

	if label.Label != "" {
		html += label.Label
	}

	html += "</label>"

	return html

}
