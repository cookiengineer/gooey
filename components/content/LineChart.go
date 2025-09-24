//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/data"
import "strconv"
import "strings"

type LineChart struct {
	Name       string                `json:"name"`
	Disabled   bool                  `json:"disabled"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Types      []string              `json:"types"`
	Dataset    *data.Dataset         `json:"dataset"`
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
	dataset   := data.NewDataset(0)

	chart.Dataset    = &dataset
	chart.Component  = &component
	chart.Name       = strings.TrimSpace(strings.ToLower(name))
	chart.Labels     = make([]string, 0)
	chart.Properties = make([]string, 0)
	chart.Types      = make([]string, 0)

	chart.ViewBox.Width  = 512
	chart.ViewBox.Height = 256

	chart.SetLabelsAndPropertiesAndTypes(labels, properties, types)
	chart.Mount()
	chart.Render()

	return chart

}

func ToLineChart(element *dom.Element) *LineChart {

	var chart LineChart

	component := components.NewComponent(element)
	dataset   := data.NewDataset(0)

	chart.Dataset    = &dataset
	chart.Component  = &component
	chart.Name       = ""
	chart.Labels     = make([]string, 0)
	chart.Properties = make([]string, 0)
	chart.Types      = make([]string, 0)

	chart.ViewBox.Width  = 512
	chart.ViewBox.Height = 256

	chart.Mount()
	chart.Render()

	return &chart

}

func (chart *LineChart) Disable() bool {

	chart.Disabled = true
	chart.Render()

	return true

}

func (chart *LineChart) Enable() bool {

	chart.Disabled = false
	chart.Render()

	return true

}

func (chart *LineChart) Mount() bool {

	if chart.Component != nil {
		chart.Component.InitEvent("mousemove")
	}

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
			dataset    := data.NewDataset(0)
			labels     := make([]string, 0)
			properties := make([]string, 0)
			types      := make(map[string]string)
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
					dataset.Add(data.ParseData(val, types))
				}

			}

			chart.Dataset    = &dataset
			chart.Labels     = labels
			chart.Properties = properties

			tmp2 := make([]string, 0)

			for _, property := range properties {
				tmp2 = append(tmp2, types[property])
			}

			chart.Types = tmp2

		}

		chart.Component.Element.AddEventListener("mousemove", dom.ToEventListener(func(event *dom.Event) {

			if chart.Disabled == false && event.Target != nil {

				if event.Target.TagName == "LABEL" {

					property := event.Target.GetAttribute("data-property")

					if property != "" {

						svg    := chart.Component.Element.QuerySelector("svg")
						layers := chart.Component.Element.QuerySelectorAll("svg g")

						var foreground *dom.Element = nil

						background := make([]*dom.Element, 0)

						for _, layer := range layers {

							if layer.GetAttribute("data-property") == property {
								layer.SetAttribute("data-state", "active")
								foreground = layer
							} else {
								layer.RemoveAttribute("data-state")
								background = append(background, layer)
							}

						}

						if svg != nil {
							svg.ReplaceChildren(background)
							svg.Append(foreground)
						}

					}

				}

			}

		}))

		return true

	} else {
		return false
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

			if chart.Dataset.Length() > 0 {

				min_value, max_value := calculateChartDatasetMinMax(chart.Dataset, chart.Properties)

				layers := make([]*dom.Element, len(chart.Properties))

				for p, property := range chart.Properties {

					layer := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "g")
					layer.SetAttribute("data-property", property)
					layer.SetAttribute("data-palette", strconv.Itoa(p+1))

					path, texts := renderLineChartDataset(
						chart.Dataset,
						chart.ViewBox.Width,
						chart.ViewBox.Height,
						min_value,
						max_value,
						property,
					)

					if path != nil {
						layer.Append(path)
					}

					for _, text := range texts {

						if text != nil {
							layer.Append(text)
						}

					}

					layers[p] = layer

				}

				svg.ReplaceChildren(layers)

			}

		}

		figcaption := chart.Component.Element.QuerySelector("figcaption")

		if figcaption != nil {

			labels := make([]*dom.Element, len(chart.Properties))

			for p, property := range chart.Properties {

				label := dom.Document.CreateElement("label")
				label.SetAttribute("data-palette", strconv.Itoa(p+1))
				label.SetAttribute("data-property", property)
				label.SetInnerHTML(chart.Labels[p])

				labels[p] = label

			}

			figcaption.ReplaceChildren(labels)

		}

	}

	return chart.Component.Element

}

func (chart *LineChart) Add(data data.Data) bool {
	return chart.Dataset.Add(data)
}

func (chart *LineChart) Remove(indexes []int) {

	entries := make([]data.Data, 0)

	for d, data := range *chart.Dataset {

		found := false

		for _, index := range indexes {

			if d == index {
				found = true
				break
			}

		}

		if found == false {
			entries = append(entries, *data)
		}

	}

	dataset := data.ToDataset(entries)

	chart.Dataset = &dataset

}

func (chart *LineChart) SetDataset(dataset data.Dataset) {
	chart.Dataset = &dataset
}

func (chart *LineChart) SetData(entries []data.Data) {
	dataset := data.ToDataset(entries)
	chart.Dataset = &dataset
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

	if chart.Name != "" {
		html += " data-name=\"" + chart.Name + "\""
	}

	html += " data-type=\"line-chart\""
	html += ">"

	html += "<datalist>"

	values, _ := chart.Dataset.Join(",")

	for p, property := range chart.Properties {

		value := values[property]
		label := chart.Labels[p]
		typ   := chart.Types[p]

		html += "<data"
		html += " data-property=\"" + property + "\""
		html += " data-type=\"" + typ + "\""
		html += " value=\"" + value + "\""
		html += ">"
		html += label
		html += "</data>"

	}

	html += "</datalist>"
	html += "<svg"
	html += " viewbox=\"0 0 " + strconv.Itoa(chart.ViewBox.Width) + " " + strconv.Itoa(chart.ViewBox.Height) + "\""
	html += " width=\"" + strconv.Itoa(chart.ViewBox.Width) + "\""
	html += " height=\"" + strconv.Itoa(chart.ViewBox.Height) + "\""
	html += ">"

	min_value, max_value := calculateChartDatasetMinMax(chart.Dataset, chart.Properties)

	for p, property := range chart.Properties {

		html += "<g"
		html += " data-property=\"" + property + "\""
		html += " data-palette=\"" + strconv.Itoa(p+1) + "\""
		html += ">"

		path, texts := renderLineChartDataset(
			chart.Dataset,
			chart.ViewBox.Width,
			chart.ViewBox.Height,
			min_value,
			max_value,
			property,
		)

		html += "<path"
		html += " d=\"" + path.GetAttribute("d") + "\""
		html += "/>"

		for _, text := range texts {

			html += "<text"
			html += " text-anchor=\"" + text.GetAttribute("text-anchor") + "\""
			html += " dominant-baseline=\"" + text.GetAttribute("dominant-baseline") + "\""
			html += " x=\"" + text.GetAttribute("x") + "\""
			html += " y=\"" + text.GetAttribute("y") + "\""
			html += ">"
			html += text.InnerHTML
			html += "</text>"

		}

		html += "</g>"

	}

	html += "</svg>"
	html += "<figcaption>"

	for p, property := range chart.Properties {
		html += "<label data-property=\"" + property + "\">" + chart.Labels[p] + "</label>"
	}

	html += "</figcaption>"
	html += "</figure>"

	return html

}

func (chart *LineChart) Unmount() bool {

	if chart.Component.Element != nil {
		chart.Component.Element.RemoveEventListener("mousemove", nil)
	}

	return true

}
