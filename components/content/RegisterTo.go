package content

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("fieldset",                         components.Wrap(ToFieldset))
	document.Register("figure[data-type=\"line-chart\"]", components.Wrap(ToLineChart))
	document.Register("figure[data-type=\"pie-chart\"]",  components.Wrap(ToPieChart))
	document.Register("table",                            components.Wrap(ToTable))

}
