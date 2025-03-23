//go:build wasm

package dom

import "syscall/js"

type Event struct {
	Bubbles          bool       `json:"bubbles"`
	Cancelable       bool       `json:"cancelable"`
	Composed         bool       `json:"composed"`
	DefaultPrevented bool       `json:"defaultPrevented"`
	EventPhase       EventPhase `json:"eventPhase"`
	IsTrusted        bool       `json:"isTrusted"`
	Target           *Element   `json:"target"`
	// TimeStamp int `json:"timeStamp"`
	Type  EventType `json:"type"`
	Value *js.Value `json:"value"`
}

func ToEvent(value js.Value) Event {

	var event Event

	event.Bubbles = value.Get("bubbles").Bool()
	event.Cancelable = value.Get("cancelable").Bool()
	event.Composed = value.Get("composed").Bool()
	event.DefaultPrevented = value.Get("defaultPrevented").Bool()
	event.EventPhase = EventPhase(value.Get("eventPhase").Int())
	event.IsTrusted = value.Get("isTrusted").Bool()
	event.Type = EventType(value.Get("type").String())

	target := ToElement(value.Get("target"))

	event.Target = &target
	event.Value = &value

	return event

}

func (event *Event) PreventDefault() {
	event.Value.Call("preventDefault")
}

func (event *Event) StopImmediatePropagation() {
	event.Value.Call("stopImmediatePropagation")
}

func (event *Event) StopPropagation() {
	event.Value.Call("stopPropagation")
}
