//go:build wasm

package virtual

func contains_element(elements []*Element, element *Element) bool {

	if element != nil {

		for _, candidate := range elements {

			if candidate.IsSameElement(element) == true {
				return true
			}

		}

	}

	return false

}
