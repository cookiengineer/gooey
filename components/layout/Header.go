package layout

import "github.com/cookiengineer/gooey/bindings/console"

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "sort"
import "strings"

type header_view_item struct {
	Name    string
	Label   string
	Path    string
	Element *dom.Element
}

type Header struct {
	Layout types.Layout `json:"layout"`
	View   string       `json:"view"`
	Content struct {
		Left  []interfaces.Component `json:"left"`
		Right []interfaces.Component `json:"right"`
	} `json:"content"`
	Component *components.Component `json:"component"`
	views map[string]*header_view_item
}

func NewHeader() Header {

	var header Header

	element   := bindings.Document.CreateElement("header")
	component := components.NewComponent(element)

	element.SetAttribute("data-layout", types.LayoutFlex.String())

	header.Component     = &component
	header.Layout        = types.LayoutFlex
	header.Content.Left  = make([]interfaces.Component, 0)
	header.Content.Right = make([]interfaces.Component, 0)
	header.View          = ""
	header.views         = make(map[string]*header_view_item)

	header.Component.InitEvent("click")
	header.Component.InitEvent("change-view")
	header.Component.InitEvent("action")

	header.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

		_, ok1 := attributes["data-action"]
		_, ok2 := attributes["data-view"]

		if ok1 == true {

			header.Component.FireEventListeners("action", map[string]string{
				"action": attributes["data-action"],
			})

		} else if ok2 == true {

			header.Component.FireEventListeners("change-view", map[string]string{
				"name": attributes["data-view"],
				"path": attributes["href"],
			})

		}

	}, false))

	header.Render()

	return header

}

func ToHeader(element *dom.Element) Header {

	var header Header

	component := components.NewComponent(element)

	header.Component     = &component
	header.Layout        = types.Layout(element.GetAttribute("data-layout"))
	header.Content.Left  = make([]interfaces.Component, 0)
	header.Content.Right = make([]interfaces.Component, 0)
	header.views         = make(map[string]*header_view_item)

	header.Parse()

	header.Component.InitEvent("click")
	header.Component.InitEvent("change-view")
	header.Component.InitEvent("action")

	header.Component.AddEventListener("click", components.ToComponentListener(func(event string, attributes map[string]string) {

		_, ok1 := attributes["data-action"]
		_, ok2 := attributes["data-view"]

		if ok1 == true {

			header.Component.FireEventListeners("action", map[string]string{
				"action": attributes["data-action"],
			})

		} else if ok2 == true {

			header.Component.FireEventListeners("change-view", map[string]string{
				"name": attributes["data-view"],
				"path": attributes["href"],
			})

		}

	}, false))

	return header

}

func (header *Header) ChangeView(name string) {

	_, ok := header.views[name]

	if ok == true {

		for other_name, item := range header.views {

			if other_name == name {

				if item.Element != nil {
					item.Element.SetClassName("active")
				}

			} else {

				if item.Element != nil {
					item.Element.SetClassName("")
				}

			}

		}

		header.View = name

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

func (header *Header) Parse() {

	if header.Component.Element != nil {

		tmp := header.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 3 && tmp[0].TagName == "DIV" && tmp[1].TagName == "UL" && tmp[2].TagName == "DIV" {

			buttons_left := tmp[0].QuerySelectorAll("button")

			for _, button := range buttons_left {
				component := ui.ToButton(button)
				header.Content.Left = append(header.Content.Left, &component)
			}

			items_center := tmp[1].QuerySelectorAll("li")

			for _, item := range items_center {

				link := item.QuerySelector("a")

				if link != nil {

					view_item := header_view_item{
						Name:    link.GetAttribute("data-view"),
						Label:   strings.TrimSpace(link.TextContent),
						Path:    link.GetAttribute("href"),
						Element: item,
					}

					if view_item.Name != "" {

						if item.ClassName == "active" {
							header.View = view_item.Name
							header.views[view_item.Name] = &view_item
						} else {
							header.views[view_item.Name] = &view_item
						}

					}

				}

			}

			buttons_right := tmp[2].QuerySelectorAll("button")

			for _, button := range buttons_right {
				component := ui.ToButton(button)
				header.Content.Right = append(header.Content.Right, &component)
			}

		} else {

			console.Group("Header: Invalid Markup")
			console.Error("Expected <div></div><ul></ul><div></div>")
			console.Error(header.Component.Element.InnerHTML)
			console.GroupEnd("Header: Invalid Markup")

		}

	}

}

func (header *Header) Render() *dom.Element {

	if header.Component.Element != nil {

		tmp := header.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 0 {
			header.Component.Element.SetInnerHTML("<div></div><ul></ul><div></div>")
			tmp = header.Component.Element.QuerySelectorAll("div, ul")
		}

		if len(tmp) == 3 {

			header.Component.Element.SetAttribute("data-layout", header.Layout.String())

			elements_left   := make([]*dom.Element, 0)
			elements_center := make([]*dom.Element, 0)
			elements_right  := make([]*dom.Element, 0)

			for _, component := range header.Content.Left {
				elements_left = append(elements_left, component.Render())
			}

			for _, item := range header.views {

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

func (header *Header) SetLeft(components []interfaces.Component) {
	header.Content.Left = components
}

func (header *Header) SetRight(components []interfaces.Component) {
	header.Content.Right = components
}

func (header *Header) SetView(view interfaces.View) {

	name  := view.GetProperty("Name")
	label := view.GetProperty("Label")
	path  := view.GetProperty("Path")

	if name != "" && label != "" && path != "" {

		item, ok := header.views[name]

		if ok == true {

			item.Name  = name
			item.Label = label
			item.Path  = path

		} else {

			item := header_view_item{
				Name:    name,
				Label:   label,
				Path:    path,
				Element: bindings.Document.CreateElement("li"),
			}

			header.views[name] = &item

		}

	}

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

	if len(header.views) > 0 {

		paths := make([]string, 0)

		for _, item := range header.views {
			paths = append(paths, item.Path)
		}

		sort.Strings(paths)

		for p := 0; p < len(paths); p++ {

			item := getHeaderItemByPath(paths[p], header.views)

			html += "<li"

			if item.Name == header.View {
				html += " class=\"active\""
			}

			html += ">"
			html += "<a data-view=\"" + item.Name + "\" href=\"" + item.Path + "\">" + item.Label + "</a>"
			html += "</li>"

		}

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
