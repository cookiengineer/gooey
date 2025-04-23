//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/interfaces"
import "slices"
import "sort"
import "strconv"
import "strings"

type TableData map[string]any

type Table struct {
	Name       string                `json:"name"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Types      []string              `json:"types"`
	Dataset    []TableData           `json:"dataset"`
	Footer     struct {
		Content struct {
			Left   []interfaces.Component `json:"left"`
			Center []interfaces.Component `json:"center"`
			Right  []interfaces.Component `json:"right"`
		} `json:"content"`
	} `json:"footer"`
	Component  *components.Component `json:"component"`
	Selectable bool                  `json:"selectable"`
	selected   []bool
	sorted     []int
	sortby     string
}

func NewTable(name string, labels []string, properties []string, types []string, selectable bool) Table {

	var table Table

	element   := bindings.Document.CreateElement("table")
	component := components.NewComponent(element)

	table.Component  = &component
	table.Name       = strings.TrimSpace(strings.ToLower(name))
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Types      = make([]string, 0)
	table.Dataset    = make([]TableData, 0)
	table.Selectable = selectable
	table.selected   = make([]bool, 0)
	table.sorted     = make([]int, 0)
	table.sortby     = ""

	table.Footer.Content.Left   = make([]interfaces.Component, 0)
	table.Footer.Content.Center = make([]interfaces.Component, 0)
	table.Footer.Content.Right  = make([]interfaces.Component, 0)

	table.SetLabelsAndPropertiesAndTypes(labels, properties, types)

	table.Component.InitEvent("change-select")
	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		// TODO: Backport click event handler from ToTable()

	}))

	return table

}

func ToTable(element *dom.Element) Table {

	var table Table

	component := components.NewComponent(element)

	table.Component  = &component
	table.Name       = ""
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Types      = make([]string, 0)
	table.Dataset    = make([]TableData, 0)
	table.Selectable = element.HasAttribute("data-selectable")
	table.selected   = make([]bool, 0)
	table.sorted     = make([]int, 0)
	table.sortby     = ""

	table.Footer.Content.Left   = make([]interfaces.Component, 0)
	table.Footer.Content.Center = make([]interfaces.Component, 0)
	table.Footer.Content.Right  = make([]interfaces.Component, 0)

	table.Parse()

	table.Component.InitEvent("change-select")
	table.Component.InitEvent("change-sort")
	table.Component.InitEvent("action")

	table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {

		if event.Target != nil {

			action := event.Target.GetAttribute("data-action")

			if action != "" {

				if action == "select" {
				} else if action == "sort" {

					th := event.Target.QueryParent("th")

					if th != nil {

						property := event.Target.GetAttribute("data-property")
						typ      := event.Target.GetAttribute("data-type")

						event.PreventDefault()
						event.StopPropagation()

						if table.sortby != property {

							tmp_sorted   := make([]int, len(table.Dataset))
							tmp_selected := make([]bool, len(table.Dataset))

							for s := 0; s < len(table.Dataset); s++ {
								tmp_sorted[s] = s
							}

							sort.Slice(tmp_sorted, func(a int, b int) bool {

								value_a, ok_a := table.Dataset[a][property]
								value_b, ok_b := table.Dataset[b][property]

								// TODO: typecast value into sortable/comparable value?

								if ok_a == true && ok_b == true {
									return value_a < value_b
								} else {
									return false
								}

							})

							for _, id := range table.sorted {

								if table.selected[id] == true {

									index := slices.Index(tmp_sorted, id)

									if index != -1 {
										tmp_selected[index] = true
									}

								}

							}

							table.sortby   = property
							table.sorted   = tmp_sorted
							table.selected = tmp_selected

							table.Render()

						}

					}

				}

			}

		}

	}))

	return table

}

func (table *Table) Disable() bool {

	var result bool

	inputs := table.Component.Element.QuerySelectorAll("input[type=\"checkbox\"]")

	if len(inputs) > 0 {

		for _, element := range inputs {
			element.SetAttribute("disabled", "")
		}

		result = true

	}

	if len(table.Footer.Content.Left) > 0 || len(table.Footer.Content.Center) > 0 || len(table.Footer.Content.Right) > 0 {

		for _, component := range table.Footer.Content.Left {
			component.Disable()
		}

		for _, component := range table.Footer.Content.Center {
			component.Disable()
		}

		for _, component := range table.Footer.Content.Right {
			component.Disable()
		}

		result = true

	}

	return result

}

func (table *Table) Enable() bool {

	var result bool

	inputs := table.Component.Element.QuerySelectorAll("input[type=\"checkbox\"]")

	if len(inputs) > 0 {

		for _, element := range inputs {
			element.RemoveAttribute("disabled")
		}

		result = true

	}

	if len(table.Footer.Content.Left) > 0 || len(table.Footer.Content.Center) > 0 || len(table.Footer.Content.Right) > 0 {

		for _, component := range table.Footer.Content.Left {
			component.Enable()
		}

		for _, component := range table.Footer.Content.Center {
			component.Enable()
		}

		for _, component := range table.Footer.Content.Right {
			component.Enable()
		}

		result = true

	}

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

			elements   := thead.QuerySelectorAll("th")
			labels     := make([]string, 0)
			properties := make([]string, 0)
			types      := make([]string, 0)
			selectable := false

			if len(elements) > 0 {

				checkbox := elements[0].QuerySelector("input[type=\"checkbox\"]")

				if checkbox != nil {
					elements   = elements[1:]
					selectable = true
				}

				for e := 0; e < len(elements); e++ {

					element := elements[e]

					label    := strings.TrimSpace(element.TextContent)
					property := element.GetAttribute("data-property")
					typ      := element.GetAttribute("data-type")

					if typ == "" {
						typ = "string"
					}

					if label != "" && property != "" {

						labels     = append(labels, label)
						properties = append(properties, property)
						types      = append(types, typ)

					}

				}

			}

			table.Labels     = labels
			table.Properties = properties
			table.Types      = types
			table.Selectable = selectable

		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {

			rows     := tbody.QuerySelectorAll("tr")
			dataset  := make([]TableData, len(rows))
			sorted   := make([]int, len(rows))
			selected := make([]bool, len(rows))

			for r, row := range rows {

				var id int = -1

				id_str := row.GetAttribute("data-id")

				if id_str != "" {

					tmp, err := strconv.ParseInt(id_str, 10, 0)

					if err == nil {
						id = int(tmp)
					}

				} else {
					id = int(r)
				}

				elements := row.QuerySelectorAll("td")

				if len(elements) > 0 {

					checkbox := elements[0].QuerySelector("input[type=\"checkbox\"]")
					values   := make(map[string]string)
					types    := make(map[string]string)

					if checkbox != nil {
						elements = elements[1:]
					}

					for e := 0; e < len(elements); e++ {

						key := table.Properties[e]
						typ := table.Types[e]
						val := strings.TrimSpace(elements[e].TextContent)

						if key != "" && typ != "" && val != "" {
							values[key] = val
							types[key]  = typ
						}

					}

					if len(values) == len(types) {

						if id != -1 && id >= 0 && id <= len(dataset) - 1 {
							dataset[id] = TableData(parseTableValues(values, types))
							sorted[r] = id
						} else {
							dataset[r] = TableData(parseTableValues(values, types))
							sorted[r] = id
						}

					}

				}

				selected[r] = row.HasAttribute("data-select")

			}

			table.Dataset = dataset
			table.sorted = sorted

		}

		console.Log(table)

	}

}

func (table *Table) Render() *dom.Element {

	if table.Component.Element != nil {

		console.Log(table.sortby)
		console.Log(table.sorted)
		// TODO: Typecast and render strings
		// TODO: Write a helper method called renderValue() or something

	}

	return table.Component.Element

}

func (table *Table) SetData(dataset []TableData) {
	table.Dataset = dataset
}

func (table *Table) SetLabelsAndPropertiesAndTypes(labels []string, properties []string, types []string) bool {

	var result bool

	if len(labels) == len(properties) && len(labels) == len(types) {

		table.Labels     = labels
		table.Properties = properties
		table.Types      = types

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
		typ      := table.Types[l]

		html += "<th data-property=\"" + property + "\" data-type=\"" + typ + "\">"
		html += "<label data-action=\"sort\">"
		html += label
		html += "</label>"
		html += "</th>"

	}

	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	for _, position := range table.sorted {

		data := table.Dataset[position]

		html += "<tr data-id=\"" + strconv.FormatInt(int64(position), 10) + "\""

		if table.selected[position] == true {
			html += " data-select=\"true\""
		}

		html += ">"

		if table.Selectable == true {
			html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
		}

		values, _ := renderTableValues(data)

		for _, property := range table.Properties {

			val, ok := values[property]

			if ok == true {
				html += "<td>" + val + "</td>"
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
