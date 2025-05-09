package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/ui"
import "time"

func main() {

	ui_button := ui.ToButton(dom.Document.QuerySelector("button"))
	ui_button.Component.AddEventListener("click", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("button click event")
		console.Log(attributes)
		console.GroupEnd("button click event")

	}, false))

	ui_input := ui.ToInput(dom.Document.QuerySelector("input"))
	ui_input.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("input change-value event")
		console.Log(attributes)
		console.GroupEnd("input change-value event")

	}, false))

	ui_select := ui.ToSelect(dom.Document.QuerySelector("select"))
	ui_select.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("select change-value event")
		console.Log(attributes)
		console.GroupEnd("select change-value event")

	}, false))

	ui_textarea := ui.ToTextarea(dom.Document.QuerySelector("textarea"))
	ui_textarea.Component.AddEventListener("change-value", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("textarea change-value event")
		console.Log(attributes)
		console.GroupEnd("textarea change-value event")

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
