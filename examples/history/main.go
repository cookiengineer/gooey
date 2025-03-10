package main

import "gooey"
import "gooey/console"
import "gooey/history"
import "gooey/location"
import "gooey/timers"
import "time"

// Customizable History States
type HistoryState struct {
	Category uint `json:"category"`
	Page     int  `json:"page"`
}

func main() {

	element := gooey.Document.QuerySelector("#location")

	history.History.PushState(history.State{
	})

	// TODO: Call history.PushState()
	// TODO: Call history.PopState()
	// TODO: Call history.Forward()
	// TODO: Call history.Back()
	// TODO: Call history.Go()


	tmp := location.Location.Href

	if tmp != "" {
		element.SetInnerHTML("This page is located at \"" + tmp + "\"!")
	}

	timers.SetTimeout(func() {
		console.Inspect(location.Location)
	}, 1000)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
