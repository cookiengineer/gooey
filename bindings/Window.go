//go:build wasm

package bindings

import "github.com/cookiengineer/gooey/bindings/dom"
import "syscall/js"

var global_window *Window

func init() {
	global_window = GetWindow()
}

type Window struct {
	listeners   map[dom.EventType][]*dom.EventListener `json:"listeners"`
	Closed      bool                                   `json:"closed"`
	InnerWidth  uint                                   `json:"innerWidth"`
	InnerHeight uint                                   `json:"innerHeight"`
	OuterWidth  uint                                   `json:"outerWidth"`
	OuterHeight uint                                   `json:"outerHeight"`
	Screen      *Screen                                `json:"screen"`
	ScrollX     uint                                   `json:"scrollX"`
	ScrollY     uint                                   `json:"scrollY"`
	Value       *js.Value                              `json:"value"`
}

// Returns the global Window instance.
func GetWindow() *Window {

	if global_window != nil {

		return global_window

	} else {

		window_value := js.Global().Get("window")
		screen_value := window_value.Get("screen")
		screen_orientation := ScreenOrientation{}

		orientation := screen_value.Get("orientation")

		if !orientation.IsNull() && !orientation.IsUndefined() {

			screen_orientation = ScreenOrientation{
				Angle: uint(orientation.Get("angle").Int()),
				Type:  orientation.Get("type").String(),
			}

		}

		screen := Screen{
			listeners:   make(map[dom.EventType][]*dom.EventListener),
			Width:       uint(screen_value.Get("width").Int()),
			Height:      uint(screen_value.Get("height").Int()),
			AvailWidth:  uint(screen_value.Get("availWidth").Int()),
			AvailHeight: uint(screen_value.Get("availHeight").Int()),
			ColorDepth:  uint(screen_value.Get("colorDepth").Int()),
			PixelDepth:  uint(screen_value.Get("pixelDepth").Int()),
			IsExtended:  false,
			Orientation: &screen_orientation,
			Value:       &screen_value,
		}

		// Firefox doesn't expose this in Tracking Protection Mode
		screen_isextended := screen_value.Get("isExtended")

		if !screen_isextended.IsNull() && !screen_isextended.IsUndefined() {
			screen.IsExtended = screen_isextended.Bool()
		}

		window := Window{
			listeners:   make(map[dom.EventType][]*dom.EventListener),
			Closed:      window_value.Get("closed").Bool(),
			InnerWidth:  uint(window_value.Get("innerWidth").Int()),
			InnerHeight: uint(window_value.Get("innerHeight").Int()),
			OuterWidth:  uint(window_value.Get("outerWidth").Int()),
			OuterHeight: uint(window_value.Get("outerHeight").Int()),
			Screen:      &screen,
			ScrollX:     uint(window_value.Get("scrollX").Int()),
			ScrollY:     uint(window_value.Get("scrollY").Int()),
			Value:       &window_value,
		}

		window.Value.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) any {

			window.InnerWidth = uint(window.Value.Get("innerWidth").Int())
			window.InnerHeight = uint(window.Value.Get("innerHeight").Int())
			window.OuterWidth = uint(window.Value.Get("outerWidth").Int())
			window.OuterHeight = uint(window.Value.Get("outerHeight").Int())

			return nil

		}))

		window.Value.Call("addEventListener", "scroll", js.FuncOf(func(this js.Value, args []js.Value) any {

			window.ScrollX = uint(window.Value.Get("scrollX").Int())
			window.ScrollY = uint(window.Value.Get("scrollY").Int())

			return nil

		}))

		if !orientation.IsNull() && !orientation.IsUndefined() {

			window.Screen.Value.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) any {

				window.Screen.Orientation.Angle = uint(orientation.Get("angle").Int())
				window.Screen.Orientation.Type = orientation.Get("type").String()

				return nil

			}))

		}

		return &window

	}

}

// Adds an EventListener to the global Window. Only the "resize" and "scroll" EventType is supported.
func (win *Window) AddEventListener(typ dom.EventType, listener *dom.EventListener) bool {

	var result bool

	check := string(typ)

	if check == "resize" || check == "scroll" {

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

		win.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
		listener.Function = &wrapped_callback

		_, ok := win.listeners[typ]

		if ok == true {
			win.listeners[typ] = append(win.listeners[typ], listener)
			result = true
		} else {
			win.listeners[typ] = make([]*dom.EventListener, 0)
			win.listeners[typ] = append(win.listeners[typ], listener)
			result = true
		}

	}

	return result

}

// Removes an EventListener from the global Window. If listener is set to nil, all EventListeners are
// removed.
func (win *Window) RemoveEventListener(typ dom.EventType, listener *dom.EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := win.listeners[typ]

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
				win.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				win.listeners[typ] = append(win.listeners[typ][:index], win.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := win.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				win.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(win.listeners, typ)

			result = true

		}

	}

	return result

}

func (win *Window) Close() {

	win.Value.Call("close")
	win.Closed = win.Value.Get("closed").Bool()

}

func (win *Window) Confirm(message string) bool {

	var result bool

	tmp := win.Value.Call("confirm", js.ValueOf(message))

	if !tmp.IsNull() && !tmp.IsUndefined() {

		if tmp.Bool() == true {
			result = true
		}

	}

	return result

}

func (win *Window) Focus() {

	win.Value.Call("focus")

}

// Moves the Window by a specified delta x and y.
func (win *Window) MoveBy(delta_x int, delta_y int) {

	win.Value.Call("moveBy", js.ValueOf(delta_x), js.ValueOf(delta_y))

}

// Moves the Window to a specified position x and y.
func (win *Window) MoveTo(x uint, y uint) {

	win.Value.Call("moveTo", js.ValueOf(x), js.ValueOf(y))

}

// Resizes the Window by a specified delta x and y.
func (win *Window) ResizeBy(delta_x int, delta_y int) {

	win.Value.Call("resizeBy", js.ValueOf(delta_x), js.ValueOf(delta_y))
	win.InnerWidth = uint(win.Value.Get("innerWidth").Int())
	win.InnerHeight = uint(win.Value.Get("innerHeight").Int())
	win.OuterWidth = uint(win.Value.Get("outerWidth").Int())
	win.OuterHeight = uint(win.Value.Get("outerHeight").Int())

}

// Resizes the Window to a specified width and height.
func (win *Window) ResizeTo(width uint, height uint) {

	win.Value.Call("resizeTo", js.ValueOf(width), js.ValueOf(height))
	win.InnerWidth = uint(win.Value.Get("innerWidth").Int())
	win.InnerHeight = uint(win.Value.Get("innerHeight").Int())
	win.OuterWidth = uint(win.Value.Get("outerWidth").Int())
	win.OuterHeight = uint(win.Value.Get("outerHeight").Int())

}

// Scrolls the Window by a specified delta x and y.
func (win *Window) ScrollBy(delta_x int, delta_y int) {

	win.Value.Call("scrollBy", js.ValueOf(delta_x), js.ValueOf(delta_y))
	win.ScrollX = uint(win.Value.Get("scrollX").Int())
	win.ScrollY = uint(win.Value.Get("scrollY").Int())

}

// Scrolls the Window to a specified position x and y.
func (win *Window) ScrollTo(x uint, y uint) {

	win.Value.Call("scrollTo", js.ValueOf(x), js.ValueOf(y))
	win.ScrollX = uint(win.Value.Get("scrollX").Int())
	win.ScrollY = uint(win.Value.Get("scrollY").Int())

}
