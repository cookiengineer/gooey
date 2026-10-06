//go:build wasm

package bindings

import "github.com/cookiengineer/gooey/bindings/dom"
import "syscall/js"

type Screen struct {
	listeners   map[dom.EventType][]*dom.EventListener `json:"listeners"`
	Width       uint                                   `json:"width"`
	Height      uint                                   `json:"height"`
	AvailWidth  uint                                   `json:"availWidth"`
	AvailHeight uint                                   `json:"availHeight"`
	ColorDepth  uint                                   `json:"colorDepth"`
	PixelDepth  uint                                   `json:"pixelDepth"`
	IsExtended  bool                                   `json:"isExtended"`
	Orientation *ScreenOrientation                     `json:"orientation"`
	Value       *js.Value                              `json:"value"`
}

// Adds an Event Listener to the Screen. Only the "change" EventType is supported.
func (screen *Screen) AddEventListener(typ dom.EventType, listener *dom.EventListener) bool {

	var result bool

	check := string(typ)

	if check == "change" {

		wrapped_type := js.ValueOf(string(typ))
		wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

			if len(args) > 0 {

				event := args[0]

				if !event.IsNull() && !event.IsUndefined() {

					wrapped_event := dom.ToEvent(event)
					listener.Callback(wrapped_event)

				}

			}

			return nil

		})
		wrapped_capture := js.ValueOf(true)

		screen.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
		listener.Function = &wrapped_callback

		_, ok := screen.listeners[typ]

		if ok == true {
			screen.listeners[typ] = append(screen.listeners[typ], listener)
			result = true
		} else {
			screen.listeners[typ] = make([]*dom.EventListener, 0)
			screen.listeners[typ] = append(screen.listeners[typ], listener)
			result = true
		}

	}

	return result

}

// Removes an Event Listener from the Screen. If listener is set to nil, all EventListeners are
// removed.
func (screen *Screen) RemoveEventListener(typ dom.EventType, listener *dom.EventListener) bool {

	var result bool

	if listener != nil {

		if listener.Function != nil {

			wrapped_type := js.ValueOf(string(typ))
			wrapped_callback := *listener.Function
			wrapped_capture := js.ValueOf(true)
			screen.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			listener.Function.Release()
			listener.Function = nil

			listeners, ok := screen.listeners[typ]

			if ok == true {

				for l := 0; l < len(listeners); l++ {

					if listeners[l].Id == listener.Id {
						screen.listeners[typ] = append(listeners[:l], listeners[l+1:]...)
						break
					}

				}

			}

			result = true

		}

	} else {

		listeners, ok := screen.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]

				if listener.Function != nil {

					wrapped_type := js.ValueOf(string(typ))
					wrapped_callback := *listener.Function
					wrapped_capture := js.ValueOf(true)
					screen.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

					listener.Function.Release()
					listener.Function = nil

				}

			}

			delete(screen.listeners, typ)

			result = true

		}

	}

	return result

}
