//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strconv"
import "strings"

type TableData map[string]any

type Table struct {
	Name       string                `json:"name"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Dataset    []TableData           `json:"dataset"`
	Component  *components.Component `json:"component"`
	Selectable bool                  `json:"selectable"`
	selected   []bool
}

func NewTable(name string, labels []string, properties []string, selectable bool) Table {

	var table Table

	element   := bindings.Document.CreateElement("table")
	component := components.NewComponent(element)

	table.Component  = &component
	table.Name       = strings.TrimSpace(strings.ToLower(name))
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Dataset    = make([]TableData, 0)
	table.Selectable = selectable
	table.selected   = make([]bool, 0)

	table.SetLabelsAndProperties(labels, properties)

	table.Component.InitEvent("change-select")
	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		if event.Target != nil {

			action := event.Target.GetAttribute("data-action")

			if action != "" {

				if action == "select" {

					row := event.Target.QueryParent("tr")

					// TODO: check for td
					// TODO: check for th

					// TODO: select all for <thead><th><input type=checkbox></th>...</thead>
					// TODO: select current for <tbody><td><input type=checkbox></td>...</tbody>

					event.PreventDefault()
					event.StopPropagation()

					table.Render()

				} else if action == "sort" {

					property := event.Target.GetAttribute("data-property")
					sort     := event.Target.GetAttribute("data-sort")

					// TODO: Change sorting via property and sort direction

					event.PreventDefault()
					event.StopPropagation()

					table.Render()

				} else {

					event.PreventDefault()
					event.StopPropagation()

					table.Component.FireEventListeners("action", map[string]string{
						"action": action,
					})

				}

			}

		}

	}))

	// TODO: table.Data = []TableData?
	// TableData = struct { Value: map[string]any }?

	return table

}

func ToTable(element *dom.Element) Table {

	var table Table

	component := components.NewComponent(element)

	table.Component  = &component
	table.Name       = ""
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Dataset    = make([]TableData, 0)
	table.Selectable = element.HasAttribute("data-selectable")
	table.selected   = make([]bool, 0)

	table.Parse()

	table.Component.InitEvent("change-select")
	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		// TODO: Port event listener from NewTable() when ready

	}))

	return table

}

func (table *Table) Disable() bool {

	var result bool

	// TODO: if table.Selectable then Disable input[type=checkbox] elements
	// TODO: Disable footer elements

	return result

}

func (table *Table) Enable() bool {

	var result bool

	// TODO: if table.Selectable then Enable input[type=checkbox] elements
	// TODO: Enable footer elements

	return result

}

func (table *Table) Parse() {

	if table.Component.Element != nil {

		name := table.Component.Element.GetAttribute("data-name")

		if name != "" {
			table.Name = strings.TrimSpace(strings.ToLower(name))
		}

		thead := table.Component.Element.QuerySelector("thead")

		if thead != nil {

			elements := thead.QuerySelectorAll("th")

			for _, element := range elements {

				label    := strings.TrimSpace(element.TextContent)
				property := element.GetAttribute("data-property")

			}

			// TODO: Parse out labels

		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {
			// TODO: Parse out content/tabledata
		}

	}

}

func (table *Table) Render() *dom.Element {

	if table.Component.Element != nil {

		// TODO: Typecast and render strings
		// TODO: Write a helper method called renderValue() or something

	}

	return table.Component.Element

}

func (table *Table) SetData(dataset []TableData) {
	table.Dataset = dataset
}

func (table *Table) SetLabelsAndProperties(labels []string, properties []string) bool {

	var result bool

	if len(labels) == len(properties) {

		table.Labels = labels
		table.Properties = properties
		result = true

	}

	return result

}

func (table *Table) String() string {

	html := "<table"
	html += ">"

	html += "<thead>"
	html += "<tr>"

	if table.Selectable == true {
		html += "<th><input type=\"checkbox\" data-action=\"select\"/></th>"
	}

	for l, label := range table.Labels {

		property := table.Properties[l]

		html += "<th data-property=\"" + property + "\">"
		html += label
		html += "</th>"

	}

	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	// TODO: Implement sorted rendering
	// TODO: Presort the ids, and then render via table.Dataset[sorted_ids[i]]

	for d := 0; d < len(table.Dataset); d++ {

		html += "<tr data-id=\"" + strconv.FormatInt(int64(d), 10) + "\""

		if table.selected[d] == true {
			html += " data-select=\"true\""
		}

		html += ">"

		if table.Selectable == true {
			html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
		}

		for _, property := range table.Properties {

			value, ok := table.Dataset[d][property]

			if ok == true {
				html += "<td>" + renderTableValue(value) + "</td>"
			} else {
				html += "<td></td>"
			}

		}

		html += "</tr>"

	}

	html += "</tbody>"

	html += "<tfoot>"

	// TODO: Render table actions?

	html += "</tfoot>"

	html += "</table>"

	return html

}
