package ui

type InputType string

const (
	InputTypeDate          InputType = "date"
	InputTypeDatetimeLocal InputType = "datetime-local"
	InputTypeEmail         InputType = "email"
	InputTypeFile          InputType = "file"
	InputTypeMonth         InputType = "month"
	InputTypeNumber        InputType = "number"
	InputTypePassword      InputType = "password"
	InputTypeSearch        InputType = "search"
	InputTypeTel           InputType = "tel"
	InputTypeText          InputType = "text"
	InputTypeTime          InputType = "time"
	InputTypeURL           InputType = "url"
	InputTypeWeek          InputType = "week"
)

func (typ InputType) String() string {
	return string(typ)
}
