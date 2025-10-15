package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "time"

func main() {

	button := ui.ToButton(dom.Document.QuerySelector("button"))
	button.Mount()

	button.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("button click event")
		console.Log(attributes)
		console.GroupEnd("button click event")

	}, false))

	input := ui.ToInput(dom.Document.QuerySelector("input"))
	input.Mount()

	input.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("input change-value event")
		console.Log(attributes)
		console.GroupEnd("input change-value event")

	}, false))

	selekt := ui.ToSelect(dom.Document.QuerySelector("select"))
	selekt.Mount()

	selekt.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("select change-value event")
		console.Log(attributes)
		console.GroupEnd("select change-value event")

	}, false))

	textarea := ui.ToTextarea(dom.Document.QuerySelector("textarea"))
	textarea.Mount()

	textarea.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("textarea change-value event")
		console.Log(attributes)
		console.GroupEnd("textarea change-value event")

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
