//go:build wasm

package virtual

import bindings_dom "github.com/cookiengineer/gooey/bindings/dom"
import "strings"

// Element adapts a *bindings_dom.Element to the Element interface.
type Element struct {
	Element *bindings_dom.Element
}

func ToElement(element *bindings_dom.Element) *Element {

	if element == nil {
		return nil
	}

	return &Element{Element: element}

}

func ToElements(elements []*bindings_dom.Element) []*Element {

	result := make([]*Element, 0, len(elements))

	for _, element := range elements {
		if element != nil {
			result = append(result, &Element{Element: element})
		}
	}

	return result

}

func (element *Element) TagName() string {
	return strings.ToUpper(element.Element.TagName)
}

// Key returns the reconciliation key, resolved from
// data-key, then data-id, then id. Empty means positional.
func (element *Element) Key() string {

	attributes := make(map[string]string)

	element.Element.RefreshAttributes()

	for key, value := range element.Element.Attributes {
		attributes[key] = value
	}

	if value, ok := attributes["data-key"]; ok == true && value != "" {
		return value
	}

	if value, ok := attributes["data-id"]; ok == true && value != "" {
		return value
	}

	if value, ok := attributes["id"]; ok == true && value != "" {
		return value
	}

	return ""

}

// Attributes returns a snapshot of all attributes.
func (element *Element) Attributes() map[string]string {

	result := make(map[string]string)

	element.Element.RefreshAttributes()

	for key, value := range element.Element.Attributes {
		result[key] = value
	}

	return result

}

func (element *Element) SetAttribute(name string, value string) {
	element.Element.SetAttribute(name, value)
}

func (element *Element) RemoveAttribute(name string) {
	element.Element.RemoveAttribute(name)
}

func (element *Element) Children() []*Element {

	children := element.Element.Children()

	if len(children) == 0 {
		return make([]*Element, 0)
	}

	return ToElements(children)

}

func (element *Element) GetTextContent() string {
	return element.Element.GetTextContent()
}

func (element *Element) SetTextContent(value string) {
	element.Element.SetTextContent(value)
}

// InsertBefore inserts child before reference. A nil reference appends.
func (element *Element) InsertBefore(child *Element, reference *Element) {

	if child != nil {

		if reference != nil {
			element.Element.InsertBefore(child.Element, reference.Element)
		} else {
			element.Element.InsertBefore(child.Element, nil)
		}

	}

}

// RemoveChild removes child from this element.
func (element *Element) RemoveChild(child *Element) {

	if child != nil {
		element.Element.RemoveChild(child.Element)
	}

}

// ReplaceChild replaces old with replacement.
func (element *Element) ReplaceChild(old_element *Element, new_element *Element) {

	if old_element != nil && new_element != nil {
		element.Element.ReplaceChild(old_element.Element, new_element.Element)
	}

}

// Focused reports whether this element currently holds focus.
func (element *Element) Focused() bool {
	return element.Element.Focused()
}

// IsSameElement reports whether other wraps the same underlying element.
func (element *Element) IsSameElement(other *Element) bool {

	if other != nil {
		return element.Element.IsSameElement(other.Element)
	} else {
		return false
	}

}

