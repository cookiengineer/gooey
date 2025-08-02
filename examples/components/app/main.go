package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "example/controllers"
import "time"

func main() {

	element := dom.Document.QuerySelector("main")

	main                := app.NewMain(element)
	controller_tasks    := controllers.NewTasks(main)
	controller_settings := controllers.NewSettings(main)

	main.SetView(controller_tasks.View)
	main.SetView(controller_settings.View)

	view := element.GetAttribute("data-view")

	if view != "" {
		main.ChangeView(view)
	}

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
