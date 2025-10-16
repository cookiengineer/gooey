//go:build wasm

package ui

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/interfaces"
import "github.com/cookiengineer/gooey/components/types"
import "strconv"
import "strings"
import "syscall/js"

type Number struct {
	Label     string                `json:"label"`
	Type      types.Input           `json:"type"`
	Value     int                   `json:"value"`
	Min       int                   `json:"min"`
	Max       int                   `json:"max"`
	Step      int                   `json:"step"`
	Disabled  bool                  `json:"disabled"`
	Component *components.Component `json:"component"`
}

func NewNumber(label string, min_value int, max_value int, step int, cur_value int) Number {

	var input Number

	element := dom.Document.CreateElement("input")
	component := components.NewComponent(element)

	element.SetAttribute("type", "number")

	input.Component = &component
	input.Label = label
	input.Type = types.InputNumber
	input.Value = 0
	input.Min = 0
	input.Max = 100
	input.Step = 1
	input.Disabled = false

	if min_value <= max_value - step && step != 0 {

		input.Min = min_value
		input.Max = max_value
		input.Step = step

		if (cur_value - min_value) % step == 0 {
			input.Value = cur_value
		} else {
			input.Value = min_value
		}

	}

	return input

}

func ToNumber(element *dom.Element) *Number {

	var input Number

	component := components.NewComponent(element)

	input.Component = &component
	input.Label = strings.TrimSpace(element.GetAttribute("placeholder"))
	input.Type = types.Input(element.GetAttribute("type"))
	input.Value = 0
	input.Min = 0
	input.Max = 100
	input.Step = 1
	input.Disabled = element.HasAttribute("disabled")

	min_str := strings.TrimSpace(element.GetAttribute("min"))
	max_str := strings.TrimSpace(element.GetAttribute("max"))
	step_str := strings.TrimSpace(element.GetAttribute("step"))

	if min_str != "" && max_str != "" && step_str != "" {

		tmp1, err1 := strconv.ParseInt(step_str, 10, 64)
		tmp2, err2 := strconv.ParseInt(min_str, 10, 64)
		tmp3, err3 := strconv.ParseInt(max_str, 10, 64)

		if err1 == nil && err2 == nil && err3 == nil {

			step := int(tmp1)
			min_value := int(tmp2)
			max_value := int(tmp3)

			if min_value <= max_value - step && step != 0 {

				input.Min = min_value
				input.Max = max_value
				input.Step = step

			}

		}

	}

	value := element.Value.Get("value")

	if !value.IsNull() && !value.IsUndefined() {

		tmp1, err1 := strconv.ParseInt(value.String(), 10, 64)

		if err1 == nil {

			cur_value := int(tmp1)

			if (cur_value - input.Min) % input.Step == 0 {
				input.Value = cur_value
			} else {
				input.Value = input.Min
			}

		} else {
			input.Value = input.Min
		}

	}

	return &input

}

func (input *Number) Disable() bool {

	input.Disabled = true
	input.Render()

	return true

}

func (input *Number) Enable() bool {

	input.Disabled = false
	input.Render()

	return true

}

func (input *Number) Mount() bool {

	if input.Component != nil {
		input.Component.InitEvent("change-value")
	}

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

func (input *Number) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if input.Component.Element != nil {

			if utils.MatchesQuery(input.Component.Element, selectors[0]) == true {
				return input
			}

		}

	}

	return nil

}

func (input *Number) Render() *dom.Element {

	if input.Component.Element != nil {

		if input.Label != "" {
			input.Component.Element.SetAttribute("placeholder", input.Label)
		} else {
			input.Component.Element.RemoveAttribute("placeholder")
		}

		input.Component.Element.SetAttribute("type", input.Type.String())

		if input.Value != 0 {
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

func (input *Number) Reset() bool {

	input.Value = input.Min
	input.Render()

	return true

}

func (input *Number) SetLabel(label string) {

	input.Label = strings.TrimSpace(label)
	input.Render()

}

func (input *Number) SetMinMaxStep(min_value int, max_value int, step int) {

	if min_value <= max_value - step {

		input.Min = min_value
		input.Max = max_value
		input.Step = step

		input.Render()

	}

}

func (input *Number) SetValue(value int) {

	if input.Step != 0 {

		if input.Min <= value && value <= input.Max {

			if (value - input.Min) % input.Step == 0 {
				input.Value = value
				input.Render()
			}

		}

	}

}

func (input *Number) String() string {

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

func (input *Number) ToValue() js.Value {

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

func (input *Number) Unmount() bool {

	if input.Component.Element != nil {
		input.Component.Element.RemoveEventListener("change", nil)
	}

	return true

}
