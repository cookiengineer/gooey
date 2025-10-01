//go:build wasm

package views

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import app_components "example/components"
import "sort"
import "strings"

type Settings struct {
	Element *dom.Element           `json:"element"`
	Layout  types.Layout           `json:"layout"`
	Content []interfaces.Component `json:"content"`
	name    string                 `json:"name"`
	label   string                 `json:"label"`
	path    string                 `json:"path"`
}

func NewSettings(name string, label string, path string) *Settings {

	var view Settings

	element := dom.Document.CreateElement("section")

	view.Element = element
	view.Layout  = types.LayoutFlow
	view.Content = make([]interfaces.Component, 0)

	view.name    = strings.ToLower(name)
	view.label   = label
	view.path    = strings.ToLower(path)

	return &view

}

func ToSettings(element *dom.Element) *Settings {

	var view Settings

	view.Element = element
	view.Layout  = types.LayoutFlow
	view.Content = make([]interfaces.Component, 0)

	view.name  = strings.ToLower(element.GetAttribute("data-name"))
	view.label = element.GetAttribute("data-label")
	view.path  = strings.ToLower(element.GetAttribute("data-path"))

	return &view

}

func (view *Settings) Disable() bool {
	return false
}

func (view *Settings) Enable() bool {
	return false
}

func (view *Settings) Enter() bool {

	if view.Element != nil {
		view.Element.SetAttribute("data-state", "active")
	}

	return true

}

func (view *Settings) Leave() bool {

	if view.Element != nil {
		view.Element.RemoveAttribute("data-state")
	}

	return true

}

func (view *Settings) Label() string {
	return view.label
}

func (view *Settings) Mount() bool {

	if view.Element != nil {

		tmp_name := view.Element.GetAttribute("data-name")

		if tmp_name != "" {
			view.name = strings.ToLower(tmp_name)
		}

		tmp_label := view.Element.GetAttribute("data-label")

		if tmp_label != "" {
			view.label = tmp_label
		}

		tmp_layout := view.Element.GetAttribute("data-layout")

		if tmp_layout != "" {
			view.Layout = types.Layout(tmp_layout)
		}

		tmp_path := view.Element.GetAttribute("data-path")

		if tmp_path != "" {
			view.path = strings.ToLower(tmp_path)
		}

		elements   := view.Element.Children()
		components := make([]interfaces.Component, 0)

		for _, element := range elements {

			if element.TagName == "ARTICLE" {
				components = append(components, layout.ToArticle(element))
			} else if element.TagName == "TABLE" {
				components = append(components, content.ToTable(element))
			} else if element.TagName == "APP-CUSTOM-COMPONENT" {
				components = append(components, app_components.ToCustomComponent(element))
			}

		}

		view.Content = components

		for _, component := range view.Content {
			component.Mount()
		}

		return true

	} else {
		return false
	}

}

func (view *Settings) Name() string {
	return view.name
}

func (view *Settings) Path() string {
	return view.path
}

func (view *Settings) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if view.Element != nil {

			if utils.MatchesQuery(view.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range view.Content {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if view.Element != nil {

			if utils.MatchesQuery(view.Element, selectors[0]) == true {
				return view
			}

		}

	}

	return nil

}

func (view *Settings) QuerySelector(query string) *dom.Element {

	if view.Element != nil {
		return view.Element.QuerySelector(query)
	}

	return nil

}

func (view *Settings) QuerySelectorAll(query string) []*dom.Element {

	result := make([]*dom.Element, 0)

	if view.Element != nil {
		result = view.Element.QuerySelectorAll(query)
	}

	return result

}


func (view *Settings) Render() *dom.Element {

	if view.Element != nil {

		if view.name != "" {
			view.Element.SetAttribute("data-name", view.name)
		}

		if view.label != "" {
			view.Element.SetAttribute("data-label", view.label)
		}

		if view.path != "" {
			view.Element.SetAttribute("data-path", view.path)
		}

		if view.Layout != types.LayoutFlow {
			view.Element.SetAttribute("data-layout", view.Layout.String())
		}

		elements := make([]*dom.Element, 0)

		for _, component := range view.Content {
			elements = append(elements, component.Render())
		}

		view.Element.ReplaceChildren(elements)

		return view.Element

	}

	return nil

}

func (view *Settings) String() string {

	html := ""

	if view.Element != nil {

		tagname := strings.ToLower(view.Element.TagName)

		html += "<" + tagname

		attributes := make([]string, 0)

		for key, _ := range view.Element.Attributes {
			attributes = append(attributes, key)
		}

		sort.Strings(attributes)

		for _, attribute := range attributes {
			html += " " + attribute + "=\"" + view.Element.Attributes[attribute] + "\""
		}

		html += ">"
		html += view.Element.InnerHTML
		html += "</" + tagname + ">"

	}

	return html

}

func (view *Settings) Unmount() bool {

	for _, component := range view.Content {
		component.Unmount()
	}

	return true

}
