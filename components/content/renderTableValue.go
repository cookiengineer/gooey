package content

import "encoding/json"
import "strconv"

func renderTableValue(value any) string {

	var result string

	switch tmp := value.(type) {
	case []byte:

		for t := 0; t < len(tmp); t++ {

			hex := strconv.FormatUint(uint64(tmp[t]), 16)

			if len(hex) == 1 {
				result += "0x0" + hex
			} else {
				result += "0x" + hex
			}

			if t < len(tmp) - 1 {
				result += " "
			}

		}

	case bool:
		result = strconv.FormatBool(tmp)
	case float32:
		result = strconv.FormatFloat(float64(tmp), 'g', -1, 32)
	case float64:
		result = strconv.FormatFloat(float64(tmp), 'g', -1, 64)
	case int:
		result = strconv.FormatInt(int64(tmp), 10)
	case int8:
		result = strconv.FormatInt(int64(tmp), 10)
	case int16:
		result = strconv.FormatInt(int64(tmp), 10)
	case int32:
		result = strconv.FormatInt(int64(tmp), 10)
	case int64:
		result = strconv.FormatInt(tmp, 10)
	case string:
		result = tmp
	case uint:
		result = strconv.FormatUint(uint64(tmp), 10)
	case uint8:
		result = strconv.FormatUint(uint64(tmp), 10)
	case uint16:
		result = strconv.FormatUint(uint64(tmp), 10)
	case uint32:
		result = strconv.FormatUint(uint64(tmp), 10)
	case uint64:
		result = strconv.FormatUint(tmp, 10)
	case map[string]any:

		bytes, err := json.Marshal(tmp)

		if err == nil {
			result = string(bytes)
		}

	default:

		bytes, err := json.Marshal(tmp)

		if err == nil {
			result = string(bytes)
		}

	}

	return result

}
