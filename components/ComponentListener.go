package components

import "github.com/cookiengineer/gooey/bindings/dom"

var component_listener_id uint = 0

type ComponentListener struct {
	Id       uint                      `json:"id"`
	Once     bool                      `json:"once"`
	Callback ComponentListenerCallback `json:"callback"`
	Listener *dom.EventListener        `json:"listener"`
}

type ComponentListenerCallback func(string)

func ToComponentListener(callback ComponentListenerCallback, once bool) ComponentListener {

	var listener ComponentListener

	listener.Id       = component_listener_id
	listener.Once     = once
	listener.Callback = callback
	listener.Listener = nil

	component_listener_id += 1

	return listener

}
