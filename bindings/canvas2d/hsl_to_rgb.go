//go:build wasm

package canvas2d

import "math"

func hsl_to_rgb(hue_degrees float64, saturation float64, lightness float64) *Color {

	var color Color

	hue := hue_degrees / 360.0

	if saturation == 0.0 {

		value := lightness * 255.0

		color.Red = math.Round(value)
		color.Green = math.Round(value)
		color.Blue = math.Round(value)

	} else {

		var q float64

		if lightness < 0.5 {
			q = lightness * (1.0 + saturation)
		} else {
			q = lightness + saturation - lightness*saturation
		}

		p := 2.0*lightness - q

		color.Red = math.Round(hue_to_rgb(p, q, hue+1.0/3.0) * 255.0)
		color.Green = math.Round(hue_to_rgb(p, q, hue) * 255.0)
		color.Blue = math.Round(hue_to_rgb(p, q, hue-1.0/3.0) * 255.0)

	}

	return &color

}

