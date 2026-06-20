//go:build wasm

package websockets

import "github.com/cookiengineer/gooey/bindings/quirks"
import "strings"
import "syscall/js"

// TODO: Update bufferedAmount on every data event and message event

type WebSocket struct {
	listeners      map[EventType][]*EventListener `json:"listeners"`
	URL            string                         `json:"url"`
	BinaryType     string                         `json:"binaryType"`
	BufferedAmount uint                           `json:"bufferedAmount"`
	Extensions     []string                       `json:"extensions"`
	Protocol       string                         `json:"protocol"`
	ReadyState     ReadyState                     `json:"readyState"`
	Value          *js.Value                      `json:"value"`
}

// Returns new WebSocket instance.
func NewWebSocket(url string) *WebSocket {

	if strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://") {

		ws := js.Global().Get("WebSocket").New(url)
		ws.Set("binaryType", "arraybuffer")

		websocket := WebSocket{
			listeners:      make(map[EventType][]*EventListener, 0),
			URL:            url,
			BinaryType:     "arraybuffer",
			BufferedAmount: 0,
			Extensions:     make([]string, 0),
			Protocol:       "",
			ReadyState:     ReadyStateConnecting,
			Value:          &ws,
		}

		websocket.Value.Call("addEventListener", js.ValueOf("open"), js.FuncOf(func(this js.Value, args []js.Value) any {

			wrapped_extensions := websocket.Value.Get("extensions")
			wrapped_protocol := websocket.Value.Get("protocol")
			wrapped_readystate := websocket.Value.Get("readyState")
			wrapped_url := websocket.Value.Get("url")

			if !wrapped_extensions.IsNull() && !wrapped_extensions.IsUndefined() {

				extensions := make([]string, 0)

				tmp1 := strings.Split(strings.TrimSpace(wrapped_extensions.String()), ",")

				for _, val := range tmp1 {

					tmp2 := strings.TrimSpace(val)

					if tmp2 != "" {
						extensions = append(extensions, tmp2)
					}

				}

				websocket.Extensions = extensions

			}

			if !wrapped_protocol.IsNull() && !wrapped_protocol.IsUndefined() {

				protocol := strings.TrimSpace(wrapped_protocol.String())

				if protocol != "" {
					websocket.Protocol = protocol
				}

			}

			if !wrapped_readystate.IsNull() && !wrapped_readystate.IsUndefined() {
				websocket.ReadyState = ReadyState(int(wrapped_readystate.Int()))
			}

			if !wrapped_url.IsNull() && !wrapped_url.IsUndefined() {

				url := strings.TrimSpace(wrapped_url.String())

				if strings.HasPrefix(url, "ws://") || strings.HasPrefix(url, "wss://") {
					websocket.URL = url
				}

			}

			return nil

		}))

		return &websocket

	}

	return nil

}

// Adds an EventListener to the WebSocket.
func (websocket *WebSocket) AddEventListener(typ EventType, listener *EventListener) bool {

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

		websocket.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
		listener.Function = &wrapped_callback

		_, ok := websocket.listeners[typ]

		if ok == true {
			websocket.listeners[typ] = append(websocket.listeners[typ], listener)
			result = true
		} else {
			websocket.listeners[typ] = make([]*EventListener, 0)
			websocket.listeners[typ] = append(websocket.listeners[typ], listener)
			result = true
		}

	}

	return result

}

func (websocket *WebSocket) RemoveEventListener(typ EventType, listener *EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := websocket.listeners[typ]

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
				websocket.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				websocket.listeners[typ] = append(websocket.listeners[typ][:index], websocket.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := websocket.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				websocket.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(websocket.listeners, typ)

			result = true

		}

	}

	return result

}

func (websocket *WebSocket) Close(status Status, reason string) error {

	err := quirks.GoTryCatch(func() {

		wrapped_status := js.ValueOf(int(status))
		wrapped_reason := js.ValueOf(strings.TrimSpace(reason))

		websocket.Value.Call("close", wrapped_status, wrapped_reason)

	})

	if err == nil {
		return nil
	} else {
		return err
	}

}

func (websocket *WebSocket) Send(data []byte) error {

	err := quirks.GoTryCatch(func() {

		wrapped_data := js.Global().Get("Uint8Array").New(len(data))
		js.CopyBytesToJS(wrapped_data, data)

		websocket.Value.Call("send", wrapped_data)

	})

	if err == nil {
		return nil
	} else {
		return err
	}

}
