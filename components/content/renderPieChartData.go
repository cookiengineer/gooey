package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/data"
import "math"
import "strconv"
import "strings"

func getArcCoordinates(percentage float64) (float64, float64) {

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
			// Do Nothing
		case float32:

			value     := val.(float32)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatFloat(float64(value), 'g', -1, 32))

		case float64:

			value     := val.(float64)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatFloat(float64(value), 'g', -1, 64))

		case int:

			value     := val.(int)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatInt(int64(value), 10))

		case int8:

			value     := val.(int8)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatInt(int64(value), 10))

		case int16:

			value     := val.(int16)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatInt(int64(value), 10))

		case int32:

			value     := val.(int32)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatInt(int64(value), 10))

		case int64:

			value     := val.(int64)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatInt(int64(value), 10))

		case string:

			value := val.(string)

			if strings.Contains(value, ".") && strings.HasSuffix(value, "%") {

				tmp, err := strconv.ParseFloat(value[0:len(value)-1], 32)

				if err == nil && tmp >= 0.0 && tmp <= 100.0 {

					percentage = float64(tmp) / float64(100.0)

					arc_start_x, arc_start_y := getArcCoordinates(offset)
					arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
					arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

					text_x := int(center_x + arc_mid_x * radius / 2)
					text_y := int(center_y - arc_mid_y * radius / 2)

					move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
					move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
					end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
					end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

					description = append(description, "M " + move_x + " " + move_y)
					description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
					description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

					text = renderTextAt(text_x, text_y, value)

				}

			} else if strings.HasSuffix(value, "%") {

				tmp, err := strconv.ParseInt(value[0:len(value)-1], 10, 32)

				if err == nil && tmp >= 0 && tmp <= 100 {

					percentage = float64(tmp) / float64(100.0)

					arc_start_x, arc_start_y := getArcCoordinates(offset)
					arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
					arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

					text_x := int(center_x + arc_mid_x * radius / 2)
					text_y := int(center_y - arc_mid_y * radius / 2)

					move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
					move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
					end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
					end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

					description = append(description, "M " + move_x + " " + move_y)
					description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
					description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

					text = renderTextAt(text_x, text_y, value)

				}

			}

		case uint:

			value     := val.(uint)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatUint(uint64(value), 10))

		case uint8:

			value     := val.(uint8)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatUint(uint64(value), 10))

		case uint16:

			value     := val.(uint16)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatUint(uint64(value), 10))

		case uint32:

			value     := val.(uint32)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatUint(uint64(value), 10))

		case uint64:

			value     := val.(uint64)
			percentage = float64(value) / delta

			arc_start_x, arc_start_y := getArcCoordinates(offset)
			arc_mid_x,   arc_mid_y   := getArcCoordinates(offset + percentage / 2)
			arc_end_x,   arc_end_y   := getArcCoordinates(offset + percentage)

			text_x := int(center_x + arc_mid_x * radius / 2)
			text_y := int(center_y - arc_mid_y * radius / 2)

			move_x := strconv.FormatFloat(center_x + arc_start_x * radius, 'f', -1, 64)
			move_y := strconv.FormatFloat(center_y - arc_start_y * radius, 'f', -1, 64)
			end_x  := strconv.FormatFloat(center_x + arc_end_x   * radius, 'f', -1, 64)
			end_y  := strconv.FormatFloat(center_y - arc_end_y   * radius, 'f', -1, 64)

			description = append(description, "M " + move_x + " " + move_y)
			description = append(description, "A " + strconv.FormatFloat(radius, 'f', -1, 64) + " " + strconv.FormatFloat(radius, 'f', -1, 64) + " 0 0 1 " + end_x + " " + end_y)
			description = append(description, "L " + strconv.FormatFloat(center_x, 'f', -1, 64) + " " + strconv.FormatFloat(center_y, 'f', -1, 64))

			text = renderTextAt(text_x, text_y, strconv.FormatUint(uint64(value), 10))

		}

	}

	path.SetAttribute("d", strings.Join(description, " ") + " Z")

	return path, text, percentage

}
