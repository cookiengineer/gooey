//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/interfaces"
import "strconv"
import "strings"

type Table struct {
	Name       string        `json:"name"`
	Labels     []string      `json:"labels"`
	Properties []string      `json:"properties"`
	Types      []string      `json:"types"`
	Dataset    *data.Dataset `json:"dataset"`
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

	element   := dom.Document.CreateElement("table")
	component := components.NewComponent(element)
	dataset   := data.NewDataset(0)

	table.Dataset    = &dataset
	table.Component  = &component
	table.Name       = strings.TrimSpace(strings.ToLower(name))
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Types      = make([]string, 0)
	table.Selectable = selectable
	table.selected   = make([]bool, 0)
	table.sorted     = make([]int, 0)
	table.sortby     = ""

	table.Footer.Content.Left   = make([]interfaces.Component, 0)
	table.Footer.Content.Center = make([]interfaces.Component, 0)
	table.Footer.Content.Right  = make([]interfaces.Component, 0)

	table.SetLabelsAndPropertiesAndTypes(labels, properties, types)
	table.Mount()
	table.Render()

	return table

}

func ToTable(element *dom.Element) *Table {

	var table Table

	component := components.NewComponent(element)
	dataset   := data.NewDataset(0)

	table.Dataset    = &dataset
	table.Component  = &component
	table.Name       = ""
	table.Labels     = make([]string, 0)
	table.Properties = make([]string, 0)
	table.Types      = make([]string, 0)
	table.Selectable = element.HasAttribute("data-selectable")
	table.selected   = make([]bool, 0)
	table.sorted     = make([]int, 0)
	table.sortby     = ""

	table.Footer.Content.Left   = make([]interfaces.Component, 0)
	table.Footer.Content.Center = make([]interfaces.Component, 0)
	table.Footer.Content.Right  = make([]interfaces.Component, 0)

	table.Mount()

	return &table

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

func (table *Table) Mount() bool {

	if table.Component != nil {
		table.Component.InitEvent("action")
	}

	if table.Component.Element != nil {

		name := table.Component.Element.GetAttribute("data-name")

		if name != "" {
			table.Name = strings.TrimSpace(strings.ToLower(name))
		}

		selectable := table.Component.Element.GetAttribute("data-selectable")

		if selectable == "true" {
			table.Selectable = true
		} else {
			table.Selectable = false
		}

		thead := table.Component.Element.QuerySelector("thead")

		if thead != nil && len(table.Labels) == 0 && len(table.Properties) == 0 && len(table.Types) == 0 {

			elements   := thead.QuerySelectorAll("th")
			labels     := make([]string, 0)
			properties := make([]string, 0)
			types      := make([]string, 0)
			selectable := table.Selectable

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
			dataset  := data.NewDataset(len(rows))
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

						if id != -1 && id >= 0 && id < dataset.Length() {
							dataset.Set(id, data.ParseData(values, types))
							selected[id] = row.HasAttribute("data-select")
							sorted[r] = id
						} else {
							dataset.Set(id, data.ParseData(values, types))
							selected[r] = row.HasAttribute("data-select")
							sorted[r] = id
						}

					}

				}

			}

			table.Dataset = &dataset
			table.sorted = sorted
			table.selected = selected

		} else {

			console.Group("Table Body: Invalid Markup")
			console.Error("Expected <tr>...</tr>")
			console.GroupEnd("Table Body: Invalid Markup")

		}

		tfoot := table.Component.Element.QuerySelector("tfoot")

		if tfoot != nil {

			tmp := tfoot.QuerySelectorAll("td")

			if len(tmp) == 3 {

				buttons_left := tmp[0].QuerySelectorAll("button")

				for _, button := range buttons_left {
					table.Footer.Content.Left = append(table.Footer.Content.Left, ui.ToButton(button))
				}

				elements_center := tmp[1].QuerySelectorAll("button, label, input")

				for _, element := range elements_center {

					if element.TagName == "BUTTON" {
						table.Footer.Content.Center = append(table.Footer.Content.Center, ui.ToButton(element))
					} else if element.TagName == "LABEL" {
						table.Footer.Content.Center = append(table.Footer.Content.Center, ui.ToLabel(element))
					} else if element.TagName == "INPUT" {
						table.Footer.Content.Center = append(table.Footer.Content.Center, ui.ToInput(element))
					}

				}

				buttons_right := tmp[2].QuerySelectorAll("button")

				for _, button := range buttons_right {
					table.Footer.Content.Right = append(table.Footer.Content.Right, ui.ToButton(button))
				}

			} else {

				console.Group("Table Footer: Invalid Markup")
				console.Error("Expected <tr><td></td><td colspan></td><td></td></tr>")
				console.GroupEnd("Table Footer: Invalid Markup")

			}

		}

		table.Component.Element.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			if event.Target != nil {

				action := event.Target.GetAttribute("data-action")

				if action == "select" {

					th := event.Target.QueryParent("th")

					if th != nil {

						is_active := event.Target.Value.Get("checked").Bool()

						if is_active == true {

							for s := 0; s < len(table.selected); s++ {
								table.selected[s] = true
							}

							table.Render()

						} else {

							for s := 0; s < len(table.selected); s++ {
								table.selected[s] = false
							}

							table.Render()

						}

					} else {

						is_active := event.Target.Value.Get("checked").Bool()
						tmp       := event.Target.QueryParent("tr").GetAttribute("data-id")

						if is_active == true {

							num, err := strconv.ParseInt(tmp, 10, 64)

							if err == nil {

								index := int(num)

								if index >= 0 && index < table.Dataset.Length() {

									table.selected[index] = true
									table.Render()

								}

							}

						} else {

							num, err := strconv.ParseInt(tmp, 10, 64)

							if err == nil {

								index := int(num)

								if index >= 0 && index < table.Dataset.Length() {

									input := table.Component.Element.QuerySelector("thead input[data-action=\"select\"]")

									if input != nil {
										input.Value.Set("checked", false)
									}

									table.selected[index] = false
									table.Render()

								}

							}

						}

						event.PreventDefault()
						event.StopPropagation()

					}

				} else if action == "sort" {

					thead := table.Component.Element.QuerySelector("thead")
					th    := event.Target.QueryParent("th")

					if thead != nil && th != nil {

						property := th.GetAttribute("data-property")
						ths      := thead.QuerySelectorAll("th")

						for _, th := range ths {
							th.RemoveAttribute("data-sort")
						}

						if table.sortby != property {

							th.SetAttribute("data-sort", "ascending")

							table.sorted = table.Dataset.SortByProperty(property)
							table.sortby = property

							table.Render()

						}

						event.PreventDefault()
						event.StopPropagation()

					}

				} else if action != "" {

					table.Component.FireEventListeners("action", map[string]any{
						"action": action,
					})

				}

			}

		}))

		return true

	} else {
		return false
	}

}

func (table *Table) Render() *dom.Element {

	if table.Component.Element != nil {

		if table.Name != "" {
			table.Component.Element.SetAttribute("data-name", table.Name)
		}

		if table.Selectable == true {
			table.Component.Element.SetAttribute("data-selectable", "true")
		}

		tbody := table.Component.Element.QuerySelector("tbody")

		if tbody != nil {

			elements := make([]*dom.Element, 0)

			for _, position := range table.sorted {

				tr := dom.Document.CreateElement("tr")

				tr.SetAttribute("data-id", strconv.FormatInt(int64(position), 10))

				if table.selected[position] == true {
					tr.SetAttribute("data-select", "true")
				}

				html := ""

				if table.Selectable == true {

					if table.selected[position] == true {
						html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
					} else {
						html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
					}

				}

				values, _ := table.Dataset.Get(position).String()

				for _, property := range table.Properties {

					val, ok := values[property]

					if ok == true {
						html += "<td>" + val + "</td>"
					} else {
						html += "<td></td>"
					}

				}

				tr.SetInnerHTML(html)

				elements = append(elements, tr)

			}

			tbody.ReplaceChildren(elements)

		}

		tfoot := table.Component.Element.QuerySelector("tfoot")

		if tfoot != nil && len(table.Labels) >= 3 {

			tmp := tfoot.QuerySelectorAll("td")

			if len(tmp) == 0 {
				tfoot.SetInnerHTML("<tr><td></td><td></td><td></td></tr>")
				tmp = tfoot.QuerySelectorAll("td")
			}

			if len(tmp) == 3 {

				elements_left   := make([]*dom.Element, 0)
				elements_center := make([]*dom.Element, 0)
				elements_right  := make([]*dom.Element, 0)

				for _, component := range table.Footer.Content.Left {
					elements_left = append(elements_left, component.Render())
				}

				for _, component := range table.Footer.Content.Center {
					elements_center = append(elements_center, component.Render())
				}

				for _, component := range table.Footer.Content.Right {
					elements_right = append(elements_right, component.Render())
				}

				colspan := len(table.Labels) - 2

				if table.Selectable == true {
					tmp[0].SetAttribute("colspan", "2")
				} else {
					tmp[0].RemoveAttribute("colspan")
				}

				tmp[1].SetAttribute("colspan", strconv.Itoa(colspan))

				tmp[0].ReplaceChildren(elements_left)
				tmp[1].ReplaceChildren(elements_center)
				tmp[2].ReplaceChildren(elements_right)

			}

		}

	}

	return table.Component.Element

}

func (table *Table) Add(data data.Data) bool {

	var result bool = false

	if table.Dataset.Add(data) == true {

		table.selected = append(table.selected, false)
		table.sorted   = append(table.sorted, table.Dataset.Length() - 1)
		result = true

	}

	return result

}

func (table *Table) Deselect(indexes []int) {

	for _, index := range indexes {
		table.selected[index] = false
	}

}

func (table *Table) Query(query string) interfaces.Component {

	selectors := utils.SplitQuery(query)

	if len(selectors) == 1 {

		if table.Component.Element != nil {

			if utils.MatchesQuery(table.Component.Element, selectors[0]) == true {
				return table
			}

		}

	}

	return nil

}

func (table *Table) Remove(indexes []int) {

	entries  := make([]data.Data, 0)
	selected := make([]bool, 0)
	sorted   := make([]int, 0)

	for d, data := range *table.Dataset {

		found := false
		is_selected := table.selected[d]

		for _, index := range indexes {

			if d == index {
				found = true
				break
			}

		}

		if found == false {
			entries  = append(entries, *data)
			selected = append(selected, is_selected)
		}

	}

	for e := 0; e < len(entries); e++ {
		sorted = append(sorted, e)
	}

	dataset := data.ToDataset(entries)

	table.Dataset  = &dataset
	table.selected = selected
	table.sortby   = ""
	table.sorted   = sorted

}

func (table *Table) Select(indexes []int) {

	for _, index := range indexes {
		table.selected[index] = true
	}

}

func (table *Table) Selected() ([]int, []data.Data) {

	result_indexes := make([]int, 0)
	result_dataset := make([]data.Data, 0)

	for s, value := range table.selected {

		if value == true {

			data := table.Dataset.Get(s)

			if data != nil {
				result_indexes = append(result_indexes, s)
				result_dataset = append(result_dataset, *data)
			}

		}

	}

	return result_indexes, result_dataset

}

func (table *Table) SetDataset(dataset data.Dataset) {

	table.Dataset = &dataset
	table.selected = make([]bool, dataset.Length())

	table.selected = make([]bool, dataset.Length())
	table.sortby = ""
	table.sorted = make([]int, dataset.Length())

	for d := 0; d < dataset.Length(); d++ {
		table.sorted[d] = d
	}

}

func (table *Table) SetData(entries []data.Data) {

	dataset := data.ToDataset(entries)
	table.Dataset = &dataset

	table.selected = make([]bool, dataset.Length())
	table.sortby = ""
	table.sorted = make([]int, dataset.Length())

	for d := 0; d < dataset.Length(); d++ {
		table.sorted[d] = d
	}

}

func (table *Table) SetCenter(components []interfaces.Component) {
	table.Footer.Content.Center = components
}

func (table *Table) SetLeft(components []interfaces.Component) {
	table.Footer.Content.Left = components
}

func (table *Table) SetRight(components []interfaces.Component) {
	table.Footer.Content.Right = components
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

func (table *Table) SortBy(prop string) bool {

	var result bool

	thead := table.Component.Element.QuerySelector("thead")

	if thead != nil {

		ths   := thead.QuerySelectorAll("th")
		found := false

		for _, th := range ths {
			th.RemoveAttribute("data-sort")
		}

		for _, th := range ths {

			property := th.GetAttribute("data-property")

			if property == prop {
				th.SetAttribute("data-sort", "ascending")
				found = true
				break
			}

		}

		if found == true {
			table.sorted = table.Dataset.SortByProperty(prop)
			table.sortby = prop
			result = true
		}

	}

	return result

}

func (table *Table) String() string {

	html := "<table"

	if table.Name != "" {
		html += " data-name=\"" + table.Name + "\""
	}

	if table.Selectable == true {
		html += " data-selectable=\"" + strconv.FormatBool(table.Selectable) + "\""
	}

	html += ">"

	html += "<thead>"
	html += "<tr>"

	if table.Selectable == true {
		html += "<th><input type=\"checkbox\" data-action=\"select\"/></th>"
	}

	for l, label := range table.Labels {

		property := table.Properties[l]
		typ      := table.Types[l]

		html += "<th data-property=\"" + property + "\" data-type=\"" + typ + "\""

		if table.sortby == property {
			html += " data-sort=\"ascending\""
		}

		html += ">"
		html += "<label data-action=\"sort\">"
		html += label
		html += "</label>"
		html += "</th>"

	}

	html += "</tr>"
	html += "</thead>"

	html += "<tbody>"

	for _, position := range table.sorted {

		html += "<tr data-id=\"" + strconv.FormatInt(int64(position), 10) + "\""

		if table.selected[position] == true {
			html += " data-select=\"true\""
		}

		html += ">"

		if table.Selectable == true {

			if table.selected[position] == true {
				html += "<td><input type=\"checkbox\" data-action=\"select\" checked/></td>"
			} else {
				html += "<td><input type=\"checkbox\" data-action=\"select\"/></td>"
			}

		}

		values, _ := table.Dataset.Get(position).String()

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
	html += "<tr>"
	html += "<td"

	if table.Selectable == true {
		html += " colspan=\"2\""
	}

	html += ">"

	if len(table.Footer.Content.Left) > 0 {

		for _, component := range table.Footer.Content.Left {
			html += component.String()
		}

	}

	html += "</td>"
	html += "<td"

	if len(table.Labels) >= 3 {
		html += " colspan=\"" + strconv.Itoa(len(table.Labels) - 2) + "\""
	}

	if len(table.Footer.Content.Center) > 0 {

		for _, component := range table.Footer.Content.Center {
			html += component.String()
		}

	}

	html += "</td>"
	html += "<td>"

	if len(table.Footer.Content.Right) > 0 {

		for _, component := range table.Footer.Content.Right {
			html += component.String()
		}

	}

	html += "</td>"
	html += "</tr>"
	html += "</tfoot>"

	html += "</table>"

	return html

}

func (table *Table) Unmount() bool {

	if table.Component.Element != nil {
		table.Component.Element.RemoveEventListener("click", nil)
	}

	if len(table.Footer.Content.Left) > 0 {

		for _, component := range table.Footer.Content.Left {
			component.Unmount()
		}

	}

	if len(table.Footer.Content.Center) > 0 {

		for _, component := range table.Footer.Content.Center {
			component.Unmount()
		}

	}

	if len(table.Footer.Content.Right) > 0 {

		for _, component := range table.Footer.Content.Right {
			component.Unmount()
		}

	}

	return true

}
