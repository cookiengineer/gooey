//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/types"
import "strconv"
import "strings"
import "syscall/js"

type Range struct {
	Label     string                `json:"label"`
	Type      types.Input           `json:"type"`
	Value     int                   `json:"value"`
	Min       int                   `json:"min"`
	Max       int                   `json:"max"`
	Step      int                   `json:"step"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewRange(label string, step int, cur_value int, min_value int, max_value int) Range {

	var input Range

	element   := dom.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", "range")

	if cur_value >= min_value && cur_value <= max_value {

		input.Min   = min_value
		input.Max   = max_value
		input.Value = cur_value

	} else if cur_value >= 0 && cur_value <= 100 {

		input.Min   = 0
		input.Max   = 100
		input.Value = cur_value

	}

	if step > 0 {
		input.Step = step
	} else {
		input.Step = 1
	}

	input.Component = &component
	input.Label     = label
	input.Type      = types.InputRange

	if input.Value > input.Max {
		input.Max = input.Value
	}

	input.Mount()
	input.Render()

	return input

}

func ToRange(element *dom.Element) *Range {

	var input Range

	tmp := element.Value.Get("value")

	if !tmp.IsNull() && !tmp.IsUndefined() {

		number, err := strconv.ParseInt(tmp.String(), 10, 64)

		if err == nil {
			input.Value = int(number)
		} else {
			input.Value = 0
		}

	} else {
		input.Value = 0
	}

	max_str  := strings.TrimSpace(element.GetAttribute("max"))
	min_str := strings.TrimSpace(element.GetAttribute("min"))

	if min_str != "" && max_str != "" {

		min_number, err1 := strconv.ParseInt(min_str, 10, 64)
		max_number, err2 := strconv.ParseInt(max_str, 10, 64)

		if err1 == nil && err2 == nil && min_number < max_number {
			input.Min = int(min_number)
			input.Max = int(max_number)
		} else {
			input.Min = 0
			input.Max = 100
		}

	} else {
		input.Min = 0
		input.Max = 100
	}

	if input.Value > input.Max {
		input.Max = input.Value
	}

	step_str := strings.TrimSpace(element.GetAttribute("step"))

	if step_str != "" {

		step_number, err1 := strconv.ParseInt(step_str, 10, 64)

		if err1 == nil && step_number > 0 {
			input.Step = int(step_number)
		} else {
			input.Step = 1
		}

	}

	component := components.NewComponent(element)

	input.Component = &component
	input.Label     = strings.TrimSpace(element.GetAttribute("placeholder"))
	input.Type      = types.Input(element.GetAttribute("type"))
	input.Disabled  = element.HasAttribute("disabled")

	input.Mount()

	return &input

}

func (input *Range) Disable() bool {

	input.Disabled = true
	input.Render()

	return true

}

func (input *Range) Enable() bool {

	input.Disabled = false
	input.Render()

	return true

}

func (input *Range) Mount() bool {

	input.Component.InitEvent("change-value")

	if input.Component.Element != nil {

		input.Component.Element.AddEventListener("change", dom.ToEventListener(func(_ *dom.Event) {

			value := input.Component.Element.Value.Get("value").String()
			number, err := strconv.ParseInt(value, 10, 64)

			if err == nil {

				input.Value = int(number)

				input.Component.FireEventListeners("change-value", map[string]any{
					"value": input.Value,
				})

			}

		}))

		return true

	} else {
		return false
	}

}

func (input *Range) Query(query string) interfaces.Component {
	return nil
}

func (input *Range) Render() *dom.Element {

	if input.Component.Element != nil {

		if input.Label != "" {
			input.Component.Element.SetAttribute("placeholder", input.Label)
		} else {
			input.Component.Element.RemoveAttribute("placeholder")
		}

		input.Component.Element.SetAttribute("type", input.Type.String())

		if input.Value >= 0 {
			input.Component.Element.Value.Set("value", input.Value)
		} else {
			input.Component.Element.Value.Set("value", "")
		}

		if input.Disabled == true {
			input.Component.Element.SetAttribute("disabled", "")
		} else {
			input.Component.Element.RemoveAttribute("disabled")
		}

		if input.Min >= 0 {
			input.Component.Element.SetAttribute("min", strconv.Itoa(input.Min))
		}

		if input.Max > 0 {
			input.Component.Element.SetAttribute("max", strconv.Itoa(input.Max))
		}

		if input.Step > 0 {
			input.Component.Element.SetAttribute("step", strconv.Itoa(input.Step))
		}

	}

	return input.Component.Element

}

func (input *Range) Reset() bool {

	input.Value = input.Min
	input.Render()

	return true

}

func (input *Range) SetChildren(children []interfaces.Component) bool {
	return false
}

func (input *Range) String() string {

	html := "<input type=\"" + input.Type.String() + "\""

	if input.Label != "" {
		html += " placeholder=\"" + input.Label + "\""
	}

	if input.Value != 0 {
		html += " value=\"" + strconv.Itoa(input.Value) + "\""
	}

	if input.Min >= 0 {
		html += " min=\"" + strconv.Itoa(input.Min) + "\""
	}

	if input.Max >= 0 {
		html += " max=\"" + strconv.Itoa(input.Max) + "\""
	}

	if input.Step >= 0 {
		html += " step=\"" + strconv.Itoa(input.Step) + "\""
	}

	if input.Disabled == true {
		html += " disabled"
	}

	html += "/>"

	return html

}

func (input *Range) ToValue() js.Value {

	var result js.Value

	if input.Component.Element != nil {

		tmp := input.Component.Element.Value.Get("value")

		if !tmp.IsNull() && !tmp.IsUndefined() {

			num, err := strconv.ParseInt(tmp.String(), 10, 64)

			if err == nil {
				result = js.ValueOf(num)
			}

		}

	}

	return result

}

func (input *Range) Unmount() bool {

	if input.Component.Element != nil {
		input.Component.Element.RemoveEventListener("change", nil)
	}

	return true

}
