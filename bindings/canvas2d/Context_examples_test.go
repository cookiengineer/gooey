//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"

func ExampleContext_DrawImage() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	image    := NewImage(42, 42, "/images/gooey.png")
	context  := canvas.GetContext()

	context.DrawImage(
		&image,
		0,
		0,
		42,
		42,
		int(canvas.Width)  / 2 - 42/2,
		int(canvas.Height) / 2 - 42/2,
		42,
		42,
	)

}

func ExampleContext_SetFillStyleColor() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()
	color, _ := NewColor("#ff0000")

	context.BeginPath()
	context.SetFillStyleColor(*color)
	context.FillRect(10, 10, 20, 20)
	context.ClosePath()

}

func ExampleContext_SetFillStylePattern() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	image    := NewImage(42, 42, "/images/gooey.png")
	context  := canvas.GetContext()
	pattern  := context.CreatePattern(&image, RepetitionRepeat)

	context.BeginPath()
	context.SetFillStylePattern(pattern)
	context.FillRect(10, 10, 20, 20)
	context.ClosePath()

}

func ExampleContext_SetStrokeStyleColor() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	context  := canvas.GetContext()
	color, _ := NewColor("#ff0000")

	context.BeginPath()
	context.SetStrokeStyleColor(*color)
	context.StrokeRect(10, 10, 20, 20)
	context.ClosePath()

}

func ExampleContext_SetStrokeStylePattern() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	image    := NewImage(42, 42, "/images/gooey.png")
	context  := canvas.GetContext()
	pattern  := context.CreatePattern(&image, RepetitionRepeat)

	context.BeginPath()
	context.SetStrokeStylePattern(pattern)
	context.StrokeRect(10, 10, 20, 20)
	context.ClosePath()

}
