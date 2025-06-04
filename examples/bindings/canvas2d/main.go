package main

import "github.com/cookiengineer/gooey/bindings/animations"
import "github.com/cookiengineer/gooey/bindings/canvas2d"
import "github.com/cookiengineer/gooey/bindings/dom"
import "time"

func main() {

	element := dom.Document.QuerySelector("canvas")
	canvas  := canvas2d.ToCanvas(element)
	image := canvas2d.NewImage(42, 42, "/images/gooey.png")

	animations.RequestAnimationFrame(func(timestamp float64) {

		context := canvas.GetContext()

		context.BeginPath()
		context.SetFillStyleColor("#ff0000")
		context.FillRect(10, 10, 20, 20)
		context.ClosePath()

		context.BeginPath()
		context.SetStrokeStyleColor("#00ff00")
		context.MoveTo(30, 20)
		context.BezierCurveTo(
			120,  30,
			 30, 120,
			120, 120,
		)
		context.Stroke()
		context.ClosePath()

		context.BeginPath()
		context.Rect(120, 110, 20, 20)
		context.SetStrokeStyleColor("#0000ff")
		context.Stroke()
		context.ClosePath()

		context.DrawImage(
			&image,
			0, 0, 42, 42,
			13, 37, 42, 42,
		)

	})

	for true {

		// Do Nothing
		time.Sleep(100 * time.Millisecond)

	}

}
