package components

import "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

type Component struct {
	Listeners map[string][]*EventListener `json:"listeners"`
	Element   *dom.Element                `json:"element"`
}

func NewComponent(element *dom.Element) Component {

	var component Component

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

func (component *Component) AddEventListener(event string, listener EventListener) bool {

	var result bool

	_, ok := component.Listeners[event]

	if ok == true {

		if event == "click" {

			if component.Element != nil {

				wrapped_listener := dom.ToEventListener(func(dom_event dom.Event) {

					attributes := make(map[string]string)

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
				listener.Listener = &wrapped_listener

			}

			component.Listeners[event] = append(component.Listeners[event], &listener)
			result = true

		} else if event == "change" {

			if component.Element != nil {

				wrapped_listener := dom.ToEventListener(func(dom_event dom.Event) {

					attributes := make(map[string]string)

					if dom_event.Target != nil {

						dom_event.Target.RefreshAttributes()

						for key, val := range dom_event.Target.Attributes {
							attributes[key] = val
						}

					}

					component.FireEventListeners(event, attributes)

				})

				component.Element.AddEventListener(dom.EventType(event), wrapped_listener)
				listener.Listener = &wrapped_listener

			}

			component.Listeners[event] = append(component.Listeners[event], &listener)
			result = true

		} else {
			component.Listeners[event] = append(component.Listeners[event], &listener)
			result = true
		}

	}

	return result

}

func (component *Component) FireEventListeners(event string, attributes map[string]string) bool {

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

func (component *Component) String() string {

	html := ""

	if component.Element != nil {

		tagname := strings.ToLower(component.Element.TagName)

		html += "<" + tagname

		// TODO: attributes?

		// TODO: Is disabled attribute actually mapped as an attribute with value "" in the HTMLElement API?

		html += ">"

		// TODO: TextContent? InnerHTML?

		html += "</" + tagname + ">"

	}

	return html

}

