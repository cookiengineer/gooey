package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/location"
import "encoding/json"
import "time"

func main() {

	element := dom.Document.QuerySelector("#location")
	details, err := json.MarshalIndent(location.Location, "", "\t")

	if err == nil {
		element.SetInnerHTML(string(details))
	}

	console.Log(location.Location)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
