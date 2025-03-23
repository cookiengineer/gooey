package components

import gooey "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/dom"

type Component struct {
	listeners map[string][]*ComponentListener `json:"listeners"`
	Element   *dom.Element                    `json:"element"`
}

func (component *Component) Init(element *dom.Element) {

	component.listeners = make(map[string][]*ComponentListener, 0)

	if element != nil {
		component.Element = element
	} else {
		component.Element = gooey.Document.CreateElement("gooey-component")
	}

}

func (component *Component) InitEvent(event string) {

	_, ok := component.listeners[event]

	if ok == false {
		component.listeners[event] = make([]*ComponentListener, 0)
	}

}

func (component *Component) AddEventListener(event string, listener ComponentListener) bool {

	var result bool = false

	_, ok := component.listeners[event]

	if ok == true {

		if event == "click" || event == "change" {

			if component.Element != nil {

				wrapped_listener := dom.ToEventListener(func(_ dom.Event) {
					component.FireEventListeners(event)
				})

				component.Element.AddEventListener(dom.EventType(event), wrapped_listener)
				listener.Listener = &wrapped_listener

			}

			component.listeners[event] = append(component.listeners[event], &listener)
			result = true

		} else {
			component.listeners[event] = append(component.listeners[event], &listener)
			result = true
		}

	}

	return result

}

func (component *Component) FireEventListeners(event string) bool {

	var result bool = false

	listeners, ok := component.listeners[event]

	if ok == true {

		indexes := make([]int, 0)

		for l := 0; l < len(listeners); l++ {

			listener := listeners[l]
			listener.Callback(event)

			if listener.Once == true {
				indexes = append(indexes, l)
			}

		}

		if len(indexes) > 0 {

			for _, index := range indexes {
				listeners = append(listeners[:index], listeners[index+1:]...)
			}

			component.listeners[event] = listeners

		}

	}

	return result

}

func (component *Component) RemoveEventListener(event string, listener *ComponentListener) bool {

	var result bool = false

	if listener != nil {

		listeners, ok := component.listeners[event]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {

				listener := component.listeners[event][index]

				if component.Element != nil && listener.Listener != nil {
					component.Element.RemoveEventListener(dom.EventType(event), listener.Listener)
				}

				component.listeners[event] = append(component.listeners[event][:index], component.listeners[event][index+1:]...)
				result = true

			}

		}

	} else {

		_, ok := component.listeners[event]

		if ok == true {

			if component.Element != nil {
				component.Element.RemoveEventListener(dom.EventType(event), nil)
			}

			component.listeners[event] = make([]*ComponentListener, 0)
			result = true

		}

	}

	return result

}

func (component *Component) Render() {
	// Render into dom Element
}

func (component *Component) String() string {
	return ""
}
