//go:build wasm

package canvas2d

import "syscall/js"

type TextMetrics struct {
	ActualBoundingBoxLeft    int `json:"actualBoundingBoxLeft"`
	ActualBoundingBoxRight   int `json:"actualBoundingBoxRight"`
	ActualBoundingBoxAscent  int `json:"actualBoundingBoxAscent"`
	ActualBoundingBoxDescent int `json:"actualBoundingBoxDescent"`
	FontBoundingBoxAscent    int `json:"fontBoundingBoxAscent"`
	FontBoundingBoxDescent   int `json:"fontBoundingBoxDescent"`
	AlphabeticBaseline       int `json:"alphabeticBaseline"`
	HangingBaseline          int `json:"hangingBaseline"`
	IdeographicBaseline      int `json:"ideographicBaseline"`
}

func ToTextMetrics(value js.Value) *TextMetrics {

	var metrics TextMetrics

	metrics.ActualBoundingBoxLeft = value.Get("actualBoundingBoxLeft").Int()
	metrics.ActualBoundingBoxRight = value.Get("actualBoundingBoxRight").Int()
	metrics.ActualBoundingBoxAscent = value.Get("actualBoundingBoxAscent").Int()
	metrics.ActualBoundingBoxDescent = value.Get("actualBoundingBoxDescent").Int()
	metrics.FontBoundingBoxAscent = value.Get("fontBoundingBoxAscent").Int()
	metrics.FontBoundingBoxDescent = value.Get("fontBoundingBoxDescent").Int()
	metrics.AlphabeticBaseline = value.Get("alphabeticBaseline").Int()
	metrics.HangingBaseline = value.Get("hangingBaseline").Int()
	metrics.IdeographicBaseline = value.Get("ideographicBaseline").Int()

	return &metrics

}
