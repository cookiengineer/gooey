package controllers

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/interfaces"

func RegisterTo(main *app.Main) {

	main.RegisterController("settings", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewSettings(main, view)
	})

	main.RegisterController("tasks", func(main *app.Main, view interfaces.View) interfaces.Controller {
		return NewTasks(main, view)
	})


}
