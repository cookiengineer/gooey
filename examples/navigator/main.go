package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/navigator"
import "time"

func main() {

	console.Log(navigator.Navigator)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
