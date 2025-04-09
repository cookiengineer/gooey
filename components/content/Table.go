//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "strings"

type TableData map[string]any

type Table struct {
	Name       string                `json:"name"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Dataset    []TableData           `json:"dataset"`
	Component  *components.Component `json:"component"`
	Selectable bool                  `json:"selectable"`
	selected   []bool                `json:"selected"`

	// TODO: internal header
	// TODO: internal footer

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

	table.SetLabels(labels)
	table.SetProperties(properties)

	table.Component.InitEvent("change-select")
	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		if event.Target != nil {

			action := event.Target.GetAttribute("data-action")

			if action != "" {

				if action == "select" {

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

	table.Parse()

	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	// TODO: Table Content
	// TODO: Table Dataset

	return table

}

func (table *Table) Disable() bool {

	var result bool

	// TODO: Disable footer elements

	return result

}

func (table *Table) Enable() bool {

	var result bool

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

		// TODO

	}

}

func (table *Table) Render() *dom.Element {

	if table.Component.Element != nil {

		// TODO

	}

	return table.Component.Element

}

func (table *Table) String() string {

	html := "<table"
	html += ">"

	html += "<thead>"

	// TODO: Render table.Labels

	html += "</thead>"

	html += "<tbody>"

	// TODO: Render table.Dataset via table.Properties

	html += "</tbody>"

	html += "<tfoot>"

	// TODO: Render table actions?

	html += "</tfoot>"

	html += "</table>"

	return html

}
