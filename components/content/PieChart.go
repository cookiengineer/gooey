//go:build wasm

package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/components/utils"
import "github.com/cookiengineer/gooey/interfaces"
import "strconv"
import "strings"

type PieChart struct {
	Name       string                `json:"name"`
	Disabled   bool                  `json:"disabled"`
	Labels     []string              `json:"labels"`
	Properties []string              `json:"properties"`
	Types      []string              `json:"types"`
	Data       *data.Data            `json:"data"`
	Component  *components.Component `json:"component"`
	ViewBox    struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"viewbox"`
}

func NewPieChart(name string, labels []string, properties []string, types []string) PieChart {

	var chart PieChart

	element   := dom.Document.CreateElement("figure")
	component := components.NewComponent(element)

	chart.Data       = &data.Data{}
	chart.Component  = &component
	chart.Name       = strings.TrimSpace(strings.ToLower(name))
	chart.Labels     = make([]string, 0)
	chart.Properties = make([]string, 0)
	chart.Types      = make([]string, 0)

	chart.ViewBox.Width  = 512
	chart.ViewBox.Height = 512

	chart.SetLabelsAndPropertiesAndTypes(labels, properties, types)
	chart.Mount()
	chart.Render()

	return chart

}

func ToPieChart(element *dom.Element) *PieChart {

	var chart PieChart

	component := components.NewComponent(element)

	chart.Data       = &data.Data{}
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

func (chart *PieChart) Disable() bool {

	chart.Disabled = true
	chart.Render()

	return true

}

func (chart *PieChart) Enable() bool {

	chart.Disabled = false
	chart.Render()

	return true

}

func (chart *PieChart) Mount() bool {

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
					chart.ViewBox.Height = int(512)
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

			if len(chart.Labels) == 0 && len(chart.Properties) == 0 && len(chart.Types) == 0 {

				elements   := datalist.QuerySelectorAll("data")
				labels     := make([]string, 0)
				properties := make([]string, 0)
				types      := make(map[string]string)
				values     := make(map[string]string)

				for _, element := range elements {

					property := element.GetAttribute("data-property")
					label    := strings.TrimSpace(element.TextContent)
					typ      := element.GetAttribute("data-type")
					val      := element.GetAttribute("value")

					types[property]  = typ
					values[property] = val

					labels     = append(labels, label)
					properties = append(properties, property)

				}

				if len(values) > 0 && len(values) == len(types) {

					tmp := data.ParseData(values, types)
					chart.Data = &tmp

				}

				chart.Labels     = labels
				chart.Properties = properties

				tmp2 := make([]string, 0)

				for _, property := range properties {
					tmp2 = append(tmp2, types[property])
				}
				
				chart.Types = tmp2

			}

		}

		chart.Component.Element.AddEventListener("mousemove", dom.ToEventListener(func(event *dom.Event) {

			if chart.Disabled == false && event.Target != nil {

				if event.Target.TagName == "LABEL" {

					property := event.Target.GetAttribute("data-property")

					if property != "" {

						svg    := chart.Component.Element.QuerySelector("svg")
						layers := chart.Component.Element.QuerySelectorAll("svg g")

						// var foreground *dom.Element = nil
						// background := make([]*dom.Element, 0)

						for _, layer := range layers {

							if layer.GetAttribute("data-property") == property {
								layer.SetAttribute("data-state", "active")
								// foreground = layer
							} else {
								layer.RemoveAttribute("data-state")
								// background = append(background, layer)
							}

						}

						if svg != nil {
							// svg.ReplaceChildren(background)
							// svg.Append(foreground)
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

func (chart *PieChart) Query(query string) interfaces.Component {

	if chart.Component.Element != nil {

		if utils.MatchesQuery(chart.Component.Element, query) == true {
			return chart
		}

	}

	return nil

}

func (chart *PieChart) Render() *dom.Element {

	if chart.Component.Element != nil {

		chart.Component.Element.SetAttribute("data-name", chart.Name)
		chart.Component.Element.SetAttribute("data-type", "pie-chart")

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

			if len(*chart.Data) > 0 {

				min_value := int64(0)
				max_value := sumChartData(chart.Data, chart.Properties)
				layers    := make([]*dom.Element, len(chart.Properties))
				offset    := 0.0

				for p, property := range chart.Properties {

					layer := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "g")
					layer.SetAttribute("data-property", property)
					layer.SetAttribute("data-palette", strconv.Itoa(p+1))

					// TODO
					// layer.SetAttribute("transform", "rotate(-180 256 256)")

					path, text, percentage := renderPieChartData(
						chart.Data,
						chart.ViewBox.Width,
						chart.ViewBox.Height,
						min_value,
						max_value,
						property,
						offset,
					)

					if path != nil && text != nil {
						layer.Append(path)
						layer.Append(text)
						offset += percentage
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

func (chart *PieChart) SetData(data data.Data) bool {

	chart.Data = &data

	return true

}

func (chart *PieChart) SetLabelsAndPropertiesAndTypes(labels []string, properties []string, types []string) bool {

	var result bool

	if len(labels) == len(properties) && len(labels) == len(types) {

		chart.Labels     = labels
		chart.Properties = properties
		chart.Types      = types

		result = true

	}

	return result

}

func (chart *PieChart) String() string {

	html := "<figure"

	if chart.Name != "" {
		html += " data-name=\"" + chart.Name + "\""
	}

	html += " data-type=\"line-chart\""
	html += ">"

	html += "<datalist>"

	values, _ := chart.Data.String()

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

	min_value := int64(0)
	max_value := sumChartData(chart.Data, chart.Properties)
	offset    := 0.0

	for p, property := range chart.Properties {

		html += "<g"
		html += " data-property=\"" + property + "\""
		html += " data-palette=\"" + strconv.Itoa(p+1) + "\""
		html += ">"

		path, text, percentage := renderPieChartData(
			chart.Data,
			chart.ViewBox.Width,
			chart.ViewBox.Height,
			min_value,
			max_value,
			property,
			offset,
		)

		html += "<path"
		html += " d=\"" + path.GetAttribute("d") + "\""
		html += "/>"

		html += "<text"
		html += " text-anchor=\"" + text.GetAttribute("text-anchor") + "\""
		html += " dominant-baseline=\"" + text.GetAttribute("dominant-baseline") + "\""
		html += " x=\"" + text.GetAttribute("x") + "\""
		html += " y=\"" + text.GetAttribute("y") + "\""
		html += ">"
		html += text.InnerHTML
		html += "</text>"

		html += "</g>"

		offset += percentage

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

func (chart *PieChart) Unmount() bool {

	if chart.Component.Element != nil {
		chart.Component.Element.RemoveEventListener("mousemove", nil)
	}

	return true

}
