package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "time"

func main() {

	layout_dialog := layout.ToDialog(bindings.Document.QuerySelector("dialog"))
	layout_dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]string) {

		action, ok := attributes["action"]

		if ok == true {

			if action == "cancel" {
				layout_dialog.Hide()
			} else if action == "save" {
				console.Log("Saved Settings!")
				layout_dialog.Hide()
			}

		}

	}, false))

	layout_header := layout.ToHeader(bindings.Document.QuerySelector("header"))
	layout_header.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("header action event")
		console.Log(attributes)
		console.GroupEnd("header action event")

		action, ok := attributes["action"]

		if ok == true {

			if action == "settings" {
				layout_dialog.Show()
			}

		}

	}, false))

	layout_header.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("header change-view event")
		console.Log(attributes)
		console.GroupEnd("header change-view event")

	}, false))

	layout_footer := layout.ToFooter(bindings.Document.QuerySelector("footer"))
	layout_footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("footer action event")
		console.Log(attributes)
		console.GroupEnd("footer action event")

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
