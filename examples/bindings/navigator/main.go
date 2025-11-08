package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/navigator"
import "encoding/json"
import "time"

func main() {

	console := console.GetConsole()
	document := dom.GetDocument()
	navigator := navigator.GetNavigator()

	element := document.QuerySelector("#navigator")
	details, err := json.MarshalIndent(navigator, "", "\t")

	if err == nil {
		element.SetInnerHTML(string(details))
	}

	console.Log(navigator)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
