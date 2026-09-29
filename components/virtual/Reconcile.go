//go:build wasm

package virtual

import "strings"

// Reconcile patches the children of parent so that they match desired.
//
// Existing children are reused by key (data-key, data-id or id), then by
// positional tag match. Unmatched children are removed and missing children
// are inserted. Elements that are already in the correct position are left
// untouched, preserving focus, selection and scroll state.
func Reconcile(parent *Element, desired []*Element) {

	current := parent.Children()
	keyed := make(map[string]*Element)

	for _, element := range current {
		if key := element.Key(); key != "" {
			keyed[key] = element
		}
	}

	used := make([]*Element, 0, len(current))
	cursor := 0

	for _, want := range desired {

		key := want.Key()

		var match *Element

		if key != "" {

			candidate, ok := keyed[key]

			if ok == true && contains_element(used, candidate) == false && strings.EqualFold(candidate.TagName(), want.TagName()) == true {
				match = candidate
			}

		} else {
			match = find_positional_element(current, used, cursor, want.TagName())
		}

		if match != nil {

			Patch(match, want)
			Move(parent, match, cursor)
			used = append(used, match)

		} else {

			children   := parent.Children()

			var element_at *Element = nil

			if cursor >= 0 && cursor < len(children) {
				element_at = children[cursor]
			}

			parent.InsertBefore(want, element_at)
			used = append(used, want)

		}

		cursor++

	}

	for _, element := range current {

		if contains_element(used, element) == false {
			parent.RemoveChild(element)
		}

	}

}

