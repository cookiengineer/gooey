//go:build wasm

package virtual

func Move(parent *Element, element *Element, index int) {

	children := parent.Children()
	current := -1

	for i, child := range children {

		if child.IsSameElement(element) == true {
			current = i
			break
		}

	}

	if current == -1 || current == index {
		return
	}

	var reference *Element

	if current < index {

		if index+1 < len(children) {
			reference = children[index+1]
		}

	} else {

		if index < len(children) {
			reference = children[index]
		}

	}

	if reference != nil && reference.IsSameElement(element) == true {
		reference = nil
	}

	parent.InsertBefore(element, reference)

}

