package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/navigator"
import "encoding/json"
import "time"

func main() {

	element := dom.Document.QuerySelector("#navigator")
	details, err := json.MarshalIndent(navigator.Navigator, "", "\t")

	if err == nil {
		element.SetInnerHTML(string(details))
	}

	console.Log(navigator.Navigator)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
