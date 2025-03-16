package main

import "github.com/cookiengineer/gooey/pkg"
import "github.com/cookiengineer/gooey/pkg/storages"
import "github.com/cookiengineer/gooey/pkg/timers"
import "strconv"
import "time"

func main() {

	element := gooey.Document.QuerySelector("#seconds")

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
