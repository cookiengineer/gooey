package main

import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "encoding/json"
import "time"

func main() {

	pre := dom.Document.QuerySelector("pre")

	fieldset := content.ToFieldset(dom.Document.QuerySelector("fieldset"))
	fieldset.Component.AddEventListener("change-field", components.ToEventListener(func(event string, attributes map[string]any) {

		console.Group("fieldset change-field event")
		console.Log(attributes)
		console.GroupEnd("fieldset change-field event")

	}, false))

	footer := layout.ToFooter(dom.Document.QuerySelector("footer"))
	footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		action, ok := attributes["action"]

		if ok == true {

			if action == "cancel" {

				fieldset.Reset()

			} else if action == "confirm" {

				bytes, err := json.MarshalIndent(map[string]any{
					"type":    fieldset.ValueOf("type").String(),
					"name":    fieldset.ValueOf("name").String(),
					"email":   fieldset.ValueOf("email").String(),
					"message": fieldset.ValueOf("message").String(),
					"rating":  fieldset.ValueOf("rating").Int(),
					"amount":  fieldset.ValueOf("amount").Int(),
					"tos":     fieldset.ValueOf("tos").Bool(),
				}, "", "\t")

				if pre != nil && err == nil {
					pre.SetInnerHTML(string(bytes))
				}

			}

		}

	}, false))

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
