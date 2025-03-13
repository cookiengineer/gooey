package main

import "gooey"
import "gooey/app"
import "gooey/console"
import "example/views"
import "time"

func main() {

	element := gooey.Document.QuerySelector("main")

	main := app.Main{}
	main.Init(element)

	view := element.GetAttribute("data-view")

	if view == "tasks" {
		main.SetView("tasks", views.NewTasks(&main))
		main.ChangeView("tasks")
	} else if view == "settings" {
		// TODO: Quick Settings example
		// main.SetView("settings", views.NewSettings(&main))
		// main.ChangeView("settings")
	}

	console.Log(main)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
