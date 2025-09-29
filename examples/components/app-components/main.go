package main

// import "github.com/cookiengineer/gooey/bindings/console"
// import "github.com/cookiengineer/gooey/bindings/dom"
// import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/components"
// import "github.com/cookiengineer/gooey/components/content"
// import "github.com/cookiengineer/gooey/components/layout"
// import "github.com/cookiengineer/gooey/components/ui"
import app_components "example/components"
import "fmt"
import "time"

func main() {

	document := components.NewDocument()

	// XXX: This is how to use Gooey Components
	// content.RegisterTo(document)
	// layout.RegisterTo(document)
	// ui.RegisterTo(document)
	// app.RegisterTo(document)

	app_components.RegisterTo(document)

	fmt.Println(document)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
