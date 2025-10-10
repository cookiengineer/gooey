package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/bindings/fetch"
import "encoding/json"
import "time"

func main() {

	element1 := dom.Document.QuerySelector("#fetch-response")
	element2 := dom.Document.QuerySelector("#fetch-error")

	response, err := fetch.Fetch("/api/test", &fetch.Request{
		Method: fetch.MethodGet,
		Mode:   fetch.ModeCORS,
	})

	details1, _ := json.MarshalIndent(response, "", "\t")
	element1.SetInnerHTML(string(details1))

	details2 := ""

	if err != nil {
		details2 = err.Error()
		element2.SetInnerHTML(string(details2))
	} else {
		element2.SetInnerHTML("(nil)")
	}

	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
