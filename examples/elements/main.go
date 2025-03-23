package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/dom"
import "strconv"
import "time"

func main() {

	var count int = 0

	listener := dom.ToEventListener(func(event dom.Event) {

		target := event.Target

		if target.Id == "clickable" {

			if target.ClassName == "active" {
				count++
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
				target.SetClassName("")
			} else {
				count++
				target.SetClassName("active")
				target.SetInnerHTML("Click me again! (" + strconv.Itoa(count) + ")")
			}

		}

	})

	bindings.Document.AddEventListener("click", listener)

	for true {

		if count > 10 {

			bindings.Document.RemoveEventListener("click", &listener)

			clickable := bindings.Document.QuerySelector("#clickable")
			clickable.SetClassName("disabled")
			clickable.SetInnerHTML("Stop clicking me!")

			count = 0

		}

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
