//go:build wasm

package gooey

import "github.com/cookiengineer/gooey/pkg/dom"
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

func (screen *Screen) AddEventListener(typ dom.EventType, listener dom.EventListener) bool {

	var result bool = false

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
			screen.listeners[typ] = append(screen.listeners[typ], &listener)
			result = true
		} else {
			screen.listeners[typ] = make([]*dom.EventListener, 0)
			screen.listeners[typ] = append(screen.listeners[typ], &listener)
			result = true
		}

	}

	return result

}

func (screen *Screen) RemoveEventListener(typ dom.EventType, listener *dom.EventListener) bool {

	var result bool = false

	if listener != nil {

		listeners, ok := screen.listeners[typ]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {

				listener := listeners[index]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				screen.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				screen.listeners[typ] = append(screen.listeners[typ][:index], screen.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := screen.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				screen.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(screen.listeners, typ)

			result = true

		}

	}

	return result

}
