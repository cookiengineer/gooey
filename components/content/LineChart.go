//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
// import "github.com/cookiengineer/gooey/interfaces"
import "strconv"
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
	ViewBox    struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"viewbox"`
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

	chart.ViewBox.Width  = 512
	chart.ViewBox.Height = 256

	chart.SetLabelsAndPropertiesAndTypes(labels, properties, types)
	chart.init_events()
	chart.Render()

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

	chart.ViewBox.Width  = 512
	chart.ViewBox.Height = 256

	chart.Parse()
	chart.init_events()
	chart.Render()

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

			viewbox := svg.GetAttribute("viewBox")
			tmp1    := strings.Split(viewbox, " ")

			if len(tmp1) == 4 {

				if tmp1[0] == "0" && tmp1[1] == "0" {

					width,  err1 := strconv.ParseInt(tmp1[2], 10, 64)
					height, err2 := strconv.ParseInt(tmp1[3], 10, 64)

					if err1 == nil && err2 == nil {
						chart.ViewBox.Width  = int(width)
						chart.ViewBox.Height = int(height)
					}

				} else {
					chart.ViewBox.Width  = int(512)
					chart.ViewBox.Height = int(256)
				}

			} else {

				tmp2 := svg.GetAttribute("width")
				tmp3 := svg.GetAttribute("height")

				if tmp2 != "" && tmp3 != "" {

					width,  err1 := strconv.ParseInt(tmp2, 10, 64)
					height, err2 := strconv.ParseInt(tmp3, 10, 64)

					if err1 == nil && err2 == nil {
						chart.ViewBox.Width  = int(width)
						chart.ViewBox.Height = int(height)
					}

				}

			}

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

			tmp2 := make([]string, 0)

			for _, property := range properties {
				tmp2 = append(tmp2, types[property])
			}

			chart.Types = tmp2

			console.Log(chart)

		}

	}

}

func (chart *LineChart) Render() *dom.Element {

	if chart.Component.Element != nil {

		chart.Component.Element.SetAttribute("data-name", chart.Name)
		chart.Component.Element.SetAttribute("data-type", "line-chart")

		svg := chart.Component.Element.QuerySelector("svg")

		if svg != nil {

			svg.SetAttribute("viewBox", strings.Join([]string{
				"0",
				"0",
				strconv.Itoa(chart.ViewBox.Width),
				strconv.Itoa(chart.ViewBox.Height),
			}, " "))

			svg.SetAttribute("width",  strconv.Itoa(chart.ViewBox.Width))
			svg.SetAttribute("height", strconv.Itoa(chart.ViewBox.Height))

			if len(chart.Dataset) > 0 {

				min_value, max_value := calculateChartValuesToMinMax(chart.Dataset, chart.Properties)

				paths := make([]*dom.Element, len(chart.Properties))

				for p, property := range chart.Properties {

					path := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "path")
					path.SetAttribute("data-property", property)

					path_description := renderChartValuesToPath(
						chart.ViewBox.Width,
						chart.ViewBox.Height,
						min_value,
						max_value,
						chart.Dataset,
						property,
					)

					path.SetAttribute("d", path_description)

					paths[p] = path

				}

				svg.ReplaceChildren(paths)

			}

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
