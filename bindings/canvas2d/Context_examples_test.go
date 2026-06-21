//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"

func Example() {

	// import "github.com/cookiengineer/gooey/bindings/dom"

	document := dom.GetDocument()
	element  := document.QuerySelector("canvas")
	canvas   := ToCanvas(element)
	image    := NewImage(42, 42, "/images/gooey.png")
	context  := canvas.GetContext()

	context.BeginPath()
	context.SetFillStyleColor("#ff0000")
	context.FillRect(10, 10, 20, 20)
	context.ClosePath()

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
