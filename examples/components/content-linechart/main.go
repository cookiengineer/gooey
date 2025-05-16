package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/content"
import "time"

func main() {

	chart := content.ToLineChart(dom.Document.QuerySelector("figure[data-type=\"line-chart\"]"))

	chart.Disable()

	go func() {

		time.Sleep(500 * time.Millisecond)
		chart.Enable()

	}()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
