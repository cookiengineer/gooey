package main

import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/layout"
import "encoding/json"
import "time"

func main() {

	pre := bindings.Document.QuerySelector("pre")

	fieldset := content.ToFieldset(bindings.Document.QuerySelector("fieldset"))
	fieldset.Component.AddEventListener("change-field", components.ToEventListener(func(event string, attributes map[string]string) {

		console.Group("fieldset change-field event")
		console.Log(attributes)
		console.GroupEnd("fieldset change-field event")

	}, false))

	footer := layout.ToFooter(bindings.Document.QuerySelector("footer"))
	footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]string) {

		action, ok := attributes["action"]

		if ok == true {

			if action == "cancel" {

				fieldset.Reset()

			} else if action == "confirm" {

				bytes, err := json.MarshalIndent(map[string]any{
					"name":    fieldset.ValueOf("name").String(),
					"email":   fieldset.ValueOf("email").String(),
					"message": fieldset.ValueOf("message").String(),
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
