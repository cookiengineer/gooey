package content

import "strconv"

func renderTableValues(values map[string]any) (map[string]string, map[string]string) {

	result_values := make(map[string]string)
	result_types  := make(map[string]string)

	for key, val := range values {

		switch tmp := val.(type) {
		case []byte:

			result := ""

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

			result_types[key]  = "bytes"
			result_values[key] = result

		case bool:

			result_types[key]  = "bool"
			result_values[key] = strconv.FormatBool(tmp)

		case float32:

			result_types[key]  = "float32"
			result_values[key] = strconv.FormatFloat(float64(tmp), 'g', -1, 32)

		case float64:

			result_types[key]  = "float64"
			result_values[key] = strconv.FormatFloat(float64(tmp), 'g', -1, 64)

		case int:

			result_types[key]  = "int"
			result_values[key] = strconv.FormatInt(int64(tmp), 10)

		case int8:

			result_types[key]  = "int8"
			result_values[key] = strconv.FormatInt(int64(tmp), 10)

		case int16:

			result_types[key]  = "int16"
			result_values[key] = strconv.FormatInt(int64(tmp), 10)

		case int32:

			result_types[key]  = "int32"
			result_values[key] = strconv.FormatInt(int64(tmp), 10)

		case int64:

			result_types[key]  = "int64"
			result_values[key] = strconv.FormatInt(tmp, 10)

		case string:

			result_types[key]  = "string"
			result_values[key] = tmp

		case uint:

			result_types[key]  = "uint"
			result_values[key] = strconv.FormatUint(uint64(tmp), 10)

		case uint8:

			result_types[key]  = "uint8"
			result_values[key] = strconv.FormatUint(uint64(tmp), 10)

		case uint16:

			result_types[key]  = "uint16"
			result_values[key] = strconv.FormatUint(uint64(tmp), 10)

		case uint32:

			result_types[key]  = "uint32"
			result_values[key] = strconv.FormatUint(uint64(tmp), 10)

		case uint64:

			result_types[key]  = "uint64"
			result_values[key] = strconv.FormatUint(tmp, 10)

		}

	}

	return result_values, result_types

}
