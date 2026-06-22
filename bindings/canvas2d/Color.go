//go:build wasm

package canvas2d

import "errors"
import "math"
import "strconv"

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

func NewColor(input string) (*Color, error) {

	color1, err1 := parseHexColor(input)

	if err1 == nil {
		return color1, nil
	}

	color2, err2 := parseRGBColor(input)

	if err2 == nil {
		return color2, nil
	}

	color3, err3 := parseRGBAColor(input)

	if err3 == nil {
		return color3, nil
	}

	color4, err4 := parseHSLColor(input)

	if err4 == nil {
		return color4, nil
	}

	color5, err5 := parseHSLAColor(input)

	if err5 == nil {
		return color5, nil
	}

	return nil, errors.New("NewColor: Unknown color syntax \"" + input + "\"")

}

func (color *Color) String() string {

	result := ""
	red    := int(math.Round(color.Red))
	green  := int(math.Round(color.Green))
	blue   := int(math.Round(color.Blue))
	alpha  := strconv.FormatFloat(color.Alpha, 'f', -1, 64)

	if color.Alpha == 1.0 {
		result = "rgb(" + strconv.Itoa(red) + "," + strconv.Itoa(green) + "," + strconv.Itoa(blue) + ")"
	} else {
		result = "rgba(" + strconv.Itoa(red) + "," + strconv.Itoa(green) + "," + strconv.Itoa(blue) + "," + alpha + ")"
	}

	return result

}
