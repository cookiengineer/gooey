package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
import "time"

func main() {

	table1 := content.ToTable(dom.Document.QuerySelector("table[data-name=\"candidates\"]"))
	table2 := content.ToTable(dom.Document.QuerySelector("table[data-name=\"interviews\"]"))

	table1.Mount()
	table2.Mount()

	table1.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		if event == "action" {

			action, ok := attributes["action"].(string)

			if ok == true {

				if action == "accept" {

					indexes, dataset := table1.Selected()

					for d := 0; d < len(dataset); d++ {

						data := dataset[d]
						data["invited"] = true

						table2.Add(data)

					}

					table1.Deselect(indexes)
					table1.Remove(indexes)

					table1.Render()
					table2.Render()

					table2.Component.Element.SetClassName("visible")

				} else if action == "deny" {

					indexes, dataset := table1.Selected()

					for d := 0; d < len(dataset); d++ {

						data := dataset[d]
						data["invited"] = false

						table2.Add(data)

					}

					table1.Deselect(indexes)
					table1.Remove(indexes)

					table1.Render()
					table2.Render()

					table2.Component.Element.SetClassName("visible")

				}

			}

		}

	}, false))

	table1.Disable()
	table2.Disable()

	go func() {

		time.Sleep(500 * time.Millisecond)

		table1.Enable()
		table2.Enable()

		table1.Render()
		table2.Render()

	}()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
