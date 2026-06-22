//go:build wasm

package canvas2d

import "errors"
import "strconv"
import "strings"

func parseHSLColor(input string) (*Color, error) {

	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "hsl(") && strings.HasSuffix(input, ")") {

		inner := input[4 : len(input)-1]
		parts := strings.Split(inner, ",")

		if len(parts) == 3 {

			hue_str := strings.TrimSpace(parts[0])
			saturation_str := strings.TrimSpace(parts[1])
			lightness_str := strings.TrimSpace(parts[2])

			saturation_str = strings.TrimSuffix(saturation_str, "%")
			lightness_str = strings.TrimSuffix(lightness_str, "%")

			hue, err1 := strconv.ParseFloat(hue_str, 64)

			if err1 == nil {

				saturation, err2 := strconv.ParseFloat(saturation_str, 64)

				if err2 == nil {

					lightness, err3 := strconv.ParseFloat(lightness_str, 64)

					if err3 == nil {

						saturation = saturation / 100.0
						lightness = lightness / 100.0

						color := hsl_to_rgb(hue, saturation, lightness)
						color.Alpha = 1.0

						return color, nil

					}

				}

			}

		}

	}

	return nil, errors.New("parseHSLColor: invalid hsl color \"" + input + "\"")

}

