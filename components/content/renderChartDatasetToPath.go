package content

import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components/data"
import "strconv"
import "strings"

func renderChartDatasetToPath(dataset *data.Dataset, width int, height int, min_value int64, max_value int64, property string) *dom.Element {

	path    := dom.Document.CreateElementNS("http://www.w3.org/2000/svg", "path")
	delta_x := float64(width)
	delta_y := float64(height)

	if dataset.Length() > 1 {
		delta_x = float64(float64(width) / (float64(dataset.Length()) - float64(1)))
	}

	if max_value > 0 {
		delta_y = float64(float64(height) / float64(max_value))
	}

	if min_value < 0 {
		delta_y = float64(float64(height) / (float64(max_value) + float64(float64(-1) * float64(min_value))))
	}

	description := make([]string, 0)

	description = append(description, "M0," + strconv.Itoa(height + 1))

	for index := 0; index < dataset.Length(); index++ {

		data    := dataset.Get(index)
		val, ok := (*data)[property]

		if ok == true {

			switch val.(type) {

			case []byte:
				// Do Nothing

			case bool:

				value := val.(bool)

				if value == true {

					pos_x := int(delta_x * float64(index))
					pos_y := 0

					description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

				} else if value == false {

					pos_x := int(delta_x * float64(index))
					pos_y := height

					description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

				}

			case float32:

				value := val.(float32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case float64:

				value := val.(float64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int:

				value := val.(int)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int8:

				value := val.(int8)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int16:

				value := val.(int16)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int32:

				value := val.(int32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case int64:

				value := val.(int64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint:

				value := val.(uint)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint8:

				value := val.(uint8)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint16:

				value := val.(uint16)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint32:

				value := val.(uint32)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			case uint64:

				value := val.(uint64)
				pos_x := int(delta_x * float64(index))
				pos_y := height - int(delta_y * (float64(value) - float64(min_value)))

				description = append(description, "L" + strconv.Itoa(pos_x) + "," + strconv.Itoa(pos_y))

			}

		}

	}

	description = append(description, "L" + strconv.Itoa(width - 1) + "," + strconv.Itoa(height + 1))

	path.SetAttribute("d", strings.Join(description, " ") + " Z")

	return path

}
