package main

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "github.com/cookiengineer/gooey/components/ui"
import "github.com/cookiengineer/gooey/interfaces"
import "example/controllers"
import "time"

func main() {

	main := app.NewMain()

	// Register Gooey Components
	content.RegisterTo(main.Document)
	layout.RegisterTo(main.Document)
	ui.RegisterTo(main.Document)

	// Register App Controllers
	main.RegisterController("settings", func(main *app.Main, view *app.View) interfaces.Controller {
		return controllers.NewSettings(main, view)
	})
	main.RegisterController("tasks", func(main *app.Main, view *app.View) interfaces.Controller {
		return controllers.NewTasks(main, view)
	})

	// Start the App
	main.Mount()
	main.Render()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
