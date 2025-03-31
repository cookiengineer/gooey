//go:build wasm

package dom

import "strings"
import "syscall/js"

type Element struct {
	listeners   map[EventType][]*EventListener `json:"listeners"`
	Id          string                         `json:"id"`
	ClassName   string                         `json:"className"`
	Attributes  map[string]string              `json:"attributes"`
	TagName     string                         `json:"tagName"`
	TextContent string                         `json:"textContent"`
	InnerHTML   string                         `json:"innerHTML"`
	Value       *js.Value                      `json:"value"`
}

func ToElement(value js.Value) Element {

	var element Element

	element.listeners = make(map[EventType][]*EventListener)
	element.Id = value.Get("id").String()
	element.ClassName = value.Get("className").String()
	element.TagName = value.Get("tagName").String()
	element.TextContent = value.Get("textContent").String()
	element.InnerHTML = value.Get("innerHTML").String()
	element.Value = &value

	element.Attributes = make(map[string]string)
	element.RefreshAttributes()

	return element

}

func (element *Element) AddEventListener(typ EventType, listener EventListener) bool {

	var result bool

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

	element.Value.Call("addEventListener", wrapped_type, wrapped_callback, wrapped_capture)

	_, ok := element.listeners[typ]

	if ok == true {
		element.listeners[typ] = append(element.listeners[typ], &listener)
		result = true
	} else {
		element.listeners[typ] = make([]*EventListener, 0)
		element.listeners[typ] = append(element.listeners[typ], &listener)
		result = true
	}

	return result

}

func (element *Element) Append(child *Element) {

	if element.Value != nil && child != nil && child.Value != nil {
		element.Value.Call("append", *child.Value)
	}

}

func (element *Element) Prepend(child *Element) {

	if element.Value != nil && child != nil && child.Value != nil {
		element.Value.Call("prepend", *child.Value)
	}

}

func (element *Element) Remove() {
	element.Value.Call("remove")
}

func (element *Element) RemoveEventListener(typ EventType, listener *EventListener) bool {

	var result bool

	if listener != nil {

		listeners, ok := element.listeners[typ]

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
				element.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

				element.listeners[typ] = append(element.listeners[typ][:index], element.listeners[typ][index+1:]...)

				result = true

			}

		}

	} else {

		listeners, ok := element.listeners[typ]

		if ok == true {

			for l := 0; l < len(listeners); l++ {

				listener := listeners[l]
				wrapped_type := js.ValueOf(string(typ))
				wrapped_callback := *listener.Function
				wrapped_capture := js.ValueOf(true)
				element.Value.Call("removeEventListener", wrapped_type, wrapped_callback, wrapped_capture)

			}

			delete(element.listeners, typ)

			result = true

		}

	}

	return result

}

func (element *Element) RefreshAttributes() {

	attributes := element.Value.Call("getAttributeNames")

	if !attributes.IsNull() && !attributes.IsUndefined() {

		for key, _ := range element.Attributes {
			delete(element.Attributes, key)
		}

		for a := 0; a < attributes.Length(); a++ {

			name := attributes.Index(a)
			value := element.Value.Call("getAttribute", name)

			if !value.IsNull() {
				element.Attributes[name.String()] = value.String()
			}

		}

	}

}

func (element *Element) GetAttribute(name string) string {

	var value string

	check := validateXMLName(name)

	if check == nil {

		tmp := element.Value.Call("getAttribute", name)

		if !tmp.IsNull() {
			element.Attributes[name] = tmp.String()
			value = element.Attributes[name]
		}

	}

	return value

}

func (element *Element) SetAttribute(name string, value string) bool {

	var result bool

	check := validateXMLName(name)

	if check == nil {
		element.Attributes[name] = value
		element.Value.Call("setAttribute", name, value)
	} else {
		result = false
	}

	return result

}

func (element *Element) HasAttribute(name string) bool {

	var result bool

	check := validateXMLName(name)

	if check == nil {

		tmp := element.Value.Call("hasAttribute", name)

		if tmp.Truthy() {
			result = true
		}

	}

	return result

}

func (element *Element) RemoveAttribute(name string) bool {

	var result bool

	check := validateXMLName(name)

	if check == nil {

		tmp := element.Value.Call("removeAttribute", name)

		if tmp.Truthy() {
			result = true
		}

	}

	return result

}

func (element *Element) GetBoundingClientRect() *Rect {

	var result *Rect = nil

	value := element.Value.Call("getBoundingClientRect")

	if !value.IsNull() && !value.IsUndefined() {
		rect := ToRect(value)
		result = &rect
	}

	return result

}

func (element *Element) InsertAdjacentElement(position string, other *Element) {

	if position == "beforebegin" {
		element.Value.Call("insertAdjacentElement", js.ValueOf(position), other.Value)
	} else if position == "afterbegin" {
		element.Value.Call("insertAdjacentElement", js.ValueOf(position), other.Value)
	} else if position == "beforeend" {
		element.Value.Call("insertAdjacentElement", js.ValueOf(position), other.Value)
	} else if position == "afterend" {
		element.Value.Call("insertAdjacentElement", js.ValueOf(position), other.Value)
	}

}

func (element *Element) InsertAdjacentHTML(position string, value string) {

	if position == "beforebegin" {
		element.Value.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(value))
	} else if position == "afterbegin" {
		element.Value.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(value))
	} else if position == "beforeend" {
		element.Value.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(value))
	} else if position == "afterend" {
		element.Value.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(value))
	}

}

func (element *Element) InsertAdjacentText(position string, value string) {

	if position == "beforebegin" {
		element.Value.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(value))
	} else if position == "afterbegin" {
		element.Value.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(value))
	} else if position == "beforeend" {
		element.Value.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(value))
	} else if position == "afterend" {
		element.Value.Call("insertAdjacentText", js.ValueOf(position), js.ValueOf(value))
	}

}

func (element *Element) ParentNode() *Element {

	var result *Element = nil

	value := element.Value.Get("parentNode")

	if !value.IsNull() && !value.IsUndefined() {
		parent := ToElement(value)
		result = &parent
	}

	return result

}

func (element *Element) QueryParent(search string) *Element {

	var result *Element = nil

	value := element.Value.Get("parentNode")
	tagname := value.Get("tagName")

	for !tagname.IsNull() && !tagname.IsUndefined() && tagname.String() != "BODY" {

		tmp := strings.ToLower(tagname.String())

		if tmp == search {
			parent := ToElement(value)
			result = &parent
			break
		} else {
			value = value.Get("parentNode")
			tagname = value.Get("tagName")
		}

	}

	return result

}

func (element *Element) QuerySelector(query string) *Element {

	var result *Element = nil

	value := element.Value.Call("querySelector", query)

	if !value.IsNull() && !value.IsUndefined() {
		child := ToElement(value)
		result = &child
	}

	return result

}

func (element *Element) QuerySelectorAll(query string) []*Element {

	var result []*Element

	values := element.Value.Call("querySelectorAll", query)

	for v := 0; v < values.Length(); v++ {

		value := values.Index(v)

		if !value.IsNull() && !value.IsUndefined() {

			child := ToElement(value)
			result = append(result, &child)

		}

	}

	return result

}

func (element *Element) ReplaceChildren(children []*Element) {

	values := make([]any, len(children))

	for c, child := range children {
		values[c] = *child.Value
	}

	element.Value.Call("replaceChildren", values...)

}

func (element *Element) SetClassName(value string) bool {

	var result bool

	value = strings.TrimSpace(value)

	element.Value.Set("className", value)
	element.ClassName = element.Value.Get("className").String()

	if element.ClassName == value {
		result = true
	}

	return result

}

func (element *Element) SetInnerHTML(value string) bool {

	element.Value.Set("innerHTML", value)
	element.InnerHTML = element.Value.Get("innerHTML").String()

	return true

}
