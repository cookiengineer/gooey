package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "syscall/js"

// Maybe field should be field.Label = ui.Label and field.Input = ui.Input?
// How to represent multiple choices with input type="radio" elements?

type field struct {
	Name  string               `json:"name"`
	Label interfaces.Component `json:"label"`
	Input interfaces.Component `json:"input"`
	Type  types.Input          `json:"type"`
}

type Fieldset struct {
	Name   string   `json:"name"`
	Label  string   `json:"label"`
	fields []*field
}

func NewFieldset(name string, label string) Fieldset {

	var fieldset Fieldset

	element   := bindings.Document.CreateElement("fieldset")
	component := components.NewComponent(element)

	element.SetAttribute("data-name", name)
	element.SetAttribute("id", toIdentifier(name))

	return fieldset

}

func (fieldset *Fieldset) Disable() bool {

	var result bool

	// TODO: Disable all elements

	return result

}

func (fieldset *Fieldset) Enable() bool {

	var result bool

	// TODO: Enable all elements

	return result

}

func (fieldset *Fieldset) ValueOf(name string) js.Value {

	var result js.Value

	for f := 0; f < len(fieldset.fields); f++ {

		field := fieldset.fields[f]

		if field.Name == name {
			// TODO: Get value?
		}

	}


	return result

}

func (fieldset *Fieldset) Parse() {

	if fieldset.Component.Element != nil {

		divs := fieldset.Component.Element.QuerySelectorAll("div")

		if len(divs) > 0 {

			for _, div := range divs {

				element1 := div.QuerySelector("label")
				element2 := div.QuerySelector("input, select, textarea")

				if element1 != nil && element2 != nil {

					// TODO: checkbox
					// TODO: radios

					if element2.TagName == "INPUT" {

						typ, ok := element2.Attributes["type"]

						if ok == true {

							if typ == "checkbox" {

								label := ui.ToLabel(element1)
								input := ui.ToCheckbox(element2)

								// TODO: Add field

							} else if typ == "radio" {

								label := ui.ToLabel(element1)
								input := ui.ToChoices(div.QuerySelectorAll("input[type=\"radio\"]"))

								// TODO: Add field

							} else {

								label := ui.ToLabel(element1)
								input := ui.ToInput(element2)

								// TODO: Add field

							}

						}

					} else if element2.TagName == "SELECT" {

						label := ui.ToLabel(element1)
						input := ui.ToSelect(element2)

						// TODO: Add field

					} else if element2.TagName == "TEXTAREA" {

						label := ui.ToLabel(element1)
						input := ui.ToTextarea(element2)

						// TODO: Add field

					}

				}

				// TODO: label

			}

		}

	}

	// TODO: Parse Fieldset <div> lines into elements

}

func (fieldset *Fieldset) Render() *dom.Element {

	// TOOD: Render() method

	return fieldset.Component.Element

}

func (fieldset *Fieldset) String() string {

	html := "<fieldset>"

	for _, field := range fieldset.fields {

		html += "<div>"
		html += field.Label.String()
		html += field.Input.String()
		html += "</div>"

	}

	html += "</fieldset>"

	return html

}

func (fieldset *Fieldset) AddField(name string, typ types.Input, Label interfaces.Component, Input interfaces.Component) {

}

func (fieldset *Fieldset) TypeOf(name string) types.Input {


}

// TODO: How to use Set/Get for type specific methods?
// TODO: Maybe generic method makes more sense?
// TODO: Fieldset.Set(name string, value bool || int || uint || string etc)

