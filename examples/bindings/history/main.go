package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/history"
import "encoding/json"
import "time"

func renderEvent(event *history.PopStateEvent) string {

	html := ""
	html += "<li>"

	data1, err1 := json.Marshal(history.GetHistory().State)
	data2, err2 := json.Marshal(event)

	if err1 == nil {
		html += string(data1)
	} else if err2 == nil {
		html += string(data2)
	} else {
		html += "(PopStateEvent)"
	}

	html += "</li>"

	return html

}

func main() {

	document := dom.GetDocument()
	history1 := history.GetHistory()

	list_events    := document.QuerySelector("main ul")
	button_back    := document.QuerySelector("main button[data-action=\"back\"]")
	button_forward := document.QuerySelector("main button[data-action=\"forward\"]")

	history1.AddEventListener(history.ToEventListener(func(event *history.PopStateEvent) {

		html := renderEvent(event)
		list_events.InsertAdjacentHTML("beforeend", html)

		console.Log("popstate event!")
		console.Log(event)

	}))

	button_back.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {
		history1.Back()
	}))

	button_forward.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {
		history1.Forward()
	}))

	history1.PushState(&map[string]any{"page": 1}, "first page",  "/page-1.html")
	history1.PushState(&map[string]any{"page": 2}, "second page", "/page-2.html")
	history1.PushState(&map[string]any{"page": 3}, "third page", "/page-3.html")

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
