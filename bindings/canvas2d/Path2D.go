//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/dom"
import "fmt"
import "math"
import "strconv"
import "strings"
import "syscall/js"

type Path2D struct {
	Children   []*Path2D       `json:"children"`
	Parameters []float64       `json:"parameters"`
	Transform  *dom.Matrix     `json:"transform"`
	Operation  Path2DOperation `json:"operation"`
	IsClosed   bool            `json:"is_closed"`
	Value      *js.Value       `json:"value"`
}

func NewPath2D() *Path2D {

	var path Path2D

	js_path := js.Global().Get("Path2D").New()

	path.Children = make([]*Path2D, 0)
	path.Parameters = make([]float64, 0)
	path.Transform = dom.NewIdentityMatrix()
	path.IsClosed = false
	path.Value = &js_path

	return &path

}

func (path *Path2D) AddPath(other *Path2D, transform *dom.Matrix) {

	child := &Path2D{
		Children:   other.Children,
		Parameters: other.Parameters,
		Transform:  transform,
		Operation:  Path2DOperationAddPath,
		IsClosed:   other.IsClosed,
	}

	path.Children = append(path.Children, child)

	if transform.Value != nil {
		path.Value.Call("addPath", other.Value, transform.Value)
	} else {
		path.Value.Call("addPath", other.Value)
	}

}

func (path *Path2D) ClosePath() {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: make([]float64, 0),
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationClosePath,
		IsClosed:   true,
	}

	path.Children = append(path.Children, child)
	path.IsClosed = true

	path.Value.Call("closePath")

}

func (path *Path2D) MoveTo(x float64, y float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationMoveTo,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("moveTo", x, y)

}

func (path *Path2D) LineTo(x float64, y float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationLineTo,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("lineTo", x, y)

}

func (path *Path2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{cp1x, cp1y, cp2x, cp2y, x, y},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationBezierCurveTo,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)

}

func (path *Path2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{cpx, cpy, x, y},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationQuadraticCurveTo,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("quadraticCurveTo", cpx, cpy, x, y)

}

func (path *Path2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x1, y1, x2, y2, radius},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationArcTo,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("arcTo", x1, y1, x2, y2, radius)

}

func (path *Path2D) Arc(x float64, y float64, radius float64, start_angle float64, end_angle float64, direction Direction) {

	counterclockwise := false

	if direction == DirectionCounterClockwise {
		counterclockwise = true
	}

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y, radius, start_angle, end_angle, 0},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationArc,
		IsClosed:   false,
	}

	if counterclockwise {
		child.Parameters[5] = 1
	}

	path.Children = append(path.Children, child)

	path.Value.Call("arc", x, y, radius, start_angle, end_angle, counterclockwise)

}

func (path *Path2D) Ellipse(x float64, y float64, radius_x float64, radius_y float64, rotation float64, start_angle float64, end_angle float64, direction Direction) {

	counterclockwise := false

	if direction == DirectionCounterClockwise {
		counterclockwise = true
	}

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y, radius_x, radius_y, rotation, start_angle, end_angle, 0},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationEllipse,
		IsClosed:   false,
	}

	if counterclockwise {
		child.Parameters[7] = 1
	}

	path.Children = append(path.Children, child)

	path.Value.Call("ellipse", x, y, radius_x, radius_y, rotation, start_angle, end_angle, counterclockwise)

}

func (path *Path2D) Rect(x float64, y float64, width float64, height float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y, width, height},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationRect,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("rect", x, y, width, height)

}

func (path *Path2D) RoundRect(x float64, y float64, width float64, height float64, radii float64) {

	child := &Path2D{
		Children:   make([]*Path2D, 0),
		Parameters: []float64{x, y, width, height, radii},
		Transform:  dom.NewIdentityMatrix(),
		Operation:  Path2DOperationRoundRect,
		IsClosed:   false,
	}

	path.Children = append(path.Children, child)

	path.Value.Call("roundRect", x, y, width, height, radii)

}

func (path *Path2D) String() string {

	var builder strings.Builder

	path.write_svg(&builder, path.Transform, 0, 0, 0)

	return builder.String()

}

func (path *Path2D) write_svg(builder *strings.Builder, accumulated *dom.Matrix, depth int, current_x float64, current_y float64) (float64, float64) {

	cx := current_x
	cy := current_y

	for c := 0; c < len(path.Children); c++ {

		child := path.Children[c]
		effective := compose_matrices(accumulated, child.Transform)

		switch child.Operation {

		case Path2DOperationMoveTo:
			cx, cy = transform_point(effective, child.Parameters[0], child.Parameters[1])
			indent(builder, depth)
			builder.WriteString("M ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationLineTo:
			cx, cy = transform_point(effective, child.Parameters[0], child.Parameters[1])
			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationHorizontalLineTo:
			cx, _ = transform_point(effective, child.Parameters[0], cy)
			indent(builder, depth)
			builder.WriteString("H ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString("\n")

		case Path2DOperationVerticalLineTo:
			_, cy = transform_point(effective, cx, child.Parameters[0])
			indent(builder, depth)
			builder.WriteString("V ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationBezierCurveTo:
			cp1x, cp1y := transform_point(effective, child.Parameters[0], child.Parameters[1])
			cp2x, cp2y := transform_point(effective, child.Parameters[2], child.Parameters[3])
			cx, cy = transform_point(effective, child.Parameters[4], child.Parameters[5])
			indent(builder, depth)
			builder.WriteString("C ")
			builder.WriteString(fmt.Sprintf("%v", cp1x))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cp1y))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cp2x))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cp2y))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationSmoothCurveTo:
			cp2x, cp2y := transform_point(effective, child.Parameters[0], child.Parameters[1])
			cx, cy = transform_point(effective, child.Parameters[2], child.Parameters[3])
			indent(builder, depth)
			builder.WriteString("S ")
			builder.WriteString(fmt.Sprintf("%v", cp2x))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cp2y))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationQuadraticCurveTo:
			cpx, cpy := transform_point(effective, child.Parameters[0], child.Parameters[1])
			cx, cy = transform_point(effective, child.Parameters[2], child.Parameters[3])
			indent(builder, depth)
			builder.WriteString("Q ")
			builder.WriteString(fmt.Sprintf("%v", cpx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cpy))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationSmoothQuadraticCurveTo:
			cx, cy = transform_point(effective, child.Parameters[0], child.Parameters[1])
			indent(builder, depth)
			builder.WriteString("T ")
			builder.WriteString(fmt.Sprintf("%v", cx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", cy))
			builder.WriteString("\n")

		case Path2DOperationArc:
			center_x := child.Parameters[0]
			center_y := child.Parameters[1]
			radius := child.Parameters[2]
			start_angle := child.Parameters[3]
			end_angle := child.Parameters[4]
			counterclockwise := child.Parameters[5] == 1

			end_x := center_x + radius*math.Cos(end_angle)
			end_y := center_y + radius*math.Sin(end_angle)

			t_end_x, t_end_y := transform_point(effective, end_x, end_y)
			cx = t_end_x
			cy = t_end_y

			t_center_x, t_center_y := transform_point(effective, center_x, center_y)
			t_start_x, t_start_y := transform_point(
				effective,
				center_x+radius*math.Cos(start_angle),
				center_y+radius*math.Sin(start_angle),
			)

			t_radius_x := math.Sqrt((t_start_x-t_center_x)*(t_start_x-t_center_x) + (t_start_y-t_center_y)*(t_start_y-t_center_y))
			t_radius_y := math.Sqrt((t_end_x-t_center_x)*(t_end_x-t_center_x) + (t_end_y-t_center_y)*(t_end_y-t_center_y))
			t_radius := (t_radius_x + t_radius_y) / 2.0

			if t_radius < 1e-9 {
				t_radius = radius
			}

			large_arc := 0
			angle_span := end_angle - start_angle

			if counterclockwise && angle_span > 0 {
				angle_span -= 2 * math.Pi
			} else if !counterclockwise && angle_span < 0 {
				angle_span += 2 * math.Pi
			}

			if math.Abs(angle_span) > math.Pi {
				large_arc = 1
			}

			sweep := 1

			if counterclockwise {
				sweep = 0
			}

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_radius))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_radius))
			builder.WriteString(" 0 ")
			builder.WriteString(strconv.Itoa(large_arc))
			builder.WriteString(" ")
			builder.WriteString(strconv.Itoa(sweep))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_end_x))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_end_y))
			builder.WriteString("\n")

		case Path2DOperationEllipse:
			center_x := child.Parameters[0]
			center_y := child.Parameters[1]
			rx := child.Parameters[2]
			ry := child.Parameters[3]
			rotation := child.Parameters[4]
			start_angle := child.Parameters[5]
			end_angle := child.Parameters[6]
			counterclockwise := child.Parameters[7] == 1

			cos_r := math.Cos(rotation)
			sin_r := math.Sin(rotation)

			end_x := center_x + rx*math.Cos(end_angle)*cos_r - ry*math.Sin(end_angle)*sin_r
			end_y := center_y + rx*math.Cos(end_angle)*sin_r + ry*math.Sin(end_angle)*cos_r

			t_end_x, t_end_y := transform_point(effective, end_x, end_y)
			cx = t_end_x
			cy = t_end_y

			t_center_x, t_center_y := transform_point(effective, center_x, center_y)

			start_x := center_x + rx*math.Cos(start_angle)*cos_r - ry*math.Sin(start_angle)*sin_r
			start_y := center_y + rx*math.Cos(start_angle)*sin_r + ry*math.Sin(start_angle)*cos_r

			t_start_x, t_start_y := transform_point(effective, start_x, start_y)

			t_rx := math.Sqrt((t_start_x-t_center_x)*(t_start_x-t_center_x) + (t_start_y-t_center_y)*(t_start_y-t_center_y))
			t_ry := math.Sqrt((t_end_x-t_center_x)*(t_end_x-t_center_x) + (t_end_y-t_center_y)*(t_end_y-t_center_y))

			if t_rx < 1e-9 {
				t_rx = rx
			}

			if t_ry < 1e-9 {
				t_ry = ry
			}

			large_arc := 0
			angle_span := end_angle - start_angle

			if counterclockwise && angle_span > 0 {
				angle_span -= 2 * math.Pi
			} else if !counterclockwise && angle_span < 0 {
				angle_span += 2 * math.Pi
			}

			if math.Abs(angle_span) > math.Pi {
				large_arc = 1
			}

			sweep := 1

			if counterclockwise {
				sweep = 0
			}

			rotation_deg := rotation * 180.0 / math.Pi

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_rx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_ry))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", rotation_deg))
			builder.WriteString(" ")
			builder.WriteString(strconv.Itoa(large_arc))
			builder.WriteString(" ")
			builder.WriteString(strconv.Itoa(sweep))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_end_x))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_end_y))
			builder.WriteString("\n")

		case Path2DOperationArcTo:
			x1 := child.Parameters[0]
			y1 := child.Parameters[1]
			x2 := child.Parameters[2]
			y2 := child.Parameters[3]
			radius := child.Parameters[4]

			t1x, t1y, t2x, t2y, arc_cx, arc_cy, sweep := compute_arc_to(cx, cy, x1, y1, x2, y2, radius)

			if !math.IsNaN(t1x) {

				t_t1x, t_t1y := transform_point(effective, t1x, t1y)
				t_t2x, t_t2y := transform_point(effective, t2x, t2y)
				t_arc_cx, t_arc_cy := transform_point(effective, arc_cx, arc_cy)

				if math.Abs(t_t1x-cx) > 1e-9 || math.Abs(t_t1y-cy) > 1e-9 {
					indent(builder, depth)
					builder.WriteString("L ")
					builder.WriteString(fmt.Sprintf("%v", t_t1x))
					builder.WriteString(" ")
					builder.WriteString(fmt.Sprintf("%v", t_t1y))
					builder.WriteString("\n")
				}

				large_arc := 0
				angle1 := math.Atan2(t_t1y-t_arc_cy, t_t1x-t_arc_cx)
				angle2 := math.Atan2(t_t2y-t_arc_cy, t_t2x-t_arc_cx)

				if sweep == 0 && angle2 > angle1 {
					angle2 -= 2 * math.Pi
				} else if sweep == 1 && angle2 < angle1 {
					angle2 += 2 * math.Pi
				}

				if math.Abs(angle2-angle1) > math.Pi {
					large_arc = 1
				}

				t_radius := math.Sqrt((t_t1x-t_arc_cx)*(t_t1x-t_arc_cx) + (t_t1y-t_arc_cy)*(t_t1y-t_arc_cy))

				if t_radius < 1e-9 {
					t_radius = radius
				}

				indent(builder, depth)
				builder.WriteString("A ")
				builder.WriteString(fmt.Sprintf("%v", t_radius))
				builder.WriteString(" ")
				builder.WriteString(fmt.Sprintf("%v", t_radius))
				builder.WriteString(" 0 ")
				builder.WriteString(strconv.Itoa(large_arc))
				builder.WriteString(" ")
				builder.WriteString(strconv.Itoa(sweep))
				builder.WriteString(" ")
				builder.WriteString(fmt.Sprintf("%v", t_t2x))
				builder.WriteString(" ")
				builder.WriteString(fmt.Sprintf("%v", t_t2y))
				builder.WriteString("\n")

				cx = t_t2x
				cy = t_t2y

			} else {

				tx, ty := transform_point(effective, x2, y2)

				indent(builder, depth)
				builder.WriteString("L ")
				builder.WriteString(fmt.Sprintf("%v", tx))
				builder.WriteString(" ")
				builder.WriteString(fmt.Sprintf("%v", ty))
				builder.WriteString("\n")

				cx = tx
				cy = ty

			}

		case Path2DOperationClosePath:
			indent(builder, depth)
			builder.WriteString("Z\n")

		case Path2DOperationRect:
			x := child.Parameters[0]
			y := child.Parameters[1]
			w := child.Parameters[2]
			h := child.Parameters[3]

			mx, my := transform_point(effective, x, y)
			lx1, ly1 := transform_point(effective, x+w, y)
			lx2, ly2 := transform_point(effective, x+w, y+h)
			lx3, ly3 := transform_point(effective, x, y+h)

			cx = mx
			cy = my

			indent(builder, depth)
			builder.WriteString("M ")
			builder.WriteString(fmt.Sprintf("%v", mx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", my))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx1))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly1))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx2))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly2))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx3))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly3))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("Z\n")

		case Path2DOperationRoundRect:
			x := child.Parameters[0]
			y := child.Parameters[1]
			w := child.Parameters[2]
			h := child.Parameters[3]
			r := child.Parameters[4]

			mx, my   := transform_point(effective, x+r, y)
			lx1, ly1 := transform_point(effective, x+w-r, y)
			ax1, ay1 := transform_point(effective, x+w, y+r)
			lx2, ly2 := transform_point(effective, x+w, y+h-r)
			ax2, ay2 := transform_point(effective, x+w-r, y+h)
			lx3, ly3 := transform_point(effective, x+r, y+h)
			ax3, ay3 := transform_point(effective, x, y+h-r)
			lx4, ly4 := transform_point(effective, x, y+r)
			ax4, ay4 := transform_point(effective, x+r, y)

			cx = mx
			cy = my

			t_rx, _ := transform_point(effective, x+r, y)
			_, t_ry := transform_point(effective, x, y+r)
			t_r := math.Sqrt((t_rx-mx)*(t_rx-mx) + (t_ry-my)*(t_ry-my))

			if t_r < 1e-9 {
				t_r = r
			}

			indent(builder, depth)
			builder.WriteString("M ")
			builder.WriteString(fmt.Sprintf("%v", mx))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", my))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx1))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly1))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" 0 0 1 ")
			builder.WriteString(fmt.Sprintf("%v", ax1))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ay1))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx2))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly2))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" 0 0 1 ")
			builder.WriteString(fmt.Sprintf("%v", ax2))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ay2))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx3))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly3))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" 0 0 1 ")
			builder.WriteString(fmt.Sprintf("%v", ax3))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ay3))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("L ")
			builder.WriteString(fmt.Sprintf("%v", lx4))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ly4))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("A ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", t_r))
			builder.WriteString(" 0 0 1 ")
			builder.WriteString(fmt.Sprintf("%v", ax4))
			builder.WriteString(" ")
			builder.WriteString(fmt.Sprintf("%v", ay4))
			builder.WriteString("\n")

			indent(builder, depth)
			builder.WriteString("Z\n")

		case Path2DOperationAddPath:
			cx, cy = child.write_svg(builder, effective, depth+1, cx, cy)

		}

	}

	return cx, cy

}

func indent(builder *strings.Builder, depth int) {

	for d := 0; d < depth; d++ {
		builder.WriteString("\t")
	}

}

func NewPath2DArc(x float64, y float64, radius float64, start_angle float64, end_angle float64, direction Direction) *Path2D {

	path := NewPath2D()

	angle_span := end_angle - start_angle

	if direction == DirectionCounterClockwise && angle_span > 0 {
		angle_span -= 2 * math.Pi
	} else if direction == DirectionClockwise && angle_span < 0 {
		angle_span += 2 * math.Pi
	}

	num_segments := int(math.Ceil(math.Abs(angle_span) / (math.Pi / 2)))
	angle_step := angle_span / float64(num_segments)

	path.MoveTo(
		x+radius*math.Cos(start_angle),
		y+radius*math.Sin(start_angle),
	)

	for s := 1; s <= num_segments; s++ {

		angle_mid := start_angle + float64(s-1)*angle_step + angle_step*0.5
		angle_end := start_angle + float64(s)*angle_step

		radius_mid := radius / math.Cos(angle_step*0.5)

		path.ArcTo(
			x+radius_mid*math.Cos(angle_mid),
			y+radius_mid*math.Sin(angle_mid),
			x+radius*math.Cos(angle_end),
			y+radius*math.Sin(angle_end),
			radius,
		)

	}

	return path

}

func NewPath2DEllipse(x float64, y float64, radius_x float64, radius_y float64, rotation float64, start_angle float64, end_angle float64, direction Direction) *Path2D {

	path := NewPath2D()

	angle_span := end_angle - start_angle

	if direction == DirectionCounterClockwise && angle_span > 0 {
		angle_span -= 2 * math.Pi
	} else if direction == DirectionClockwise && angle_span < 0 {
		angle_span += 2 * math.Pi
	}

	num_segments := int(math.Ceil(math.Abs(angle_span) / (math.Pi / 2)))
	angle_step := angle_span / float64(num_segments)

	cos_rot := math.Cos(rotation)
	sin_rot := math.Sin(rotation)

	start_on_ellipse := ellipse_point(radius_x, radius_y, cos_rot, sin_rot, start_angle)
	path.MoveTo(x+start_on_ellipse[0], y+start_on_ellipse[1])

	for s := 1; s <= num_segments; s++ {

		theta_start := start_angle + float64(s-1)*angle_step
		theta_end := start_angle + float64(s)*angle_step

		delta_theta := theta_end - theta_start
		kappa := 4.0 / 3.0 * math.Tan(delta_theta/4.0)

		start_pt := ellipse_point(radius_x, radius_y, cos_rot, sin_rot, theta_start)
		end_pt := ellipse_point(radius_x, radius_y, cos_rot, sin_rot, theta_end)

		start_tangent := ellipse_tangent(radius_x, radius_y, cos_rot, sin_rot, theta_start)
		end_tangent := ellipse_tangent(radius_x, radius_y, cos_rot, sin_rot, theta_end)

		cp1_x := start_pt[0] + kappa*start_tangent[0]
		cp1_y := start_pt[1] + kappa*start_tangent[1]
		cp2_x := end_pt[0] - kappa*end_tangent[0]
		cp2_y := end_pt[1] - kappa*end_tangent[1]

		path.BezierCurveTo(
			x+cp1_x, y+cp1_y,
			x+cp2_x, y+cp2_y,
			x+end_pt[0], y+end_pt[1],
		)

	}

	return path

}

func ellipse_point(rx float64, ry float64, cos_rot float64, sin_rot float64, theta float64) [2]float64 {

	ux := rx * math.Cos(theta)
	uy := ry * math.Sin(theta)

	return [2]float64{
		ux*cos_rot - uy*sin_rot,
		ux*sin_rot + uy*cos_rot,
	}

}

func ellipse_tangent(rx float64, ry float64, cos_rot float64, sin_rot float64, theta float64) [2]float64 {

	tx := -rx * math.Sin(theta)
	ty := ry * math.Cos(theta)

	return [2]float64{
		tx*cos_rot - ty*sin_rot,
		tx*sin_rot + ty*cos_rot,
	}

}

func NewPath2DRect(x float64, y float64, width float64, height float64) *Path2D {

	path := NewPath2D()

	path.MoveTo(x, y)
	path.LineTo(x+width, y)
	path.LineTo(x+width, y+height)
	path.LineTo(x, y+height)
	path.ClosePath()

	return path

}
