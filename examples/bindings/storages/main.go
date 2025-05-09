package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/storages"
import "github.com/cookiengineer/gooey/bindings/timers"
import "strconv"
import "time"

func main() {

	element := dom.Document.QuerySelector("#seconds")

	var seconds int = 0

	tmp := storages.SessionStorage.GetItemInt("seconds")

	if tmp > 0 {
		seconds = tmp
	}

	timers.SetInterval(func() {
		seconds++
		storages.SessionStorage.SetItem("seconds", seconds)
		element.SetInnerHTML("This page has been running for " + strconv.Itoa(seconds) + " seconds!")
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
