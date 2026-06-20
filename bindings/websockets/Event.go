//go:build wasm

package websockets

import "syscall/js"

type Event struct {
	Data   []byte    `json:"data"`
	Origin string    `json:"origin"`
	Type   EventType `json:"type"`
	Value  *js.Value `json:"value"`
}

func ToEvent(value js.Value) *Event {

	event := &Event{}
	data  := value.Get("data")

	if !data.IsNull() && !data.IsUndefined() {

		switch {

		// Text Frames
		case data.Type() == js.TypeString:

			event.Data = []byte(data.String())

		// Binary Frames
		case data.InstanceOf(js.Global().Get("ArrayBuffer")):

			array := js.Global().Get("Uint8Array").New(data)
			event.Data = make([]byte, array.Length())
			js.CopyBytesToGo(event.Data, array)

		// Unsupported Frames
		default:

			event.Data = make([]byte, 0)

		}

	} else {

		event.Data = make([]byte, 0)

	}

	event.Origin = value.Get("origin").String()
	event.Type = EventType(value.Get("type").String())

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
