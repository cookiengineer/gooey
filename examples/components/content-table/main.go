package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "time"

func main() {

	table := content.ToTable(bindings.Document.QuerySelector("table"))
	table.Component.AddEventListener("change-select", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("table change-select event")
		console.Log(attributes)
		console.GroupEnd("table change-select event")

	}, false))
	table.Component.AddEventListener("change-sort", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("table change-sort event")
		console.Log(attributes)
		console.GroupEnd("table change-sort event")

	}, false))

	table.Disable()

	go func() {

		time.Sleep(1 * time.Second)
		table.Enable()

	}()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
