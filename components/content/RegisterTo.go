package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/interfaces"

func RegisterTo(document *components.Document) {

	document.Register("fieldset", components.WrapComponent(ToFieldset))
	document.Register("figure", func(element *dom.Element) interfaces.Component {

		typ := element.GetAttribute("data-type")

		if typ == "line-chart" {
			return ToLineChart(element)
		} else if typ == "pie-chart" {
			return ToPieChart(element)
		} else {
			return components.ToComponent(element)
		}

	})
	document.Register("table", components.WrapComponent(ToTable))

}
