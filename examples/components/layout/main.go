package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/layout"
import "time"

func main() {

	console := console.GetConsole()
	document := dom.GetDocument()

	dialog := layout.ToDialog(document.QuerySelector("dialog"))
	dialog.Mount()

	dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "cancel" {

				console.Warn("Cancelled Settings!")
				dialog.Hide()

			} else if action == "save" {

				console.Info("Saved Settings!")
				dialog.Hide()

			}

		}

	}, false))

	header := layout.ToHeader(document.QuerySelector("header"))
	aside := layout.ToAside(document.QuerySelector("aside"))

	header.Mount()
	aside.Mount()

	header.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("header action event")
		console.Log(attributes)
		console.GroupEnd()

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "settings" {
				dialog.Show()
			}

		}

	}, false))
	header.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]any) {

		name, ok := attributes["name"].(string)

		if ok == true {
			aside.ChangeView(name)
		}

		console.Group("header change-view event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	aside.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("aside action event")
		console.Log(attributes)
		console.GroupEnd()

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "settings" {
				dialog.Show()
			}

		}

	}, false))
	aside.Component.AddEventListener("change-view", components.ToEventListener(func(event string, attributes map[string]any) {

		name, ok := attributes["name"].(string)

		if ok == true {
			header.ChangeView(name)
		}

		console.Group("aside change-view event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	footer := layout.ToFooter(document.QuerySelector("footer"))
	footer.Mount()

	footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("footer action event")
		console.Log(attributes)
		console.GroupEnd()

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
