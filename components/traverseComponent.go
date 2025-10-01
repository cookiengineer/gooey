package components

import "github.com/cookiengineer/gooey/interfaces"
import "strings"

func traverseComponent(document *Document, component *Component) {

	nested_content  := make([]interfaces.Component, 0)
	nested_children := component.Element.Children()

	for _, nested_element := range nested_children {

		wrapper, ok := document.Registry[strings.ToLower(nested_element.TagName)]

		if ok == true {

			nested_component := wrapper(nested_element)

			if nested_component != nil {
				nested_content = append(nested_content, nested_component)
			}

		} else {

			nested_component := NewComponent(nested_element)
			traverseComponent(document, &nested_component)
			nested_content = append(nested_content, &nested_component)

		}

	}

	component.SetContent(nested_content)

}
