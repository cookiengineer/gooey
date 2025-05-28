package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"

type Canvas struct {
	Width   uint         `json:"width"`
	Height  uint         `json:"height"`
	Element *dom.Element `json:"element"`
	Context *Context     `json:"context"`
}

func ToCanvas(element *dom.Element) *Canvas {

	if element.TagName == "CANVAS" {

		var canvas Canvas

		canvas.Width = uint(element.Value.Get("width").Int())
		canvas.Height = uint(element.Value.Get("height").Int())
		canvas.Element = element

		tmp := element.Value.Call("getContext", "2d")

		if !tmp.IsNull() && !tmp.IsUndefined() {
			canvas.Context = ToContext(tmp)
		}

		return &canvas

	}

	return nil

}

func (canvas *Canvas) GetContext() *Context {
	return canvas.Context
}

func (canvas *Canvas) SetWidth(width uint) {

	canvas.Element.Value.Set("width", width)
	canvas.Width = width

}

func (canvas *Canvas) SetHeight(height uint) {

	canvas.Element.Value.Set("height", height)
	canvas.Height = height

}
