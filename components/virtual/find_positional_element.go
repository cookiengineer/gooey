//go:build wasm

package virtual

import "strings"

func find_positional_element(current []*Element, used []*Element, cursor int, tag string) *Element {

	if cursor < len(current) {

		candidate := current[cursor]

		if contains_element(used, candidate) == false && strings.EqualFold(candidate.TagName(), tag) == true {
			return candidate
		}

	}

	for _, candidate := range current {

		if contains_element(used, candidate) == true {
			continue
		}

		if strings.EqualFold(candidate.TagName(), tag) == true {
			return candidate
		}

	}

	return nil

}
