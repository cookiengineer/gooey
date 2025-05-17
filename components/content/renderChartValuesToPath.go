package content

import "strconv"
import "strings"

func renderChartValuesToPath(width int, height int, min_value int64, max_value int64, dataset []ChartData, property string) string {

	delta_x := int(width  / len(dataset))
	delta_y := int(int64(height) / max_value)

	if min_value < 0 {
		delta_y = int(int64(height) / (int64(max_value) + int64(-1 * min_value)))
	}

	result := make([]string, 0)
	result = append(result, "M0," + strconv.Itoa(height + 1))

	for d := 0; d < len(dataset); d++ {

		val, ok := dataset[d][property]

		if ok == true {

			switch val.(type) {

			case []byte:
				// Do Nothing

			case bool:

				value := val.(bool)

				if value == true {

					pos_x := int(delta_x * (d + 1))
					pos_y := 0

					result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

				} else if value == false {

					pos_x := int(delta_x * (d + 1))
					pos_y := height

					result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

				}

			case float32:

				value := val.(float32)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(float32(delta_y) * (value - float32(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case float64:

				value := val.(float64)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(float64(delta_y) * (value - float64(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int:

				value := val.(int)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(int(delta_y) * (value - int(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int8:

				value := val.(int8)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(int8(delta_y) * (value - int8(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int16:

				value := val.(int16)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(int16(delta_y) * (value - int16(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int32:

				value := val.(int32)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(int32(delta_y) * (value - int32(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int64:

				value := val.(int64)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(int64(delta_y) * (value - int64(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint:

				value := val.(uint)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(uint(delta_y) * (value - uint(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint8:

				value := val.(uint8)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(uint8(delta_y) * (value - uint8(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint16:

				value := val.(uint16)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(uint16(delta_y) * (value - uint16(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint32:

				value := val.(uint32)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(uint32(delta_y) * (value - uint32(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint64:

				value := val.(uint64)
				pos_x := int(delta_x * (d + 1))
				pos_y := height - int(uint64(delta_y) * (value - uint64(min_value)))

				result = append(result, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			}

		}

	}

	result = append(result, "L" + strconv.Itoa(width - 1) + "," + strconv.Itoa(height + 1))

	return strings.Join(result, " ") + " Z"

}
