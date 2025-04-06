package components

import "github.com/cookiengineer/gooey/bindings/dom"

var component_listener_id uint = 0

type EventListener struct {
	Id       uint                  `json:"id"`
	Once     bool                  `json:"once"`
	Callback EventListenerCallback `json:"callback"`
	Listener *dom.EventListener    `json:"listener"`
}

type EventListenerCallback func(string, map[string]string)

func ToEventListener(callback EventListenerCallback, once bool) EventListener {

	var listener EventListener

	listener.Id       = component_listener_id
	listener.Once     = once
	listener.Callback = callback
	listener.Listener = nil

	component_listener_id += 1

	return listener

}
