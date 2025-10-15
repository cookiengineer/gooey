package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "time"

func main() {

	dialog := layout.ToDialog(dom.Document.QuerySelector("dialog"))
	dialog.Mount()

	dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "cancel" {
				dialog.Hide()
			} else if action == "save" {
				console.Log("Saved Settings!")
				dialog.Hide()
			}

		}

	}, false))

	header := layout.ToHeader(dom.Document.QuerySelector("header"))
	header.Mount()

	header.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("header action event")
		console.Log(attributes)
		console.GroupEnd("header action event")

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "settings" {
				dialog.Show()
			}

		}

	}, false))
	header.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("header change-view event")
		console.Log(attributes)
		console.GroupEnd("header change-view event")

	}, false))

	footer := layout.ToFooter(dom.Document.QuerySelector("footer"))
	footer.Mount()

	footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("footer action event")
		console.Log(attributes)
		console.GroupEnd("footer action event")

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
