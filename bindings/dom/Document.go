//go:build wasm

package dom

import "syscall/js"

var Document document

func init() {

	doc := js.Global().Get("document")
	html := ToElement(doc.Call("querySelector", "html"))
	head := ToElement(doc.Get("head"))
	body := ToElement(doc.Get("body"))

	Document = document{
		listeners: make(map[EventType][]*EventListener),
		Element:   html,
		Head:      head,
		Body:      body,
		Value:     &doc,
	}

}

type document struct {
	listeners map[EventType][]*EventListener `json:"listeners"`
	Element   *Element                       `json:"element"`
	Head      *Element                       `json:"head"`
	Body      *Element                       `json:"body"`
	Value     *js.Value                      `json:"value"`
}

func (doc *document) AddEventListener(typ EventType, listener *EventListener) bool {

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

		doc.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)
		listener.Function = &wrapped_callback

		_, ok := doc.listeners[typ]

		if ok == true {
			doc.listeners[typ] = append(doc.listeners[typ], listener)
			result = true
		} else {
			doc.listeners[typ] = make([]*EventListener, 0)
			doc.listeners[typ] = append(doc.listeners[typ], listener)
			result = true
		}

	}

	return result

}

func (doc *document) CreateElement(tagname string) *Element {

	var result *Element = nil

	value := doc.Value.Call("createElement", js.ValueOf(tagname))

	if !value.IsNull() && !value.IsUndefined() {
		result = ToElement(value)
	}

	return result

}

func (doc *document) CreateElementNS(url string, tagname string) *Element {

	var result *Element = nil

	value := doc.Value.Call("createElementNS", js.ValueOf(url), js.ValueOf(tagname))

	if !value.IsNull() && !value.IsUndefined() {
		result = ToElement(value)
	}

	return result

}


func (doc *document) QuerySelector(query string) *Element {

	var result *Element = nil

	value := doc.Value.Call("querySelector", query)

	if !value.IsNull() && !value.IsUndefined() {
		result = ToElement(value)
	}

	return result

}

func (doc *document) QuerySelectorAll(query string) []*Element {

	var result []*Element

	values := doc.Value.Call("querySelectorAll", query)

	for v := 0; v < values.Length(); v++ {

		value := values.Index(v)

		if !value.IsNull() && !value.IsUndefined() {
			result = append(result, ToElement(value))
		}

	}

	return result

}

func (doc *document) RemoveEventListener(typ EventType, listener *EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := doc.listeners[typ]

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
				doc.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				doc.listeners[typ] = append(doc.listeners[typ][:index], doc.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := doc.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				doc.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(doc.listeners, typ)

			result = true

		}

	}

	return result

}
