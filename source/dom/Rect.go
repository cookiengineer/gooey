//go:build wasm

package dom

import "syscall/js"

type Rect struct {
	X      int       `json:"x"`
	Y      int       `json:"y"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Top    int       `json:"top"`
	Right  int       `json:"right"`
	Bottom int       `json:"bottom"`
	Left   int       `json:"left"`
	Value  *js.Value `json:"value"`
}

func ToRect(value js.Value) Rect {

	var rect Rect

	rect.X = value.Get("x").Int()
	rect.Y = value.Get("y").Int()
	rect.Width = value.Get("width").Int()
	rect.Height = value.Get("height").Int()
	rect.Top = value.Get("top").Int()
	rect.Right = value.Get("right").Int()
	rect.Bottom = value.Get("bottom").Int()
	rect.Left = value.Get("left").Int()
	rect.Value = &value

	return rect

}
