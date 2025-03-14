//go:build wasm

package history

import "syscall/js"

type PopStateEvent struct {
	Bubbles               bool       `json:"bubbles"`
	Cancelable            bool       `json:"cancelable"`
	Composed              bool       `json:"composed"`
	DefaultPrevented      bool       `json:"defaultPrevented"`
	EventPhase            EventPhase `json:"eventPhase"`
	HasUAVisualTransition bool       `json:"hasUAVisualTransition"`
	IsTrusted             bool       `json:"isTrusted"`
	Type                  EventType  `json:"type"`
	Value                 *js.Value  `json:"value"`
	// TODO: State *any `json:"state"`
}

func ToPopStateEvent(value js.Value) PopStateEvent {

	var event PopStateEvent

	event.Bubbles = value.Get("bubbles").Bool()
	event.Cancelable = value.Get("cancelable").Bool()
	event.Composed = value.Get("composed").Bool()
	event.DefaultPrevented = value.Get("defaultPrevented").Bool()
	event.EventPhase = EventPhase(value.Get("eventPhase").Int())
	event.HasUAVisualTransition = value.Get("hasUAVisualTransition").Bool()
	event.IsTrusted = value.Get("isTrusted").Bool()
	event.Type = EventType(value.Get("type").String())
	event.Value = &value

	// TODO: Target?

	return event

}

func (event *PopStateEvent) PreventDefault() {
	event.Value.Call("preventDefault")
}

func (event *PopStateEvent) StopImmediatePropagation() {
	event.Value.Call("stopImmediatePropagation")
}

func (event *PopStateEvent) StopPropagation() {
	event.Value.Call("stopPropagation")
}
