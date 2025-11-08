package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "time"

func main() {

	console := console.GetConsole()
	document := dom.GetDocument()

	button := ui.ToButton(document.QuerySelector("button"))
	button.Mount()

	button.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("button click event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	input := ui.ToInput(document.QuerySelector("input"))
	input.Mount()

	input.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("input change-value event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	selekt := ui.ToSelect(document.QuerySelector("select"))
	selekt.Mount()

	selekt.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("select change-value event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	textarea := ui.ToTextarea(document.QuerySelector("textarea"))
	textarea.Mount()

	textarea.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("textarea change-value event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
