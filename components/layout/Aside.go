//go:build wasm

package layout

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/types"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "strings"

type aside_item struct {
	Name    string
	Label   string
	Path    string
	Element *dom.Element
}

type Aside struct {
	Layout  types.Layout `json:"layout"`
	Content struct {
		Top    []interfaces.Component `json:"top"`
		Bottom []interfaces.Component `json:"bottom"`
	} `json:"content"`
	Component *components.Component `json:"component"`
	active    string                `json:"-"`
	items     []*aside_item         `json:"-"`
}

func NewAside() Aside {

	var aside Aside

	element := dom.GetDocument().CreateElement("aside")
	component := components.NewComponent(element)

	aside.Component = &component
	aside.Layout = types.LayoutFlex
	aside.Content.Bottom = make([]interfaces.Component, 0)
	aside.active = ""
	aside.items = make([]*aside_item, 0)

	return aside

}

func ToAside(element *dom.Element) *Aside {

	var aside Aside

	component := components.NewComponent(element)

	aside.Component = &component
	aside.Layout = types.LayoutFlex
	aside.Content.Bottom = make([]interfaces.Component, 0)
	aside.active = ""
	aside.items = make([]*aside_item, 0)

	return &aside

}

func (aside *Aside) ChangeView(name string) bool {

	var found *aside_item = nil

	for _, item := range aside.items {

		if item.Name == name {
			found = item
			break
		}

	}

	if found != nil {

		for _, item := range aside.items {

			if item.Element != nil {

				if item.Name == found.Name {
					item.Element.SetAttribute("data-state", "active")
				} else {
					item.Element.RemoveAttribute("data-state")
				}

			}

		}

		aside.active = name

		return true

	}

	return false

}

func (aside *Aside) Disable() bool {

	var result bool

	if len(aside.Content.Bottom) > 0 {

		for _, component := range aside.Content.Bottom {
			component.Disable()
		}

		result = true

	}

	return result

}

func (aside *Aside) Enable() bool {

	var result bool

	if len(aside.Content.Bottom) > 0 {

		for _, component := range aside.Content.Bottom {
			component.Enable()
		}

		result = true

	}

	return result

}

func (aside *Aside) Mount() bool {

	if aside.Component != nil {

		aside.Component.InitEvent("change-view")
		aside.Component.InitEvent("action")

	}

	if aside.Component.Element != nil {

		aside.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target != nil {

				action := event.Target.GetAttribute("data-action")
				view := event.Target.GetAttribute("data-view")
				path := event.Target.GetAttribute("href")

				if action != "" {

					event.PreventDefault()
					event.StopPropagation()

					aside.Component.FireEventListeners("action", map[string]any{
						"action": action,
					})

				} else if view != "" && path != "" {

					event.PreventDefault()
					event.StopPropagation()

					aside.ChangeView(view)
					aside.Component.FireEventListeners("change-view", map[string]any{
						"name": view,
						"path": path,
					})

				}

			}

		}))

		layout := aside.Component.Element.GetAttribute("data-layout")

		if layout != "" {
			aside.Layout = types.Layout(layout)
		}

		tmp := aside.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 3 && tmp[0].TagName == "UL" && tmp[1].TagName == "DIV" && tmp[2].TagName == "DIV" {

			elements_top := tmp[0].QuerySelectorAll("li")

			for _, element := range elements_top {

				link := element.QuerySelector("a")

				if link != nil {

					item := aside_item{
						Name:    link.GetAttribute("data-view"),
						Label:   strings.TrimSpace(link.TextContent),
						Path:    link.GetAttribute("href"),
						Element: element,
					}

					if item.Name != "" {

						if element.GetAttribute("data-state") == "active" {
							aside.active = item.Name
							aside.items = append(aside.items, &item)
						} else {
							aside.items = append(aside.items, &item)
						}

					}

				}

			}

			// Layout constraint: Middle DIV must be empty

			buttons_bottom := tmp[2].QuerySelectorAll("button")
			content_bottom := make([]interfaces.Component, 0)

			for _, button := range buttons_bottom {
				content_bottom = append(content_bottom, ui.ToButton(button))
			}

			aside.Content.Bottom = content_bottom

			for _, component := range aside.Content.Bottom {
				component.Mount()
			}

			return true

		} else {

			console1 := console.GetConsole()
			console1.Group("gooey/components/layout.Aside: Invalid markup")
			console1.Error("Expected <ul/><div/><div/> but got this instead:")
			console1.Error(aside.Component.Element.InnerHTML)
			console1.GroupEnd()

			return false

		}

	} else {
		return false
	}

}

func (aside *Aside) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if aside.Component.Element != nil {

			if utils.MatchesQuery(aside.Component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range aside.Content.Bottom {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if aside.Component.Element != nil {

			if utils.MatchesQuery(aside.Component.Element, selectors[0]) == true {
				return aside
			}

		}

	}

	return nil

}

func (aside *Aside) RegisterView(view interfaces.View) bool {

	name := view.Name()
	label := view.Label()
	path := view.Path()

	if name != "" && label != "" && path != "" {

		var found *aside_item = nil

		for _, item := range aside.items {

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

			item := aside_item{
				Name:    name,
				Label:   label,
				Path:    path,
				Element: dom.GetDocument().CreateElement("li"),
			}

			aside.items = append(aside.items, &item)

			return true

		}

	}

	return false

}

func (aside *Aside) Render() *dom.Element {

	if aside.Component.Element != nil {

		tmp := aside.Component.Element.QuerySelectorAll("div, ul")

		if len(tmp) == 0 {
			aside.Component.Element.SetInnerHTML("<ul></ul><div></div><div></div>")
			tmp = aside.Component.Element.QuerySelectorAll("div, ul")
		}

		if len(tmp) == 3 {

			if aside.Layout != types.LayoutFlex {
				aside.Component.Element.SetAttribute("data-layout", aside.Layout.String())
			}

			elements_top := make([]*dom.Element, 0)
			elements_middle := make([]*dom.Element, 0)
			elements_bottom := make([]*dom.Element, 0)

			for _, item := range aside.items {

				if item.Name == aside.active {
					item.Element.SetAttribute("data-state", "active")
				} else {
					item.Element.RemoveAttribute("data-state")
				}

				item.Element.SetInnerHTML("<a data-view=\"" + item.Name + "\" href=\"" + item.Path + "\">" + item.Label + "</a>")
				elements_top = append(elements_top, item.Element)

			}

			// Layout constraint: Middle DIV must be empty

			for _, component := range aside.Content.Bottom {
				elements_bottom = append(elements_bottom, component.Render())
			}

			tmp[0].ReplaceChildren(elements_top)
			tmp[1].ReplaceChildren(elements_middle)
			tmp[2].ReplaceChildren(elements_bottom)

		}

	}

	return aside.Component.Element

}

func (aside *Aside) SetContentBottom(components []interfaces.Component) {
	aside.Content.Bottom = components
}

func (aside *Aside) String() string {

	html := "<aside"

	if aside.Layout != types.LayoutFlex {
		html += " data-layout=\"" + aside.Layout.String() + "\""
	}

	html += ">"
	html += "<ul>"

	for _, item := range aside.items {

		html += "<li"

		if item.Name == aside.active {
			html += " data-state=\"active\""
		}

		html += ">"
		html += "<a data-view=\"" + item.Name + "\" href=\"" + item.Path + "\">" + item.Label + "</a>"
		html += "</li>"

	}

	html += "</ul>"
	html += "<div>"

	// Layout constraint: Middle DIV must be empty

	html += "</div>"
	html += "<div>"

	if len(aside.Content.Bottom) > 0 {

		for _, component := range aside.Content.Bottom {
			html += component.String()
		}

	}

	html += "</div>"
	html += "</aside>"

	return html

}

func (aside *Aside) Unmount() bool {

	if aside.Component.Element != nil {
		aside.Component.Element.RemoveEventListener("click", nil)
	}

	return true

}
