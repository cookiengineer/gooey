package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/data"
import "time"

func main() {

	chart1 := content.ToPieChart(dom.Document.QuerySelector("figure[data-name=\"integers\"]"))
	chart2 := content.ToPieChart(dom.Document.QuerySelector("figure[data-name=\"floats\"]"))
	chart3 := content.ToPieChart(dom.Document.QuerySelector("figure[data-name=\"percentages\"]"))

	chart1.Disable()
	chart2.Disable()
	chart3.Disable()

	go func() {

		time.Sleep(500 * time.Millisecond)

		chart1.Enable()
		chart2.Enable()
		chart3.Enable()

	}()

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
