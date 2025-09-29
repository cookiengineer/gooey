package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
// import "github.com/cookiengineer/gooey/components/content"
// import "github.com/cookiengineer/gooey/components/layout"
// import "github.com/cookiengineer/gooey/components/ui"
import app_components "example/components"
import "time"

func main() {

	document := components.NewDocument()

	// XXX: This is how to use Gooey Components
	// content.RegisterTo(document)
	// layout.RegisterTo(document)
	// ui.RegisterTo(document)
	// app.RegisterTo(document)

	app_components.RegisterTo(document)

	document.Mount()



	main_component, ok1 := components.Unwrap[*components.Component](document.Query("main"))

	if ok1 == true {

		console.Group("Main Component")
		console.Log(main_component)
		console.GroupEnd("Main Component")

	} else {
		console.Error("Can't typecast to components.Component")
	}

	example_component, ok2 := components.Unwrap[*app_components.Example](document.Query("main > app-example"))

	if ok2 == true {

		console.Group("App Example Component")
		console.Log(example_component)
		console.GroupEnd("App Example Component")

	} else {
		console.Error("Can't typecast to app_components.Example")
	}

	h3_component, ok3 := components.Unwrap[*components.Component](document.Query("main > app-example > h3"))

	if ok3 == true {

		console.Group("H3 Component")
		console.Log(h3_component)
		console.GroupEnd("H3 Component")

	} else {
		console.Error("Can't typecast to components.Component")
	}


	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
