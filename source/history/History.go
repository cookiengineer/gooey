//go:build wasm

package history

import "fmt"
import "syscall/js"

var History history

type history struct {
	listeners []*EventListener `json:"listeners"`
	stack     []*HistoryState  `json:"stack"`
	Length    uint             `json:"length"`
	State     *HistoryState    `json:"state"`
	Value     *js.Value        `json:"value"`
}

func init() {

	value := js.Global().Get("window").Get("history")

	History = history{
		listeners: make([]*EventListener, 0),
		Value:     &value,
	}

}

func (history *history) AddEventListener(listener EventListener) bool {

	wrapped_type     := js.ValueOf("popstate")
	wrapped_callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		if len(args) > 0 {

			event := args[0]

			if !event.IsNull() && !event.IsUndefined() {

				// TODO: How to get the state map now?

				wrapped_event := ToPopStateEvent(event)
				listener.Callback(wrapped_event)

			}

		}

		return nil

	})
	wrapped_capture := js.ValueOf(true)

	js.Global().Get("window").Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)

	listener.Function = &wrapped_callback
	history.listeners = append(history.listeners, &listener)

	return true

}

func (history *history) getListenerById(id uint) *EventListener {

	var result *EventListener

	for l := 0; l < len(history.listeners); l++ {

		if history.listeners[l].Id == id {
			result = history.listeners[l]
			break
		}

	}

	return result

}

func (history *history) RemoveEventListener(listener *EventListener) bool {

	var result bool = false

	if listener != nil {

		var index int = -1

		for l := 0; l < len(history.listeners); l++ {

			if history.listeners[l].Id == listener.Id {
				index = l
				break
			}

		}

		if index != -1 {

			listener := history.listeners[index]
			wrapped_type := js.ValueOf("popstate")
			wrapped_callback := *listener.Function
			wrapped_capture := js.ValueOf(true)
			js.Global().Get("window").Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			history.listeners = append(history.listeners[:index], history.listeners[index+1:]...)

			result = true

		}

	} else {

		for l := 0; l < len(history.listeners); l++ {

			listener := history.listeners[l]
			wrapped_type := js.ValueOf("popstate")
			wrapped_callback := *listener.Function
			wrapped_capture := js.ValueOf(true)
			js.Global().Get("window").Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

		}

		history.listeners = make([]*EventListener, 0)

		result = true

	}

	return result

}

func (history *history) Back() {

	if len(history.stack) > 0 {

		if history.State != nil {

			var index = -1

			for s := 0; s < len(history.stack); s++ {

				if history.stack[s] == history.State {
					index = s
					break
				}

			}

			if index != -1 && index > 0 {
				history.State = history.stack[index - 1]
			} else {
				history.State = nil
			}

		} else {

			history.State = nil

		}

		history.Value.Call("back")

	}


}

func (history *history) Forward() {

	if len(history.stack) > 0 {

		if history.State != nil {

			var index = -1

			for s := 0; s < len(history.stack); s++ {

				if history.stack[s] == history.State {
					index = s
					break
				}

			}

			if index != -1 && index < len(history.stack) - 1 {
				history.State = history.stack[index + 1]
			} else {
				history.State = history.stack[0]
			}

		} else {

			history.State = history.stack[0]

		}

		history.Value.Call("forward")

	}

}

func (history *history) Go(delta int) {

	if delta > 0 {

		wrapped_delta := js.ValueOf(delta)

		fmt.Println(wrapped_delta)

	} else if delta < 0 {

		// TODO: Negative direction / backwards

	} else if delta == 0 {

		// TODO? Nothing?

	}

}

func (history *history) PushState(statemap *map[string]any, title string, url string) bool {

	var result bool = false

	if title != "" && url != "" {

		var state HistoryState


		if statemap != nil {

			wrapped_statemap := js.ValueOf(*statemap)

			state = HistoryState{
				State: statemap,
				Title: title,
				URL:   url,
				value: &wrapped_statemap,
			}

		} else {

			wrapped_statemap := js.ValueOf(map[string]any{})

			state = HistoryState{
				State: nil,
				Title: title,
				URL:   url,
				value: &wrapped_statemap,
			}

		}

		history.State = &state
		history.stack = append(history.stack, &state)

		history.Value.Call("pushState", *state.value, js.ValueOf(state.Title), js.ValueOf(state.URL))

		result = true

	}

	return result

}

func (history *history) ReplaceState(statemap *map[string]any, title string, url string) bool {

	var result bool = false

	if title != "" && url != "" {

		var state HistoryState

		if statemap != nil {

			wrapped_statemap := js.ValueOf(*statemap)

			state = HistoryState{
				State: statemap,
				Title: title,
				URL:   url,
				value: &wrapped_statemap,
			}

		} else {

			wrapped_statemap := js.ValueOf(map[string]any{})

			state = HistoryState{
				State: nil,
				Title: title,
				URL:   url,
				value: &wrapped_statemap,
			}

		}

		history.State = &state
		history.stack[len(history.stack)-1] = &state

		history.Value.Call("replaceState", *state.value, js.ValueOf(state.Title), js.ValueOf(state.URL))

		result = true

	}

	return result

}
