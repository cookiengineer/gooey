//go:build wasm

package canvas2d

import "syscall/js"

type Context struct {
	Font                     string             `json:"font"`
	GlobalAlpha              float64            `json:"globalAlpha"`
	GlobalCompositeOperation CompositeOperation `json:"globalCompositeOperation"`
	LetterSpacing            string             `json:"letterSpacing"`
	LineCap                  LineCap            `json:"lineCap"`
	LineJoin                 LineJoin           `json:"lineJoin"`
	LineDash                 []float64          `json:"lineDash"`
	LineDashOffset           float64            `json:"lineDashOffset"`
	LineWidth                float64            `json:"lineWidth"`
	TextAlign                TextAlign          `json:"textAlign"`
	TextBaseline             TextBaseline       `json:"textBaseline"`
	WordSpacing              string             `json:"wordSpacing"`
	Value                    *js.Value          `json:"value"`
}

func ToContext(value js.Value) *Context {

	var context Context

	context.Font = string(value.Get("font").String())
	context.GlobalAlpha = float64(value.Get("globalAlpha").Float())
	context.GlobalCompositeOperation = CompositeOperation(value.Get("globalCompositeOperation").String())
	context.LetterSpacing = string(value.Get("letterSpacing").String())
	context.LineCap = LineCap(value.Get("lineCap").String())
	context.LineDash = make([]float64, 0)
	context.LineDashOffset = float64(value.Get("lineDashOffset").Float())
	context.LineJoin = LineJoin(value.Get("lineJoin").String())
	context.LineWidth = float64(value.Get("lineWidth").Float())
	context.TextAlign = TextAlign(value.Get("textAlign").String())
	context.TextBaseline = TextBaseline(value.Get("textBaseline").String())
	context.WordSpacing = string(value.Get("wordSpacing").String())
	context.Value = &value

	line_dash := value.Call("getLineDash")

	if !line_dash.IsNull() && !line_dash.IsUndefined() {

		for l := 0; l < line_dash.Length(); l++ {
			context.LineDash = append(context.LineDash, line_dash.Index(l).Float())
		}

	}

	return &context

}

func (context *Context) IsContextLost() bool {

	var result bool

	value := context.Value.Call("isContextLost")

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Bool()
	}

	return result

}

func (context *Context) Arc(x int, y int, radius int, start_angle float64, end_angle float64, counterclockwise bool) {
	context.Value.Call("arc", x, y, radius, start_angle, end_angle, counterclockwise)
}

func (context *Context) ArcTo(x1 int, y1 int, x2 int, y2 int, radius int) {
	context.Value.Call("arcTo", x1, y1, x2, y2, radius)
}

func (context *Context) BeginPath() {
	context.Value.Call("beginPath")
}

func (context *Context) BezierCurveTo(cp1x int, cp1y int, cp2x int, cp2y int, x int, y int) {
	context.Value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (context *Context) ClearRect(x int, y int, width int, height int) {
	context.Value.Call("clearRect", x, y, width, height)
}

func (context *Context) Clip(fillrule FillRule) {
	context.Value.Call("clip", string(fillrule))
}

func (context *Context) ClosePath() {
	context.Value.Call("closePath")
}

func (context *Context) DrawImage(image *Image, sx int, sy int, swidth int, sheight int, dx int, dy int, dwidth int, dheight int) {
	context.Value.Call("drawImage", *image.Value, sx, sy, swidth, sheight, dx, dy, dwidth, dheight)
}

func (context *Context) Ellipse(x int, y int, radius_x int, radius_y int, rotation float64, start_angle float64, end_angle float64, counterclockwise bool) {
	context.Value.Call("ellipse", x, y, radius_x, radius_y, rotation, start_angle, end_angle, counterclockwise)
}

func (context *Context) Fill() {
	context.Value.Call("fill")
}

func (context *Context) FillRect(x int, y int, width int, height int) {
	context.Value.Call("fillRect", x, y, width, height)
}

func (context *Context) FillText(text string, x int, y int) {
	context.Value.Call("fillText", text, x, y)
}

func (context *Context) GetLineDash() []float64 {
	return context.LineDash
}

func (context *Context) IsPointInPath(x int, y int) bool {

	var result bool

	tmp := context.Value.Call("isPointInPath", x, y)

	if !tmp.IsNull() && !tmp.IsUndefined() && tmp.Bool() == true {
		result = true
	}

	return result

}

func (context *Context) IsPointInStroke(x int, y int) bool {

	var result bool

	tmp := context.Value.Call("isPointInStroke", x, y)

	if !tmp.IsNull() && !tmp.IsUndefined() && tmp.Bool() == true {
		result = true
	}

	return result

}

func (context *Context) LineTo(x int, y int) {
	context.Value.Call("lineTo", x, y)
}

func (context *Context) MeasureText(text string) *TextMetrics {

	var metrics *TextMetrics

	value := context.Value.Call("measureText", text)

	if !value.IsNull() && !value.IsUndefined() {
		metrics = ToTextMetrics(value)
	}

	return metrics

}

func (context *Context) MoveTo(x int, y int) {
	context.Value.Call("moveTo", x, y)
}

func (context *Context) Rect(x int, y int, width int, height int) {
	context.Value.Call("rect", x, y, width, height)
}

func (context *Context) Reset() {
	context.Value.Call("reset")
}

func (context *Context) Restore() {
	context.Value.Call("restore")
}

func (context *Context) Rotate(angle float64) {
	context.Value.Call("rotate", angle)
}

func (context *Context) Save() {
	context.Value.Call("save")
}

func (context *Context) Scale(x float64, y float64) {
	context.Value.Call("scale", x, y)
}

func (context *Context) SetFillStyleColor(color string) {
	// TODO: Validate CSS color syntax
	context.Value.Set("fillStyle", string(color))
}

func (context *Context) SetFont(font string) {
	// TODO: Validate CSS font syntax
	context.Value.Set("font", string(font))
	context.Font = font
}

func (context *Context) SetGlobalAlpha(alpha float64) {

	if alpha >= 0.0 && alpha <= 1.0 {
		context.Value.Set("globalAlpha", alpha)
		context.GlobalAlpha = alpha
	}

}

func (context *Context) SetGlobalCompositeOperation(operation CompositeOperation) {
	context.Value.Set("globalCompositeOperation", string(operation))
	context.GlobalCompositeOperation = operation
}

func (context *Context) SetLetterSpacing(css_length string) {
	// TODO: Validate CSS lengths (with units?)
	context.Value.Set("letterSpacing", string(css_length))
	context.WordSpacing = css_length
}

func (context *Context) SetLineCap(linecap LineCap) {
	context.Value.Set("lineCap", string(linecap))
	context.LineCap = linecap
}

func (context *Context) SetLineJoin(linejoin LineJoin) {
	context.Value.Set("lineJoin", string(linejoin))
	context.LineJoin = linejoin
}

func (context *Context) SetLineDash(linedash []float64) {

	wrapped_linedash := js.Global().Get("Array").New(len(linedash))

	for l := 0; l < len(linedash); l++ {
		wrapped_linedash.SetIndex(l, linedash[l])
	}

	context.Value.Call("setLineDash", wrapped_linedash)
	context.LineDash = linedash

}

func (context *Context) SetLineDashOffset(offset float64) {
	context.Value.Set("lineDashOffset", float64(offset))
	context.LineDashOffset = offset
}

func (context *Context) SetStrokeStyleColor(color string) {
	// TODO: Validate CSS color syntax
	context.Value.Set("strokeStyle", string(color))
}

func (context *Context) SetTextAlign(align TextAlign) {
	context.Value.Set("textAlign", string(align))
	context.TextAlign = align
}

func (context *Context) SetTextBaseline(baseline TextBaseline) {
	context.Value.Set("textBaseline", string(baseline))
	context.TextBaseline = baseline
}

func (context *Context) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) {
	context.Value.Call("setTransform", a, b, c, d, e, f)
}

func (context *Context) SetWordSpacing(css_length string) {
	// TODO: Validate CSS lengths (with units?)
	context.Value.Set("wordSpacing", string(css_length))
	context.WordSpacing = css_length
}

func (context *Context) Stroke() {
	context.Value.Call("stroke")
}

func (context *Context) StrokeRect(x int, y int, width int, height int) {
	context.Value.Call("strokeRect", x, y, width, height)
}

func (context *Context) StrokeText(text string, x int, y int) {
	context.Value.Call("strokeText", text, x, y)
}

func (context *Context) Transform(a float64, b float64, c float64, d float64, e float64, f float64) {
	context.Value.Call("transform", a, b, c, d, e, f)
}

func (context *Context) Translate(x int, y int) {
	context.Value.Call("translate", x, y)
}
