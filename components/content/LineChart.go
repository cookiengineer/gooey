//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
// import "github.com/cookiengineer/gooey/interfaces"
// import "strconv"
import "strings"

// TODO: Custom numeric type
// TODO: uint
// TODO: string? could rendered as different entries
// type numeric float32 | float64 | int | int8 | int16 | int32 | int64

type ChartData map[string]any

type LineChart struct {
	Name       string                `json:"name"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Types      []string              `json:"types"`
	Dataset    []ChartData           `json:"dataset"`
	Component  *components.Component `json:"component"`
}

func NewLineChart(name string, labels []string, properties []string, types []string) LineChart {

	var chart LineChart

	element   := dom.Document.CreateElement("figure")
	component := components.NewComponent(element)

	chart.Component  = &component
	chart.Name       = strings.TrimSpace(strings.ToLower(name))
	chart.Labels     = make([]string, 0)
	chart.Properties = make([]string, 0)
	chart.Types      = make([]string, 0)
	chart.Dataset    = make([]ChartData, 0)

	chart.SetLabelsAndPropertiesAndTypes(labels, properties, types)
	chart.init_events()

	return chart

}

func ToLineChart(element *dom.Element) *LineChart {

	var chart LineChart

	component := components.NewComponent(element)

	chart.Component  = &component
	chart.Name       = ""
	chart.Labels     = make([]string, 0)
	chart.Properties = make([]string, 0)
	chart.Types      = make([]string, 0)
	chart.Dataset    = make([]ChartData, 0)

	chart.Parse()
	chart.init_events()

	return &chart

}

func (chart *LineChart) Disable() bool {

	var result bool

	// TODO: Can <figure> element have custom attributes like disabled?

	return result

}

func (chart *LineChart) Enable() bool {

	var result bool

	// TODO: Can <figure> element have custom attributes like disabled?

	return result

}

func (chart *LineChart) init_events() {

	chart.Component.InitEvent("mousemove")

	chart.Component.Element.AddEventListener("mousemove", dom.ToEventListener(func(event *dom.Event) {

		if event.Target != nil {

			label := event.Target.GetAttribute("data-label")

			if label != "" {
				// TODO: Tooltip integration
			}

		}

	}))

}

func (chart *LineChart) Parse() {

	if chart.Component.Element != nil {

		name := chart.Component.Element.GetAttribute("data-name")

		if name != "" {
			chart.Name = strings.TrimSpace(strings.ToLower(name))
		}

		datalist := chart.Component.Element.QuerySelector("datalist")
		svg      := chart.Component.Element.QuerySelector("svg")

		if datalist != nil && svg != nil {

			elements   := datalist.QuerySelectorAll("data")
			dataset    := make([]ChartData, 0)
			labels     := make([]string, 0)
			properties := make([]string, 0)
			types      := make(map[string]string, 0)
			values     := make([]map[string]string, 0)

			for _, element := range elements {

				property := element.GetAttribute("data-property")
				label    := strings.TrimSpace(element.TextContent)
				typ      := element.GetAttribute("data-type")
				value    := strings.Split(element.GetAttribute("value"), ",")

				if len(value) > 0 {

					if len(values) < len(value) {

						for v := len(values); v < len(value); v++ {
							values = append(values, map[string]string{})
						}

					}

					for v, val := range value {
						values[v][property] = val
					}

					types[property] = typ
					labels          = append(labels, label)
					properties      = append(properties, property)

				}

			}

			if len(values) > 0 && len(values[0]) == len(types) {

				for _, val := range values {
					dataset = append(dataset, ChartData(parseChartValues(val, types)))
				}

			}

			chart.Labels     = labels
			chart.Properties = properties
			chart.Dataset    = dataset

			tmp := make([]string, 0)

			for _, property := range properties {
				tmp = append(tmp, types[property])
			}

			chart.Types = tmp

			console.Log(chart)

		}

		caption := chart.Component.Element.QuerySelector("figcaption")

		if caption != nil {

			// TODO: Parse out caption to label?

		}

	}

}

func (chart *LineChart) Render() *dom.Element {

	if chart.Component.Element != nil {

		svg := chart.Component.Element.QuerySelector("svg")

		if svg != nil {

			// TODO

		}

		caption := chart.Component.Element.QuerySelector("figcaption")

		if caption != nil {

			// TODO

		}

	}

	return chart.Component.Element

}

func (chart *LineChart) Add(data ChartData) {
	chart.Dataset = append(chart.Dataset, data)
}

func (chart *LineChart) Remove(indexes []int) {

	dataset := make([]ChartData, 0)

	for d, data := range chart.Dataset {

		found := false

		for _, index := range indexes {

			if d == index {
				found = true
				break
			}

		}

		if found == false {
			dataset = append(dataset, data)
		}

	}

	chart.Dataset = dataset

}

func (chart *LineChart) SetData(dataset []ChartData) {
	chart.Dataset = dataset
}

func (chart *LineChart) SetLabelsAndPropertiesAndTypes(labels []string, properties []string, types []string) bool {

	var result bool

	if len(labels) == len(properties) && len(labels) == len(types) {

		chart.Labels     = labels
		chart.Properties = properties
		chart.Types      = types

		result = true

	}

	return result

}

func (chart *LineChart) String() string {

	html := "<figure"
	html += " data-type=\"line-chart\""

	if chart.Name != "" {
		html += " data-name=\"" + chart.Name + "\""
	}

	html += ">"

	// TODO

	html += "</figure>"

	return html

}
