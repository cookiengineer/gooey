package main

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/data"
import "time"

func main() {

	performance_chart := content.ToLineChart(dom.Document.QuerySelector("figure[data-name=\"performance\"]"))
	performance_chart.Mount()

	performance_chart.Disable()

	go func() {

		time.Sleep(500 * time.Millisecond)
		performance_chart.Enable()
		performance_chart.Render()

	}()

	activity_chart := content.ToLineChart(dom.Document.QuerySelector("figure[data-name=\"activity\"]"))
	activity_chart.Mount()

	activity_dataset := data.ToDataset([]data.Data{
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
		data.Data(map[string]any{"mouse-x": 0, "mouse-y": 0}),
	})
	activity_chart.SetDataset(activity_dataset)

	var index int = 0

	listener := dom.ToEventListener(func(event *dom.Event) {

		index++

		if index >= 10 {
			index = 0
		}

		activity_chart.Dataset.Set(index, data.Data(map[string]any{
			"mouse-x": int(event.Value.Get("clientX").Int()),
			"mouse-y": int(event.Value.Get("clientY").Int()),
		}))

		activity_chart.Render()

	})

	dom.Document.AddEventListener("mousemove", listener)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
