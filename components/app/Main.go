//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/location"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "strings"

type controller_constructor func(*Main,*View)  interfaces.Controller
type view_constructor       func(*dom.Element) interfaces.View

type Main struct {

	Client             *Client                           `json:"client"`
	Storage            *Storage                          `json:"storage"`

	Header             *layout.Header                    `json:"header"`
	Footer             *layout.Footer                    `json:"footer"`
	Dialog             *layout.Dialog                    `json:"dialog"`

	Controller         interfaces.Controller             `json:"controller"`
	ControllerRegistry map[string]controller_constructor `json:"controller_registry"`
	View               interfaces.View                   `json:"view"`
	ViewRegistry       map[string]view_constructor       `json:"view_registry"`
	Document           *components.Document              `json:"document"`
	Element            *dom.Element                      `json:"element"`

	controllers        map[string]interfaces.Controller  `json:"-"`
	views              map[string]interfaces.View        `json:"-"`

}

func NewMain() *Main {

	var main Main

	main.Client  = NewClient()
	main.Storage = NewStorage()

	main.Controller         = nil
	main.ControllerRegistry = make(map[string]controller_constructor)
	main.controllers        = make(map[string]interfaces.Controller)

	main.View         = nil
	main.ViewRegistry = make(map[string]view_constructor)
	main.views        = make(map[string]interfaces.View)

	main.Document = components.NewDocument()
	main.Element  = main.Document.QuerySelector("main")

	return &main

}

func ToMain(document *components.Document) *Main {

	var main Main

	main.Client  = NewClient()
	main.Storage = NewStorage()

	main.Controller         = nil
	main.ControllerRegistry = make(map[string]controller_constructor)
	main.controllers        = make(map[string]interfaces.Controller)

	main.View         = nil
	main.ViewRegistry = make(map[string]view_constructor)
	main.views        = make(map[string]interfaces.View)

	main.Document = document
	main.Element  = main.Document.QuerySelector("main")

	return &main

}

func (main *Main) ChangeView(name string) bool {

	var result bool

	view, ok := main.views[name]

	if ok == true {

		if main.View != nil {
			main.View.Leave()
			main.View = nil
		}

		main.Element.SetAttribute("data-view", name)

		if main.Header != nil {
			main.Header.ChangeView(name)
		}

		main.View = view
		main.View.Enter()

		result = true

	}

	return result

}

func (main *Main) Mount() bool {

	main.Document.Register("header",  components.Wrap(layout.ToHeader))
	main.Document.Register("footer",  components.Wrap(layout.ToFooter))
	main.Document.Register("dialog",  components.Wrap(layout.ToDialog))
	main.Document.Register("section", components.Wrap(ToView))

	// Don't mount Components
	main.Document.Mount()

	header, ok1 := components.Unwrap[*layout.Header](main.Document.QueryComponent("body > header"))

	if header != nil && ok1 == true {

		main.Header = header
		main.Header.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]any) {

			name, ok1 := attributes["name"].(string)
			path, ok2 := attributes["path"].(string)

			if ok1 == true && ok2 == true {

				_, ok3 := main.views[name]

				if ok3 == true {

					// Single-page web app
					main.ChangeView(name)

				} else {

					// TODO: History API integration?
					// Multi-page web app
					location.Location.Replace(path)

				}

			}

		}, false))

	} else {
		main.Header = nil
	}

	footer, ok2 := components.Unwrap[*layout.Footer](main.Document.QueryComponent("body > footer"))

	if footer != nil && ok2 == true {
		main.Footer = footer
	} else {
		main.Footer = nil
	}

	dialog, ok3 := components.Unwrap[*layout.Dialog](main.Document.QueryComponent("body > dialog"))

	if dialog != nil && ok3 == true {
		main.Dialog = dialog
	} else {
		main.Dialog = nil
	}

	main_component, ok4 := components.Unwrap[*components.Component](main.Document.QueryComponent("body > main"))

	if main_component != nil && ok4 == true {

		if len(main_component.Content) > 0 {

			for _, component := range main_component.Content {

				view, ok0 := component.(*View)

				if ok0 == true {

					name  := view.Name()
					label := view.Label()
					path  := view.Path()

					if name != "" && label != "" && path != "" {

						view_wrapper, ok1 := main.ViewRegistry[name]

						if ok1 == true {

							custom_view := view_wrapper(view.Element)

							if custom_view != nil {
								main.views[name] = custom_view
							} else {
								main.views[name] = view
							}

							main.views[name].Mount()

						} else {

							main.views[name] = view
							main.views[name].Mount()

						}

						if main.Header != nil {
							main.Header.RegisterView(main.views[name])
						}

						controller_wrapper, ok2 := main.ControllerRegistry[name]

						if ok2 == true {

							controller := controller_wrapper(main, view)

							if controller != nil {
								main.controllers[name] = controller
							}

						}

					} else {
						console.Group("View: Invalid Markup")
						console.Error("Expected <section data-name=\"...\" data-label=\"...\" data-path=\"...\"></section>")
						console.GroupEnd("View: Invalid Markup")
					}

				} else {
					console.Group("Main: Invalid Markup")
					console.Error("Expected <main><section data-name=\"...\" data-label=\"...\" data-path=\"...\"></section></main>")
					console.GroupEnd("Main: Invalid Markup")
				}

			}

		}

	}

	if main.Element != nil {

		default_view := main.Element.GetAttribute("data-view")

		if default_view != "" {
			main.ChangeView(default_view)
		}

		return true

	}

	return false

}

func (main *Main) QueryComponent(view_name string, query string) interfaces.Component {

	view, ok := main.views[view_name]

	if ok == true {

		selectors := utils.SplitQuery(query)

		if len(selectors) > 0 && strings.HasPrefix(selectors[0], "section") == true {

			return view.Query(utils.JoinQuery(selectors))

		} else if len(selectors) > 0 {

			tmp_selectors := make([]string, 0)
			tmp_selectors = append(tmp_selectors, "section")
			tmp_selectors = append(tmp_selectors, selectors...)

			return view.Query(utils.JoinQuery(tmp_selectors))

		}

	}

	return nil

}

func (main *Main) QuerySelector(view_name string, query string) *dom.Element {

	_, ok := main.views[view_name]

	if ok == true {

		selectors := utils.SplitQuery(query)

		if len(selectors) > 0 && strings.HasPrefix(selectors[0], "section") == true {

			return main.Document.QuerySelector(query)

		} else {

			tmp_selectors := make([]string, 0)
			tmp_selectors = append(tmp_selectors, "section[data-name=\"" + view_name + "\"]")
			tmp_selectors = append(tmp_selectors, selectors...)

			return main.Document.QuerySelector(utils.JoinQuery(tmp_selectors))

		}

	}

	return nil

}

func (main *Main) QuerySelectorAll(view_name string, query string) []*dom.Element {

	result := make([]*dom.Element, 0)

	_, ok := main.views[view_name]

	if ok == true {

		selectors := utils.SplitQuery(query)

		if len(selectors) > 0 && strings.HasPrefix(selectors[0], "section") == true {

			result = main.Document.QuerySelectorAll(query)

		} else {

			tmp_selectors := make([]string, 0)
			tmp_selectors = append(tmp_selectors, "section[data-name=\"" + view_name + "\"]")
			tmp_selectors = append(tmp_selectors, selectors...)

			result = main.Document.QuerySelectorAll(utils.JoinQuery(tmp_selectors))

		}

	}

	return result

}

func (main *Main) RegisterController(name string, wrapper controller_constructor) {
	main.ControllerRegistry[strings.ToLower(name)] = wrapper
}

func (main *Main) RegisterView(name string, wrapper view_constructor) {
	main.ViewRegistry[strings.ToLower(name)] = wrapper
}

func (main *Main) Render() {

	if main.Header != nil {
		main.Header.Render()
	}

	if main.View != nil {
		main.View.Render()
	}

	if main.Footer != nil {
		main.Footer.Render()
	}

	if main.Dialog != nil {
		main.Dialog.Render()
	}

}

func (main *Main) Unmount() bool {

	if main.Header != nil {
		main.Header.Component.RemoveEventListener("change-view", nil)
	}

	return true

}
