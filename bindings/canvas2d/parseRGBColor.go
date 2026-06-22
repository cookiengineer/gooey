//go:build wasm

package canvas2d

import "errors"
import "strconv"
import "strings"

func parseRGBColor(input string) (*Color, error) {

	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "rgb(") && strings.HasSuffix(input, ")") {

		inner := input[4 : len(input)-1]
		parts := strings.Split(inner, ",")

		if len(parts) == 3 {

			red_str := strings.TrimSpace(parts[0])
			green_str := strings.TrimSpace(parts[1])
			blue_str := strings.TrimSpace(parts[2])

			red, err1 := strconv.Atoi(red_str)

			if err1 == nil {

				green, err2 := strconv.Atoi(green_str)

				if err2 == nil {

					blue, err3 := strconv.Atoi(blue_str)

					if err3 == nil {

						if red >= 0 && red <= 255 && green >= 0 && green <= 255 && blue >= 0 && blue <= 255 {

							var color Color

							color.Red = float64(red)
							color.Green = float64(green)
							color.Blue = float64(blue)
							color.Alpha = 1.0

							return &color, nil

						}

					}

				}

			}

		}

	}

	return nil, errors.New("parseRGBColor: invalid rgb color \"" + input + "\"")

}

