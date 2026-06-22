//go:build wasm

package canvas2d

func hue_to_rgb(p float64, q float64, t float64) float64 {

	if t < 0.0 {
		t = t + 1.0
	} else if t > 1.0 {
		t = t - 1.0
	}

	if t < 1.0/6.0 {
		return p + (q-p)*6.0*t
	} else if t < 1.0/2.0 {
		return q
	} else if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6.0
	}

	return p

}
