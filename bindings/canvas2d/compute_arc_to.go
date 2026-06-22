//go:build wasm

package canvas2d

import "math"

func compute_arc_to(x0 float64, y0 float64, x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (t1x float64, t1y float64, t2x float64, t2y float64, cx float64, cy float64, sweep int) {

	if radius <= 0 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0
	}

	dx0 := x0 - x1
	dy0 := y0 - y1
	len0 := math.Sqrt(dx0*dx0 + dy0*dy0)

	if len0 < 1e-12 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0
	}

	dx1 := x2 - x1
	dy1 := y2 - y1
	len1 := math.Sqrt(dx1*dx1 + dy1*dy1)

	if len1 < 1e-12 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0
	}

	ux0 := dx0 / len0
	uy0 := dy0 / len0
	ux1 := dx1 / len1
	uy1 := dy1 / len1

	dot := ux0*ux1 + uy0*uy1

	if dot > 1.0 {
		dot = 1.0
	} else if dot < -1.0 {
		dot = -1.0
	}

	half_angle := math.Acos(dot) / 2.0
	sin_half := math.Sin(half_angle)

	if sin_half < 1e-12 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0
	}

	distance := radius / math.Tan(half_angle)

	if distance > len0 || distance > len1 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 0
	}

	t1x = x1 + ux0*distance
	t1y = y1 + uy0*distance
	t2x = x1 + ux1*distance
	t2y = y1 + uy1*distance

	cross := ux0*uy1 - uy0*ux1

	if cross >= 0 {
		sweep = 1
	} else {
		sweep = 0
	}

	mid_x := (t1x + t2x) / 2.0
	mid_y := (t1y + t2y) / 2.0

	perp_x := -(t2y - t1y)
	perp_y := t2x - t1x
	perp_len := math.Sqrt(perp_x*perp_x + perp_y*perp_y)

	offset := radius / math.Sin(half_angle)
	center_dist := math.Sqrt(offset*offset - radius*radius)

	cx = mid_x + (perp_x/perp_len)*center_dist
	cy = mid_y + (perp_y/perp_len)*center_dist

	check_angle1 := math.Atan2(t1y-cy, t1x-cx)
	check_angle2 := math.Atan2(t2y-cy, t2x-cx)
	mid_angle := math.Atan2((y1 - cy), (x1 - cx))

	diff1 := compute_angle_difference(check_angle1, mid_angle)
	diff2 := compute_angle_difference(check_angle2, mid_angle)

	if diff1 < 0 && diff2 < 0 {
		sweep = 1 - sweep
	} else if diff1 > 0 && diff2 > 0 {

	} else {
		sweep = 1 - sweep
	}

	return t1x, t1y, t2x, t2y, cx, cy, sweep

}

