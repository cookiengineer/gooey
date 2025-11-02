package layout

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/types"
import "github.com/cookiengineer/gooey/components/ui"
import "fmt"

type aside_menu_item struct {
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
	categories []string
	menus      map[string][]*aside_menu_item
}

func NewAside() Aside {

	var aside Aside

	element := dom.Document.CreateElement("aside")
	component := components.NewComponent(element)

	aside.Component = &component
	aside.Layout = types.LayoutFlex
	aside.Content.Top = make([]interfaces.Component, 0)
	aside.Content.Bottom = make([]interfaces.Component, 0)
	aside.menus = make(map[string][]*aside_menu_item, 0)

	return aside

}

func ToAside(element *dom.Element) *Aside {

	var aside Aside

	component := components.NewComponent(element)

	aside.Component = &component
	aside.Layout = types.LayoutFlex
	aside.Content.Top = make([]interfaces.Component, 0)
	aside.Content.Bottom = make([]interfaces.Component, 0)
	aside.menus = make(map[string][]*aside_menu_item, 0)

	return &aside

}

func (aside *Aside) Disable() bool {

	var result bool

	if len(aside.Content.Top) > 0 || len(aside.Content.Bottom) > 0 {

		for _, component := range aside.Content.Top {
			component.Disable()
		}

		for _, component := range aside.Content.Bottom {
			component.Disable()
		}

		result = true

	}

	return result

}

func (aside *Aside) Enable() bool {

	var result bool

	if len(aside.Content.Top) > 0 || len(aside.Content.Bottom) > 0 {

		for _, component := range aside.Content.Top {
			component.Enable()
		}

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

		if len(tmp) == 2 && tmp[0].TagName == "DIV" && tmp[1].TagName == "DIV" {

			lists_top := tmp[0].QuerySelectorAll("ul")

			for _, list := range lists_top {

				// TODO: Parse Top Menus into categories
				// TODO: ul data-category="..."?
				// TODO: How to sort correctly?

				if 1 == 2 {
					fmt.Println(list)
				}

			}

			buttons_bottom := tmp[1].QuerySelectorAll("button")
			content_bottom := make([]interfaces.Component, 0)

			for _, button := range buttons_bottom {
				content_bottom = append(content_bottom, ui.ToButton(button))
			}

			return true

		} else {

			console.Group("Aside: Invalid Markup")
			console.Error("Expected <ul></ul><div></div>")
			console.Error(aside.Component.Element.InnerHTML)
			console.GroupEnd("Aside: Invalid Markup")

			return false

		}

	} else {
		return false
	}

}
