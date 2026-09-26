//go:build wasm

package webgl

import "github.com/cookiengineer/gooey/bindings/dom"
import "syscall/js"

type Canvas struct {
	listeners map[EventType][]*EventListener `json:"listeners"`
	Width     uint                           `json:"width"`
	Height    uint                           `json:"height"`
	Element   *dom.Element                   `json:"element"`
	Context   *Context                       `json:"context"`
}

func ToCanvas(element *dom.Element) *Canvas {

	attributes := NewContextAttributes()

	return ToCanvasWithAttributes(element, &attributes)

}

func ToCanvasWithAttributes(element *dom.Element, attributes *ContextAttributes) *Canvas {

	if element != nil && element.TagName == "CANVAS" && attributes != nil {

		var canvas Canvas

		canvas.listeners = make(map[EventType][]*EventListener)
		canvas.Width = uint(element.Value.Get("width").Int())
		canvas.Height = uint(element.Value.Get("height").Int())
		canvas.Element = element

		tmp := element.Value.Call("getContext", "webgl2", attributes.ToValue())

		if !tmp.IsNull() && !tmp.IsUndefined() {
			canvas.Context = ToContext(tmp)
		}

		return &canvas

	}

	return nil

}

// The Context loss Events are dispatched at the Canvas Element, not at the
// Context, because a lost Context is exactly the object that is gone.
func (canvas *Canvas) AddEventListener(typ EventType, listener *EventListener) bool {

	var result bool

	if listener != nil {

		wrapped_type := js.ValueOf(string(typ))
		wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

			if len(args) > 0 {

				event := args[0]

				if !event.IsNull() && !event.IsUndefined() {

					wrapped_event := ToEvent(event)
					listener.Callback(wrapped_event)

				}

			}

			return nil

		})
		wrapped_capture := js.ValueOf(true)

		canvas.Element.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
		listener.Function = &wrapped_callback

		_, ok := canvas.listeners[typ]

		if ok == true {
			canvas.listeners[typ] = append(canvas.listeners[typ], listener)
			result = true
		} else {
			canvas.listeners[typ] = make([]*EventListener, 0)
			canvas.listeners[typ] = append(canvas.listeners[typ], listener)
			result = true
		}

	}

	return result

}

func (canvas *Canvas) RemoveEventListener(typ EventType, listener *EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := canvas.listeners[typ]

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
				canvas.Element.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				canvas.listeners[typ] = append(canvas.listeners[typ][:index], canvas.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := canvas.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				canvas.Element.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(canvas.listeners, typ)

			result = true

		}

	}

	return result

}

func (canvas *Canvas) GetContext() *Context {
	return canvas.Context
}

// Unlike a CanvasRenderingContext2D, a WebGL2RenderingContext does not follow
// the size of its Canvas Element, so the viewport is updated along with it.
func (canvas *Canvas) SetWidth(width uint) {

	canvas.Element.Value.Set("width", width)
	canvas.Width = width

	if canvas.Context != nil {

		canvas.Context.DrawingBufferWidth = canvas.Context.Value.Get("drawingBufferWidth").Int()
		canvas.Context.Viewport(0, 0, int(canvas.Width), int(canvas.Height))

	}

}

// Unlike a CanvasRenderingContext2D, a WebGL2RenderingContext does not follow
// the size of its Canvas Element, so the viewport is updated along with it.
func (canvas *Canvas) SetHeight(height uint) {

	canvas.Element.Value.Set("height", height)
	canvas.Height = height

	if canvas.Context != nil {

		canvas.Context.DrawingBufferHeight = canvas.Context.Value.Get("drawingBufferHeight").Int()
		canvas.Context.Viewport(0, 0, int(canvas.Width), int(canvas.Height))

	}

}
