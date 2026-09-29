//go:build wasm

package virtual

// Patch copies attributes and, for leaf elements, the text from new_element into old_element.
// Elements with children are reconciled recursively.
func Patch(old_element *Element, new_element *Element) {

	for name, value := range new_element.Attributes() {

		if old_element.Attributes()[name] != value {
			old_element.SetAttribute(name, value)
		}

	}

	for name := range old_element.Attributes() {

		if _, ok := new_element.Attributes()[name]; ok == false {
			old_element.RemoveAttribute(name)
		}

	}

	old_children := old_element.Children()
	new_children := new_element.Children()

	if len(new_children) == 0 && len(old_children) == 0 {

		// Text only for leaves; never clobber a focused element.
		if old_element.Focused() == false && old_element.GetTextContent() != new_element.GetTextContent() {
			old_element.SetTextContent(new_element.GetTextContent())
		}

	} else {
		Reconcile(old_element, new_children)
	}

}

