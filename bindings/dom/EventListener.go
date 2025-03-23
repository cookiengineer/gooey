//go:build wasm

package dom

import "syscall/js"

var event_listener_id uint = 0

type EventListener struct {
	Id       uint                  `json:"id"`
	Callback EventListenerCallback `json:"callback"`
	Function *js.Func              `json:"function"`
}

type EventListenerCallback func(Event)

func ToEventListener(callback EventListenerCallback) EventListener {

	var listener EventListener

	listener.Id = event_listener_id
	listener.Callback = callback

	event_listener_id += 1

	return listener

}
