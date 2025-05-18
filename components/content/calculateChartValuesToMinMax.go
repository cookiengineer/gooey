package content

import "github.com/cookiengineer/gooey/components/data"
import "math"

func calculateChartDatasetMinMax(dataset *data.Dataset, properties []string) (int64, int64) {

	var min_value int64 = 0
	var max_value int64 = 0

	for _, property := range properties {

		for index := 0; index < dataset.Length(); index++ {

			data    := dataset.Get(index)
			val, ok := (*data)[property]

			if ok == true {

				switch val.(type) {
				case []byte:
					// Do Nothing
				case bool:

					value := val.(bool)

					if value == true || value == false {
						max_value = 1
						min_value = 0
					}

				case float32:

					value1 := math.Floor(float64(val.(float32)))
					value2 := math.Round(float64(val.(float32)))

					if int64(value1) < min_value {
						min_value = int64(value1)
					}

					if int64(value2) > max_value {
						max_value = int64(value2)
					}

				case float64:

					value1 := math.Floor(val.(float64))
					value2 := math.Round(val.(float64))

					if int64(value1) < min_value {
						min_value = int64(value1)
					}

					if int64(value2) > max_value {
						max_value = int64(value2)
					}

				case int:

					value := int64(val.(int))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case int8:

					value := int64(val.(int8))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case int16:

					value := int64(val.(int16))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case int32:

					value := int64(val.(int32))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case int64:

					value := val.(int64)

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case string:

					// Do Nothing

				case uint:

					value := int64(val.(uint))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case uint8:

					value := int64(val.(uint8))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case uint16:

					value := int64(val.(uint16))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case uint32:

					value := int64(val.(uint32))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				case uint64:

					value := int64(val.(uint64))

					if value < min_value {
						min_value = value
					}

					if value > max_value {
						max_value = value
					}

				}

			}

		}

	}

	return min_value, max_value

}
