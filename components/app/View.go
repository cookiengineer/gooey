//go:build wasm

package app

import "fmt"

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
// import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"

type View struct {
	Element *dom.Element           `json:"element"`
	Layout  types.Layout           `json:"layout"`
	Name    string                 `json:"name"`
	Label   string                 `json:"label"`
	Path    string                 `json:"path"`
	Content []interfaces.Component `json:"components"`
}

func NewView(name string, label string, path string) View {

	var view View

	element := bindings.Document.CreateElement("section")

	view.Element = element
	view.Name    = strings.ToLower(name)
	view.Label   = label
	view.Layout  = types.LayoutFlow
	view.Path    = strings.ToLower(path)
	view.Content = make([]interfaces.Component, 0)

	return view

}

func ToView(element *dom.Element, label string, path string) View {

	var view View

	view.Element = element
	view.Layout  = types.LayoutFlow
	view.Label   = label
	view.Path    = strings.ToLower(path)
	view.Content = make([]interfaces.Component, 0)

	view.Parse()

	return view

}

func (view *View) Enter() bool {

	if view.Element != nil {
		view.Element.SetAttribute("data-state", "active")
	}

	return true

}

func (view *View) Leave() bool {

	if view.Element != nil {
		view.Element.RemoveAttribute("data-state")
	}

	return true

}

func (view *View) GetProperty(name string) string {

	var result string

	switch name {
	case "Name":
		result = view.Name
	case "Label":
		result = view.Label
	case "Path":
		result = view.Path
	}

	return result

}

func (view *View) Parse() {

	if view.Element != nil {

		name := view.Element.GetAttribute("data-name")

		if name != "" {
			view.Name = strings.ToLower(name)
		}

		layout := view.Element.GetAttribute("data-layout")

		if layout != "" {
			view.Layout = types.Layout(layout)
		}

		elements   := view.Element.Children()
		components := make([]interfaces.Component, 0)

		for _, element := range elements {

			if element.TagName == "ARTICLE" {

				// component := content.ToArticle(element)
				// components = append(components, &component)

			} else if element.TagName == "TABLE" {

				// component := content.ToTable(element)
				// components = append(components, &component)

			}

			fmt.Println(element.TagName)

		}

		view.Content = components

	}

}

func (view *View) Render() {

	if view.Element != nil {

		if view.Name != "" {
			view.Element.SetAttribute("data-name", strings.ToLower(view.Name))
		}

		if view.Layout != types.LayoutFlow {
			view.Element.SetAttribute("data-layout", view.Layout.String())
		}

		elements := make([]*dom.Element, 0)

		for _, component := range view.Content {
			elements = append(elements, component.Render())
		}

		view.Element.ReplaceChildren(elements)

	}

}

func (view *View) SetProperty(name string, value string) bool {

	var result bool

	switch name {
	case "Name":
		view.Name = value
		result    = true
	case "Label":
		view.Label = value
		result     = true
	case "Path":
		view.Path = value
		result    = true
	}

	return result

}

