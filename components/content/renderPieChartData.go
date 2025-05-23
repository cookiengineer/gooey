package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/data"
import "math"
import "strconv"
import "strings"
import "fmt"

func getArcCoordinates(percentage float64) (float64, float64) {

	// TODO: This is kind of wrong, it should start at the top and not at the bottom?
	// TODO: largeArcFlag if percentage > 50% then 1 else 0?
	x := math.Sin(2 * math.Pi * percentage)
	y := math.Cos(2 * math.Pi * percentage)

	return x, y

}

func renderPieChartData(data *data.Data, width int, height int, min_value int64, max_value int64, property string, offset float64) (*dom.Element, *dom.Element, float64) {

	path       := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "path")
	text       := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "text")
	percentage := 0.0

	delta := float64(max_value)

	if min_value > 0 {
		delta = float64(max_value) - float64(min_value)
	} else if min_value < 0 {
		delta = float64(max_value) + (float64(-1) * float64(min_value))
	}

	radius      := math.Min(float64(width / 2), float64(height / 2))
	center_x    := float64(width / 2)
	center_y    := float64(height / 2)
	description := make([]string, 0)

	val, ok := (*data)[property]

	if ok == true {

		switch val.(type) {
		case []byte:
			// Do Nothing
		case bool:

			value := val.(bool)

			if value == true {

				// TODO

			} else if value == false {

				// TODO

			}

		case float32:

			value            := val.(float32)
			percentage        = float64(value) / delta
			start_x, start_y := getArcCoordinates(offset)
			end_x,   end_y   := getArcCoordinates(offset + percentage)

			fmt.Println(property, offset, percentage)
			fmt.Println(value, delta)

			description = append(description, "M " + strconv.FormatFloat(center_x + start_x * radius, 'f', -1, 64) + " " + strconv.FormatFloat(center_y + start_y * radius, 'f', -1, 64))
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 0 " + strconv.FormatFloat(center_x + end_x * radius, 'f', -1, 64) + " " + strconv.FormatFloat(center_y + end_y * radius, 'f', -1, 64))
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			// TODO: A rx ry x-axis-rotation large-arc-flag sweep-flag x y

		case float64:
		case int:
		case int8:
		case int16:
		case int32:
		case int64:
		case string:
			// TODO: Parse percentage string
		case uint:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		}

	}

	path.SetAttribute("d", strings.Join(description, " ") + " Z")

	return path, text, percentage

}
