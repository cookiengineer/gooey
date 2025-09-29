package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "sort"
import "strings"

type Component struct {
	Content   []interfaces.Component      `json:"content"`
	Listeners map[string][]*EventListener `json:"listeners"`
	Element   *dom.Element                `json:"element"`
}

func NewComponent(element *dom.Element) Component {

	var component Component

	component.Content = make([]interfaces.Component, 0)
	component.Element = element
	component.Listeners = make(map[string][]*EventListener, 0)

	return component

}

func (component *Component) Disable() bool {

	var result bool

	if component.Element != nil {
		component.Element.SetAttribute("disabled", "")
	}

	return result

}

func (component *Component) Enable() bool {

	var result bool

	if component.Element != nil {
		component.Element.RemoveAttribute("disabled")
	}

	return result

}

func (component *Component) AddEventListener(event string, listener *EventListener) bool {

	var result bool

	if listener != nil {

		_, ok := component.Listeners[event]

		if ok == true {

			if event == "click" {

				if component.Element != nil {

					wrapped_listener := dom.ToEventListener(func(dom_event *dom.Event) {

						attributes := make(map[string]any)

						if dom_event.Target != nil {

							dom_event.Target.RefreshAttributes()

							for key, val := range dom_event.Target.Attributes {
								attributes[key] = val
							}

						}

						component.FireEventListeners(event, attributes)

						// XXX: This prevents <a> elements triggering History navigation
						dom_event.PreventDefault()
						dom_event.StopPropagation()

					})

					component.Element.AddEventListener(dom.EventType(event), wrapped_listener)
					listener.Listener = wrapped_listener

				}

				component.Listeners[event] = append(component.Listeners[event], listener)
				result = true

			} else if event == "change" {

				if component.Element != nil {

					wrapped_listener := dom.ToEventListener(func(dom_event *dom.Event) {

						attributes := make(map[string]any)

						if dom_event.Target != nil {

							dom_event.Target.RefreshAttributes()

							for key, val := range dom_event.Target.Attributes {
								attributes[key] = val
							}

						}

						component.FireEventListeners(event, attributes)

					})

					component.Element.AddEventListener(dom.EventType(event), wrapped_listener)
					listener.Listener = wrapped_listener

				}

				component.Listeners[event] = append(component.Listeners[event], listener)
				result = true

			} else {
				component.Listeners[event] = append(component.Listeners[event], listener)
				result = true
			}

		}

	}

	return result

}

func (component *Component) FireEventListeners(event string, attributes map[string]any) bool {

	var result bool

	listeners, ok := component.Listeners[event]

	if ok == true {

		indexes := make([]int, 0)

		for l := 0; l < len(listeners); l++ {

			listener := listeners[l]
			listener.Callback(event, attributes)

			if listener.Once == true {
				indexes = append(indexes, l)
			}

		}

		if len(indexes) > 0 {

			for _, index := range indexes {
				listeners = append(listeners[:index], listeners[index+1:]...)
			}

			component.Listeners[event] = listeners

		}

	}

	return result

}

func (component *Component) HasEvent(event string) bool {

	var result bool

	_, ok := component.Listeners[event]

	if ok == true {
		result = true
	}

	return result

}

func (component *Component) InitEvent(event string) {

	_, ok := component.Listeners[event]

	if ok == false {
		component.Listeners[event] = make([]*EventListener, 0)
	}

}

func (component *Component) Mount() bool {
	return true
}

func (component *Component) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) >= 2 {

		if component.Element != nil {

			if utils.MatchesQuery(component.Element, selectors[0]) == true {

				tmp_query := utils.JoinQuery(selectors[1:])

				for _, content := range component.Content {

					tmp_component := content.Query(tmp_query)

					if tmp_component != nil {
						return tmp_component
					}

				}

			}

		}

	} else if len(selectors) == 1 {

		if component.Element != nil {

			if utils.MatchesQuery(component.Element, selectors[0]) == true {
				return component
			}

		}

	}

	return nil

}

func (component *Component) Render() *dom.Element {
	return component.Element
}

func (component *Component) RemoveEventListener(event string, listener *EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := component.Listeners[event]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {

				listener := component.Listeners[event][index]

				if component.Element != nil && listener.Listener != nil {
					component.Element.RemoveEventListener(dom.EventType(event), listener.Listener)
				}

				component.Listeners[event] = append(component.Listeners[event][:index], component.Listeners[event][index+1:]...)
				result = true

			}

		}

	} else {

		_, ok := component.Listeners[event]

		if ok == true {

			if component.Element != nil {
				component.Element.RemoveEventListener(dom.EventType(event), nil)
			}

			component.Listeners[event] = make([]*EventListener, 0)
			result = true

		}

	}

	return result

}

func (component *Component) SetContent(components []interfaces.Component) {
	component.Content = components
}

func (component *Component) String() string {

	html := ""

	if component.Element != nil {

		tagname := strings.ToLower(component.Element.TagName)

		html += "<" + tagname

		attributes := make([]string, 0)

		for key, _ := range component.Element.Attributes {
			attributes = append(attributes, key)
		}

		sort.Strings(attributes)

		for _, attribute := range attributes {
			html += " " + attribute + "=\"" + component.Element.Attributes[attribute] + "\""
		}

		html += ">"
		html += component.Element.InnerHTML
		html += "</" + tagname + ">"

	}

	return html

}

func (component *Component) Unmount() bool {
	return true
}
