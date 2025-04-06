package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/location"
import "github.com/cookiengineer/gooey/bindings/timers"
import "time"

func main() {

	element := bindings.Document.QuerySelector("#location")

	tmp := location.Location.Href

	if tmp != "" {
		element.SetInnerHTML("This page is located at \"" + tmp + "\"!")
	}

	timers.SetTimeout(func() {
		console.Log(location.Location)
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
