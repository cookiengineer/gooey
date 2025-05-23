package data

import "strconv"
import "strings"

type Data map[string]any

func ParseData(values map[string]string, types map[string]string) Data {

	var result Data = Data(map[string]any{})

	if len(values) == len(types) {

		for key, val := range values {

			typ, ok := types[key]

			if ok == true {

				switch typ {
				case "bytes":

					bytes  := make([]byte, 0)
					chunks := strings.Split(strings.TrimSpace(val), " ")

					for _, chunk := range chunks {

						if strings.HasPrefix(chunk, "0x") {

							tmp, err := strconv.ParseUint(string(chunk[2:4]), 16, 8)

							if err == nil {
								bytes = append(bytes, byte(tmp))
							}

						}

					}

					result[key] = bytes

				case "bool":

					if val == "true" {
						result[key] = true
					} else {
						result[key] = false
					}

				case "float32":

					tmp, err := strconv.ParseFloat(val, 32)

					if err == nil {
						result[key] = float32(tmp)
					} else {
						result[key] = float32(0.0)
					}

				case "float64":

					tmp, err := strconv.ParseFloat(val, 64)

					if err == nil {
						result[key] = float64(tmp)
					} else {
						result[key] = float64(0.0)
					}

				case "int":

					tmp, err := strconv.ParseInt(val, 10, 0)

					if err == nil {
						result[key] = int(tmp)
					} else {
						result[key] = int(0)
					}

				case "int8":

					tmp, err := strconv.ParseInt(val, 10, 8)

					if err == nil {
						result[key] = int8(tmp)
					} else {
						result[key] = int8(0)
					}

				case "int16":

					tmp, err := strconv.ParseInt(val, 10, 16)

					if err == nil {
						result[key] = int16(tmp)
					} else {
						result[key] = int16(0)
					}

				case "int32":

					tmp, err := strconv.ParseInt(val, 10, 32)

					if err == nil {
						result[key] = int32(tmp)
					} else {
						result[key] = int32(0)
					}

				case "int64":

					tmp, err := strconv.ParseInt(val, 10, 64)

					if err == nil {
						result[key] = int64(tmp)
					} else {
						result[key] = int64(0)
					}

				case "string":

					result[key] = val

				case "uint":

					tmp, err := strconv.ParseUint(val, 10, 0)

					if err == nil {
						result[key] = uint(tmp)
					} else {
						result[key] = uint(0)
					}

				case "uint8":

					tmp, err := strconv.ParseUint(val, 10, 8)

					if err == nil {
						result[key] = uint8(tmp)
					} else {
						result[key] = uint8(0)
					}

				case "uint16":

					tmp, err := strconv.ParseUint(val, 10, 16)

					if err == nil {
						result[key] = uint16(tmp)
					} else {
						result[key] = uint16(0)
					}

				case "uint32":

					tmp, err := strconv.ParseUint(val, 10, 32)

					if err == nil {
						result[key] = uint32(tmp)
					} else {
						result[key] = uint32(0)
					}

				case "uint64":

					tmp, err := strconv.ParseUint(val, 10, 64)

					if err == nil {
						result[key] = uint64(tmp)
					} else {
						result[key] = uint64(0)
					}

				}

			}

		}

	}

	return result

}

func (data *Data) String() (map[string]string, map[string]string) {

	result_values := make(map[string]string)
	result_types  := make(map[string]string)

	for key, val := range *data {

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

func (data *Data) StringProperty(property string) (string, string) {

	result_value := ""
	result_type  := ""

	if property != "" {

		val, ok := (*data)[property]

		if ok == true {

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

				result_type  = "bytes"
				result_value = result

			case bool:

				result_type  = "bool"
				result_value = strconv.FormatBool(tmp)

			case float32:

				result_type  = "float32"
				result_value = strconv.FormatFloat(float64(tmp), 'g', -1, 32)

			case float64:

				result_type  = "float64"
				result_value = strconv.FormatFloat(float64(tmp), 'g', -1, 64)

			case int:

				result_type  = "int"
				result_value = strconv.FormatInt(int64(tmp), 10)

			case int8:

				result_type  = "int8"
				result_value = strconv.FormatInt(int64(tmp), 10)

			case int16:

				result_type  = "int16"
				result_value = strconv.FormatInt(int64(tmp), 10)

			case int32:

				result_type  = "int32"
				result_value = strconv.FormatInt(int64(tmp), 10)

			case int64:

				result_type  = "int64"
				result_value = strconv.FormatInt(tmp, 10)

			case string:

				result_type  = "string"
				result_value = tmp

			case uint:

				result_type  = "uint"
				result_value = strconv.FormatUint(uint64(tmp), 10)

			case uint8:

				result_type  = "uint8"
				result_value = strconv.FormatUint(uint64(tmp), 10)

			case uint16:

				result_type  = "uint16"
				result_value = strconv.FormatUint(uint64(tmp), 10)

			case uint32:

				result_type  = "uint32"
				result_value = strconv.FormatUint(uint64(tmp), 10)

			case uint64:

				result_type  = "uint64"
				result_value = strconv.FormatUint(tmp, 10)

			}

		}

	}

	return result_value, result_type

}

