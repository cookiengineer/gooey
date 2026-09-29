//go:build wasm

package webgl

import "syscall/js"

type Event struct {
	StatusMessage string    `json:"statusMessage"`
	Type          EventType `json:"type"`
	Value         *js.Value `json:"value"`
}

func ToEvent(value js.Value) *Event {

	event := &Event{}

	status_message := value.Get("statusMessage")

	if !status_message.IsNull() && !status_message.IsUndefined() {
		event.StatusMessage = status_message.String()
	}

	event.Type = EventType(value.Get("type").String())
	event.Value = &value

	return event

}

// The Web Browser Engine only restores a lost Context when the
// webglcontextlost Event was cancelled, so a listener that intends to recover
// has to call this.
func (event *Event) PreventDefault() {
	event.Value.Call("preventDefault")
}

func (event *Event) StopImmediatePropagation() {
	event.Value.Call("stopImmediatePropagation")
}

func (event *Event) StopPropagation() {
	event.Value.Call("stopPropagation")
}
