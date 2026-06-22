//go:build wasm

package canvas2d

import "errors"
import "strconv"

func parseHexColor(input string) (*Color, error) {

	input_length := len(input)

	if input_length == 0 {

		return nil, errors.New("parseHexColor: empty input")

	} else if input[0] == '#' {

		hex_part := input[1:]
		hex_length := len(hex_part)

		var red_str string
		var green_str string
		var blue_str string
		var alpha_str string

		if hex_length == 3 {

			red_str = string([]byte{hex_part[0], hex_part[0]})
			green_str = string([]byte{hex_part[1], hex_part[1]})
			blue_str = string([]byte{hex_part[2], hex_part[2]})
			alpha_str = "ff"

		} else if hex_length == 4 {

			red_str = string([]byte{hex_part[0], hex_part[0]})
			green_str = string([]byte{hex_part[1], hex_part[1]})
			blue_str = string([]byte{hex_part[2], hex_part[2]})
			alpha_str = string([]byte{hex_part[3], hex_part[3]})

		} else if hex_length == 6 {

			red_str = hex_part[0:2]
			green_str = hex_part[2:4]
			blue_str = hex_part[4:6]
			alpha_str = "ff"

		} else if hex_length == 8 {

			red_str = hex_part[0:2]
			green_str = hex_part[2:4]
			blue_str = hex_part[4:6]
			alpha_str = hex_part[6:8]

		} else {

			return nil, errors.New("parseHex: invalid hex color length \"" + input + "\"")

		}

		red, err1 := strconv.ParseUint(red_str, 16, 8)

		if err1 == nil {

			green, err2 := strconv.ParseUint(green_str, 16, 8)

			if err2 == nil {

				blue, err3 := strconv.ParseUint(blue_str, 16, 8)

				if err3 == nil {

					alpha, err4 := strconv.ParseUint(alpha_str, 16, 8)

					if err4 == nil {

						var color Color

						color.Red = float64(red)
						color.Green = float64(green)
						color.Blue = float64(blue)
						color.Alpha = float64(alpha) / 255.0

						return &color, nil

					}

				}

			}

		}

	}

	return nil, errors.New("parseHexColor: not a hex color \"" + input + "\"")

}
