package views

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterView("settings", func(element *dom.Element) interfaces.View {
		return ToSettings(element)
	})

}
