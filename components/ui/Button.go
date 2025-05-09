package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type Button struct {
	Label     string                `json:"label"`
	Action    string                `json:"action"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewButton(label string, action string) Button {

	var button Button

	element   := dom.Document.CreateElement("button")
	component := components.NewComponent(element)

	button.Component = &component
	button.Label     = label
	button.Action    = strings.ToLower(action)
	button.Disabled  = false

	button.Component.InitEvent("click")
	button.Render()

	return button

}

func ToButton(element *dom.Element) *Button {

	var button Button

	component := components.NewComponent(element)

	button.Component = &component
	button.Label     = strings.TrimSpace(element.TextContent)
	button.Action    = strings.ToLower(element.GetAttribute("data-action"))
	button.Disabled  = element.HasAttribute("disabled")

	button.Component.InitEvent("click")

	return &button

}

func (button *Button) Disable() bool {

	button.Disabled = true
	button.Render()

	return true

}

func (button *Button) Enable() bool {

	button.Disabled = false
	button.Render()

	return true

}

func (button *Button) Render() *dom.Element {

	if button.Component.Element != nil {

		if button.Label != "" {
			button.Component.Element.SetInnerHTML(button.Label)
		} else {
			button.Component.Element.SetInnerHTML("")
		}

		if button.Action != "" {
			button.Component.Element.SetAttribute("data-action", button.Action)
		} else {
			button.Component.Element.RemoveAttribute("data-action")
		}

		if button.Disabled == true {
			button.Component.Element.SetAttribute("disabled", "")
		} else {
			button.Component.Element.RemoveAttribute("disabled")
		}

	}

	return button.Component.Element

}

func (button *Button) String() string {

	html := "<button"

	if button.Action != "" {
		html += " data-action=\"" + button.Action + "\""
	}

	if button.Disabled == true {
		html += " disabled"
	}

	html += ">"

	if button.Label != "" {
		html += button.Label
	}

	html += "</button>"

	return html

}
