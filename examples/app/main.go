package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components/app"
import "example/views"
import "time"

func main() {

	element := bindings.Document.QuerySelector("main")

	main := app.Main{}
	main.Init(element)
	main.SetView(views.NewTasks(&main))
	main.ChangeView("tasks")

	console.Log(main)

	console.Group("Header Component")
	console.Log(main.Header)
	main.Header.Render()
	console.Log(main.Header.String())
	console.GroupEnd("Header Component")

	// console.Group("Footer Component")
	// console.Log(main.Footer)
	// console.Log(main.Footer.Render())
	// console.Log(main.Footer.String())
	// console.GroupEnd("Footer Component")

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
