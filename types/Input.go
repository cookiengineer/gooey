package types

type Input string

const (

	// ui.Checkbox
	InputCheckbox      Input = "checkbox"

	// ui.Input
	InputDate          Input = "date"
	InputDatetimeLocal Input = "datetime-local"
	InputEmail         Input = "email"
	InputFile          Input = "file"
	InputMonth         Input = "month"
	InputNumber        Input = "number"
	InputPassword      Input = "password"
	InputSearch        Input = "search"
	InputTel           Input = "tel"
	InputText          Input = "text"
	InputTime          Input = "time"
	InputURL           Input = "url"
	InputWeek          Input = "week"

	// ui.Textarea uses same as InputText
	InputTextarea      Input = "text"

)

func (typ Input) String() string {
	return string(typ)
}
