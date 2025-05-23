package content

import "github.com/cookiengineer/gooey/components/data"
import "math"
import "strconv"
import "strings"

func sumChartData(data *data.Data, properties []string) (int64) {

	var result int64 = 0

	for _, property := range properties {

		val, ok := (*data)[property]

		if ok == true {

			switch val.(type) {
			case []byte:
				// Do Nothing
			case bool:

				value := val.(bool)

				if value == true || value == false {
					result += 1
				}

			case float32:

				value1 := math.Floor(float64(val.(float32)))
				value2 := math.Round(float64(val.(float32)))

				if value1 > 0.0 && value2 < 1.0 {
					result = 1
					break
				} else {
					result += int64(val.(float32))
				}

			case float64:

				value1 := math.Floor(val.(float64))
				value2 := math.Round(val.(float64))

				if value1 > 0.0 && value2 < 1.0 {
					result = 1
					break
				} else {
					result += int64(val.(float64))
				}

			case int:
				result += int64(val.(int))
			case int8:
				result += int64(val.(int8))
			case int16:
				result += int64(val.(int16))
			case int32:
				result += int64(val.(int32))
			case int64:
				result += int64(val.(int64))
			case string:

				value := val.(string)

				if strings.HasSuffix(value, "%") {

					value = value[0:len(value)-1]

					if strings.Contains(value, ".") {

						tmp, err := strconv.ParseFloat(value, 32)

						if err == nil && tmp >= 0.0 && tmp <= 100.0 {
							result += int64(tmp)
						}

					} else {

						tmp, err := strconv.ParseInt(value, 10, 32)

						if err == nil && tmp >= 0 && tmp <= 100 {
							result += int64(tmp)
						}

					}

				}

			case uint:
				result += int64(val.(uint))
			case uint8:
				result += int64(val.(uint8))
			case uint16:
				result += int64(val.(uint16))
			case uint32:
				result += int64(val.(uint32))
			case uint64:
				result += int64(val.(uint64))
			}

		}

	}

	if result >= 0 && result <= 1 {
		// float percentage 1.0
		result = 1
	} else if result >= 1 && result <= 100 {
		// string percentage 100%
		result = 100
	} else {
		// Do Nothing
	}

	return result

}
