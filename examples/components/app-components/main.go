package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import app_components "example/components"
import "example/controllers"
import "example/views"
import "time"

func main() {

	main := app.NewMain()

	// Register Gooey Components
	content.RegisterTo(main.Document)
	layout.RegisterTo(main.Document)
	ui.RegisterTo(main.Document)

	// Register App Components
	app_components.RegisterTo(main.Document)

	// Register App Controllers
	main.RegisterController("settings", func(main *app.Main, view *app.View) interfaces.Controller {
		return controllers.NewSettings(main, view)
	})

	// Register App Views
	main.RegisterView("settings", func(element *dom.Element) interfaces.View {
		return views.ToSettings(element)
	})

	// Start the App
	main.Mount()
	main.Render()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
