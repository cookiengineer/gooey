//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strings"
import "syscall/js"

// Maybe field should be field.Label = ui.Label and field.Input = ui.Input?
// How to represent multiple choices with input type="radio" elements?

type field struct {
	Name  string               `json:"name"`
	Label interfaces.Component `json:"label"`
	Input interfaces.Component `json:"input"`
	Type  types.Input          `json:"type"`
	ctype string // component type
}

type Fieldset struct {
	Name      string                `json:"name"`
	Label     string                `json:"label"`
	Component *components.Component `json:"component"`
	fields    []*field
}

func NewFieldset(name string, label string) Fieldset {

	var fieldset Fieldset

	element   := bindings.Document.CreateElement("fieldset")
	component := components.NewComponent(element)

	element.SetAttribute("data-name", name)
	element.SetAttribute("id", toIdentifier(name))

	fieldset.Name      = strings.TrimSpace(name)
	fieldset.Label     = strings.TrimSpace(label)
	fieldset.Component = &component
	fieldset.fields    = make([]*field, 0)

	fieldset.Component.InitEvent("change")

	fieldset.Render()

	return fieldset

}

func ToFieldset(element *dom.Element) Fieldset {

	var fieldset Fieldset

	component := components.NewComponent(element)

	fieldset.Name      = strings.TrimSpace(element.GetAttribute("data-name"))
	fieldset.Component = &component
	fieldset.fields    = make([]*field, 0)

	fieldset.Parse()

	fieldset.Component.InitEvent("change")

	return fieldset

}

func (fieldset *Fieldset) AddField(name string, typ types.Input, Label interfaces.Component, Input interfaces.Component) {

	if name != "" {

		fieldset.fields = append(fieldset.fields, &field{
			Name:  name,
			Label: &label,
			Input: &input,
			Type:  typ,
		})

	}

}

func (fieldset *Fieldset) Disable() bool {

	var result bool

	if len(fieldset.fields) > 0 {

		for _, field := range fieldset.fields {
			field.Input.Disable()
		}

		result = true

	}

	return result

}

func (fieldset *Fieldset) Enable() bool {

	var result bool

	if len(fieldset.fields) > 0 {

		for _, field := range fieldset.fields {
			field.Input.Enable()
		}

		result = true

	}

	return result

}

func (fieldset *Fieldset) Parse() {

	if fieldset.Component.Element != nil {

		legend := fieldset.Component.Element.QuerySelector("legend")

		if legend != nil {
			fieldset.Label = strings.TrimSpace(legend.TextContent)
		}

		divs := fieldset.Component.Element.QuerySelectorAll("div")

		if len(divs) > 0 {

			for _, div := range divs {

				name     := div.GetAttribute("data-name")
				element1 := div.QuerySelector("label")
				element2 := div.QuerySelector("input, select, textarea")

				if element1 != nil && element2 != nil {

					if element2.TagName == "INPUT" {

						typ, ok := element2.Attributes["type"]

						if ok == true {

							if typ == "checkbox" {

								label := ui.ToLabel(element1)
								input := ui.ToCheckbox(element2)

								input.Component.AddEventListener("change", components.ToEventListener(func(event string, attributes map[string]string) {
									fieldset.Component.FireEventListeners("change", attributes)
								}, false))

								fieldset.fields = append(fieldset.fields, &field{
									Name:  name,
									Label: &label,
									Input: &input,
									Type:  input.Type,
									ctype: "ui.Checkbox",
								})

							} else if typ == "radio" {

								// TODO: Support ui.RadioGroup
								// label := ui.ToLabel(element1)
								// input := ui.ToRadioGroup(div.QuerySelectorAll("input[type=\"radio\"]"))

								// input.Component.AddEventListener("change", components.ToEventListener(func(event string, attributes map[string]string) {
								// 	fieldset.Component.FireEventListeners("change", attributes)
								// }, false))

								// fieldset.fields = append(fieldset.fields, &field{
								// 	Name:  name,
								// 	Label: &label,
								// 	Input: &input,
								// 	Type:  input.Type,
								//	ctype: "ui.RadioGroup",
								// })

							} else {

								label := ui.ToLabel(element1)
								input := ui.ToInput(element2)

								input.Component.AddEventListener("change", components.ToEventListener(func(event string, attributes map[string]string) {
									fieldset.Component.FireEventListeners("change", attributes)
								}, false))

								fieldset.fields = append(fieldset.fields, &field{
									Name:  name,
									Label: &label,
									Input: &input,
									Type:  input.Type,
									ctype: "ui.Input",
								})

							}

						}

					} else if element2.TagName == "SELECT" {

						label := ui.ToLabel(element1)
						input := ui.ToSelect(element2)

						input.Component.AddEventListener("change", components.ToEventListener(func(event string, attributes map[string]string) {
							fieldset.Component.FireEventListeners("change", attributes)
						}, false))

						fieldset.fields = append(fieldset.fields, &field{
							Name:  name,
							Label: &label,
							Input: &input,
							Type:  input.Type,
							ctype: "ui.Select",
						})

					} else if element2.TagName == "TEXTAREA" {

						label := ui.ToLabel(element1)
						input := ui.ToTextarea(element2)

						input.Component.AddEventListener("change", components.ToEventListener(func(event string, attributes map[string]string) {
							fieldset.Component.FireEventListeners("change", attributes)
						}, false))

						fieldset.fields = append(fieldset.fields, &field{
							Name:  name,
							Label: &label,
							Input: &input,
							Type:  input.Type,
							ctype: "ui.Textarea",
						})

					}

				}

			}

		}

	}

}

func (fieldset *Fieldset) RemoveField(name string) bool {

	var result bool

	var index int = -1

	for f := 0; f < len(fieldset.fields); f++ {

		if fieldset.fields[f].Name == name {
			index = f
			break
		}

	}

	if index != -1 {
		fieldset.fields = append(fieldset.fields[:index], fieldset.fields[index+1:]...)
		result = true
	}

	return result

}

func (fieldset *Fieldset) Render() *dom.Element {

	if fieldset.Component.Element != nil {

		elements := make([]*dom.Element, 0)

		if fieldset.Label != "" {

			legend := fieldset.Component.QuerySelector("legend")

			if legend != nil {
				legend.SetInnerHTML(fieldset.Label)
				elements = append(elements, legend)
			}

		}

		for _, field := range fieldset.fields {

			div := bindings.Document.CreateElement("div")

			label := field.Label.Render()
			input := field.Input.Render()

			div.ReplaceChildren([]*dom.Element{
				label,
				input,
			})

			elements = append(elements, div)

		}

		fieldset.ReplaceChildren(elements)

	}

	return fieldset.Component.Element

}

func (fieldset *Fieldset) String() string {

	html := "<fieldset>"

	if fieldset.Label != "" {
		html += "<legend>" + fieldset.Label + "</legend>"
	}

	for _, field := range fieldset.fields {

		html += "<div>"
		html += field.Label.String()
		html += field.Input.String()
		html += "</div>"

	}

	html += "</fieldset>"

	return html

}

func (fieldset *Fieldset) TypeOf(name string) types.Input {

	var result types.Input

	for _, field := range fieldset.fields {

		if field.Name == name {
			result = fieldset.Type
			break
		}

	}

	return result

}

func (fieldset *Fieldset) ValueOf(name string) js.Value {

	var result js.Value

	for _, field := range fieldset.fields {
	
		if field.Name == name {

			if field.ctype == "ui.Checkbox" {

				component, ok := field.Input.(ui.Checkbox)

				if ok == true {
					result = component.ToValue()
				}

			} else if field.ctype == "ui.Input" {

				component, ok := field.Input.(ui.Input)

				if ok == true {
					result = component.ToValue()
				}

			} else if field.ctype == "ui.RadioGroup" {

				// TODO: Support ui.RadioGroup
				// component, ok := field.Input.(ui.Radio)

				// if ok == true {
				// 	result = component.ToValue()
				// }

			} else if field.ctype == "ui.Select" {

				component, ok := field.Input.(ui.Select)

				if ok == true {
					result = component.ToValue()
				}

			} else if field.ctype == "ui.Textarea" {

				component, ok := field.Input.(ui.Textarea)

				if ok == true {
					result = component.ToValue()
				}

			}

			break

		}

	}

	return result

}

