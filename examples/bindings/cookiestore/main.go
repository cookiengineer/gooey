package main

import "github.com/cookiengineer/gooey/bindings/cookiestore"
import "github.com/cookiengineer/gooey/bindings/dom"
import "fmt"
import "time"

func main() {

	document := dom.GetDocument()
	store := cookiestore.GetCookieStore()

	get_button := document.QuerySelector("button[data-action=\"get-cookie\"]")

	if get_button != nil {

		get_button.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			input := document.QuerySelector("input[data-name=\"get-name\"]")
			div := document.QuerySelector("div[data-name=\"get-output\"]")

			if input != nil && div != nil {

				go func() {

					cookie, err := store.Get(cookiestore.GetOptions{
						Name: input.Value.Get("value").String(),
					})

					if err == nil {
						div.SetInnerHTML(cookie.Value)
					} else {
						div.SetInnerHTML(err.Error())
					}

				}()

			}

		}))

	}

	set_button := document.QuerySelector("button[data-action=\"set-cookie\"]")

	if set_button != nil {

		set_button.AddEventListener("click", dom.ToEventListener(func(event *dom.Event) {

			input1 := document.QuerySelector("input[data-name=\"set-name\"]")
			input2 := document.QuerySelector("input[data-name=\"set-value\"]")
			div := document.QuerySelector("div[data-name=\"set-output\"]")

			if input1 != nil && input2 != nil {

				go func() {

					name  := input1.Value.Get("value").String()
					value := input2.Value.Get("value").String()
					err   := store.Set(cookiestore.SetOptions{
						Name:  name,
						Value: value,
					})

					if err == nil {
						div.SetInnerHTML(fmt.Sprintf("Cookie \"%s\" was set", name))
					} else {
						div.SetInnerHTML(err.Error())
					}

				}()

			}


		}))

	}

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
