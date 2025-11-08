package layout

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/types"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "strings"

type header_item struct {
	Name    string
	Label   string
	Path    string
	Element *dom.Element
}

type Header struct {
	Layout  types.Layout `json:"layout"`
	Content struct {
		Left  []interfaces.Component `json:"left"`
		Right []interfaces.Component `json:"right"`
	} `json:"content"`
	Component *components.Component `json:"component"`
	active    string                `json:"-"`
	items     []*header_item        `json:"-"`
}

func NewHeader() Header {

	var header Header

	element := dom.GetDocument().CreateElement("header")
	component := components.NewComponent(element)

	header.Component = &component
	header.Layout = types.LayoutFlex
	header.Content.Left = make([]interfaces.Component, 0)
	header.Content.Right = make([]interfaces.Component, 0)
	header.active = ""
	header.items = make([]*header_item, 0)

	return header

}

func ToHeader(element *dom.Element) *Header {

	var header Header

	component := components.NewComponent(element)

	header.Component = &component
	header.Layout = types.LayoutFlex
	header.Content.Left = make([]interfaces.Component, 0)
	header.Content.Right = make([]interfaces.Component, 0)
	header.active = ""
	header.items = make([]*header_item, 0)

	return &header

}

func (header *Header) ChangeView(name string) {

	var found *header_item = nil

	for _, item := range header.items {

		if item.Name == name {
			found = item
			break
		}

	}

	if found != nil {

		for _, item := range header.items {

			if item.Element != nil {

				if item.Name == found.Name {
					item.Element.SetAttribute("data-state", "active")
				} else {
					item.Element.RemoveAttribute("data-state")
				}

			}

		}

		header.active = name

	}

}

func (header *Header) Disable() bool {

	var result bool

	if len(header.Content.Left) > 0 || len(header.Content.Right) > 0 {

		for _, component := range header.Content.Left {
			component.Disable()
		}

		for _, component := range header.Content.Right {
			component.Disable()
		}

		result = true

	}

	return result

}

func (header *Header) Enable() bool {

	var result bool

	if len(header.Content.Left) > 0 || len(header.Content.Right) > 0 {

		for _, component := range header.Content.Left {
			component.Enable()
		}

		for _, component := range header.Content.Right {
			component.Enable()
		}

		result = true

	}

	return result

}

func (header *Header) Mount() bool {

	if header.Component != nil {

		header.Component.InitEvent("change-view")
		header.Component.InitEvent("action")

	}

	if header.Component.Element != nil {

		header.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target != nil {

				action := event.Target.GetAttribute("data-action")
				view := event.Target.GetAttribute("data-view")
				path := event.Target.GetAttribute("href")

				if action != "" {

					event.PreventDefault()
					event.StopPropagation()

					header.Component.FireEventListeners("action", map[string]any{
						"action": action,
					})

				} else if view != "" && path != "" {

					event.PreventDefault()
					event.StopPropagation()

					header.ChangeView(view)
					header.Component.FireEventListeners("change-view", map[string]any{
						"name": view,
						"path": path,
					})

				}

			}

		}))

		layout := header.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			header.Layout = types.Layout(layout)
		}

		tmp := header.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 3 && tmp[0].TagName == "DIV" && tmp[1].TagName == "UL" && tmp[2].TagName == "DIV" {

			buttons_left := tmp[0].QuerySelectorAll("button")
			content_left := make([]interfaces.Component, 0)

			for _, button := range buttons_left {
				content_left = append(content_left, ui.ToButton(button))
			}

			elements_center := tmp[1].QuerySelectorAll("li")

			for _, element := range elements_center {

				link := element.QuerySelector("a")

				if link != nil {

					item := header_item{
						Name:    link.GetAttribute("data-view"),
						Label:   strings.TrimSpace(link.TextContent),
						Path:    link.GetAttribute("href"),
						Element: element,
					}

					if item.Name != "" {

						if element.GetAttribute("data-state") == "active" {
							header.active = item.Name
							header.items = append(header.items, &item)
						} else {
							header.items = append(header.items, &item)
						}

					}

				}

			}

			buttons_right := tmp[2].QuerySelectorAll("button")
			content_right := make([]interfaces.Component, 0)

			for _, button := range buttons_right {
				content_right = append(content_right, ui.ToButton(button))
			}

			header.Content.Left = content_left
			header.Content.Right = content_right

			for _, component := range header.Content.Left {
				component.Mount()
			}

			for _, component := range header.Content.Right {
				component.Mount()
			}

			return true

		} else {

			console1 := console.GetConsole()
			console1.Group("gooey/components/layout.Header: Invalid markup")
			console1.Error("Expected <div/><ul/><div/> but got this instead:")
			console1.Error(header.Component.Element.InnerHTML)
			console1.GroupEnd()

			return false

		}

	} else {
		return false
	}

}

func (header *Header) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if header.Component.Element != nil {

			if utils.MatchesQuery(header.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range header.Content.Left {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

				for _, content := range header.Content.Right {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if header.Component.Element != nil {

			if utils.MatchesQuery(header.Component.Element, selectors[0]) == true {
				return header
			}

		}

	}

	return nil

}

func (header *Header) RegisterView(view interfaces.View) bool {

	name := view.Name()
	label := view.Label()
	path := view.Path()

	if name != "" && label != "" && path != "" {

		var found *header_item = nil

		for _, item := range header.items {

			if item.Name == name {
				found = item
				break
			}

		}

		if found != nil {

			found.Name = name
			found.Label = label
			found.Path = path

			return true

		} else {

			item := header_item{
				Name:    name,
				Label:   label,
				Path:    path,
				Element: dom.GetDocument().CreateElement("li"),
			}

			header.items = append(header.items, &item)

			return true

		}

	}

	return false

}

func (header *Header) Render() *dom.Element {

	if header.Component.Element != nil {

		tmp := header.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 0 {
			header.Component.Element.SetInnerHTML("<div></div><ul></ul><div></div>")
			tmp = header.Component.Element.QuerySelectorAll("div, ul")
		}

		if len(tmp) == 3 {

			if header.Layout != types.LayoutFlex {
				header.Component.Element.SetAttribute("data-layout", header.Layout.String())
			}

			elements_left := make([]*dom.Element, 0)
			elements_center := make([]*dom.Element, 0)
			elements_right := make([]*dom.Element, 0)

			for _, component := range header.Content.Left {
				elements_left = append(elements_left, component.Render())
			}

			for _, item := range header.items {

				if item.Name == header.active {
					item.Element.SetAttribute("data-state", "active")
				} else {
					item.Element.RemoveAttribute("data-state")
				}

				item.Element.SetInnerHTML("<a data-view=\"" + item.Name + "\" href=\"" + item.Path + "\">" + item.Label + "</a>")
				elements_center = append(elements_center, item.Element)

			}

			for _, component := range header.Content.Right {
				elements_right = append(elements_right, component.Render())
			}

			tmp[0].ReplaceChildren(elements_left)
			tmp[1].ReplaceChildren(elements_center)
			tmp[2].ReplaceChildren(elements_right)

		}

	}

	return header.Component.Element

}

func (header *Header) SetContentLeft(components []interfaces.Component) {
	header.Content.Left = components
}

func (header *Header) SetContentRight(components []interfaces.Component) {
	header.Content.Right = components
}

func (header *Header) String() string {

	html := "<header"

	if header.Layout != types.LayoutFlex {
		html += " data-layout=\"" + header.Layout.String() + "\""
	}

	html += ">"
	html += "<div>"

	if len(header.Content.Left) > 0 {

		for _, component := range header.Content.Left {
			html += component.String()
		}

	}

	html += "</div>"
	html += "<ul>"

	for _, item := range header.items {

		html += "<li"

		if item.Name == header.active {
			html += " data-state=\"active\""
		}

		html += ">"
		html += "<a data-view=\"" + item.Name + "\" href=\"" + item.Path + "\">" + item.Label + "</a>"
		html += "</li>"

	}

	html += "</ul>"
	html += "<div>"

	if len(header.Content.Right) > 0 {

		for _, component := range header.Content.Right {
			html += component.String()
		}

	}

	html += "</div>"
	html += "</header>"

	return html

}

func (header *Header) Unmount() bool {

	if header.Component.Element != nil {
		header.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}
