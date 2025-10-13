package data

import "strconv"
import "strings"

type Data map[string]any

func ParseData(values map[string]string, types map[string]string) Data {

	data := Data(map[string]any{})

	if len(values) == len(types) {

		for property, value := range values {

			typ, ok := types[property]

			if ok == true {

				switch typ {

				case "[]bool":

					bools  := make([]bool, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						if chunk == "true" {
							bools = append(bools, true)
						} else {
							bools = append(bools, false)
						}

					}

					data[property] = bools

				case "bool":

					if value == "true" {
						data[property] = true
					} else {
						data[property] = false
					}

				case "[]byte":

					bytes  := make([]byte, 0)
					chunks := strings.Split(strings.TrimSpace(value), " ")

					for _, chunk := range chunks {

						if strings.HasPrefix(chunk, "0x") {

							tmp, err := strconv.ParseUint(string(chunk[2:4]), 16, 8)

							if err == nil {
								bytes = append(bytes, byte(tmp))
							}

						}

					}

					data[property] = bytes

				case "[]float32":

					floats := make([]float32, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseFloat(chunk, 32)

						if err == nil {
							floats = append(floats, float32(tmp))
						} else {
							floats = append(floats, float32(0.0))
						}

					}

					data[property] = floats

				case "float32":

					tmp, err := strconv.ParseFloat(value, 32)

					if err == nil {
						data[property] = float32(tmp)
					} else {
						data[property] = float32(0.0)
					}

				case "[]float64":

					floats := make([]float64, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseFloat(chunk, 64)

						if err == nil {
							floats = append(floats, float64(tmp))
						} else {
							floats = append(floats, float64(0.0))
						}

					}

					data[property] = floats

				case "float64":

					tmp, err := strconv.ParseFloat(value, 64)

					if err == nil {
						data[property] = float64(tmp)
					} else {
						data[property] = float64(0.0)
					}

				case "[]int":

					ints   := make([]int, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseInt(chunk, 10, 0)

						if err == nil {
							ints = append(ints, int(tmp))
						} else {
							ints = append(ints, int(0))
						}

					}

					data[property] = ints

				case "int":

					tmp, err := strconv.ParseInt(value, 10, 0)

					if err == nil {
						data[property] = int(tmp)
					} else {
						data[property] = int(0)
					}

				case "[]int8":

					ints   := make([]int8, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseInt(chunk, 10, 8)

						if err == nil {
							ints = append(ints, int8(tmp))
						} else {
							ints = append(ints, int8(0))
						}

					}

					data[property] = ints

				case "int8":

					tmp, err := strconv.ParseInt(value, 10, 8)

					if err == nil {
						data[property] = int8(tmp)
					} else {
						data[property] = int8(0)
					}

				case "[]int16":

					ints   := make([]int16, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseInt(chunk, 10, 16)

						if err == nil {
							ints = append(ints, int16(tmp))
						} else {
							ints = append(ints, int16(0))
						}

					}

					data[property] = ints

				case "int16":

					tmp, err := strconv.ParseInt(value, 10, 16)

					if err == nil {
						data[property] = int16(tmp)
					} else {
						data[property] = int16(0)
					}

				case "[]int32":

					ints   := make([]int32, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseInt(chunk, 10, 32)

						if err == nil {
							ints = append(ints, int32(tmp))
						} else {
							ints = append(ints, int32(0))
						}

					}

					data[property] = ints

				case "int32":

					tmp, err := strconv.ParseInt(value, 10, 32)

					if err == nil {
						data[property] = int32(tmp)
					} else {
						data[property] = int32(0)
					}

				case "[]int64":

					ints   := make([]int64, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseInt(chunk, 10, 64)

						if err == nil {
							ints = append(ints, int64(tmp))
						} else {
							ints = append(ints, int64(0))
						}

					}

					data[property] = ints

				case "int64":

					tmp, err := strconv.ParseInt(value, 10, 64)

					if err == nil {
						data[property] = int64(tmp)
					} else {
						data[property] = int64(0)
					}

				case "[]string":

					strs   := make([]string, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {
						strs = append(strs, chunk)
					}

					data[property] = strs

				case "string":

					data[property] = string(value)

				case "[]uint":

					uints  := make([]uint, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseUint(chunk, 10, 0)

						if err == nil {
							uints = append(uints, uint(tmp))
						} else {
							uints = append(uints, uint(0))
						}

					}

					data[property] = uints

				case "uint":

					tmp, err := strconv.ParseUint(value, 10, 0)

					if err == nil {
						data[property] = uint(tmp)
					} else {
						data[property] = uint(0)
					}

				case "[]uint8":

					uints  := make([]uint8, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseUint(chunk, 10, 8)

						if err == nil {
							uints = append(uints, uint8(tmp))
						} else {
							uints = append(uints, uint8(0))
						}

					}

					data[property] = uints

				case "uint8":

					tmp, err := strconv.ParseUint(value, 10, 8)

					if err == nil {
						data[property] = uint8(tmp)
					} else {
						data[property] = uint8(0)
					}

				case "[]uint16":

					uints  := make([]uint16, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseUint(chunk, 10, 16)

						if err == nil {
							uints = append(uints, uint16(tmp))
						} else {
							uints = append(uints, uint16(0))
						}

					}

					data[property] = uints

				case "uint16":

					tmp, err := strconv.ParseUint(value, 10, 16)

					if err == nil {
						data[property] = uint16(tmp)
					} else {
						data[property] = uint16(0)
					}

				case "[]uint32":

					uints  := make([]uint32, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseUint(chunk, 10, 32)

						if err == nil {
							uints = append(uints, uint32(tmp))
						} else {
							uints = append(uints, uint32(0))
						}

					}

					data[property] = uints

				case "uint32":

					tmp, err := strconv.ParseUint(value, 10, 32)

					if err == nil {
						data[property] = uint32(tmp)
					} else {
						data[property] = uint32(0)
					}

				case "[]uint64":

					uints  := make([]uint64, 0)
					chunks := strings.Split(strings.TrimSpace(value), ",")

					for _, chunk := range chunks {

						tmp, err := strconv.ParseUint(chunk, 10, 64)

						if err == nil {
							uints = append(uints, uint64(tmp))
						} else {
							uints = append(uints, uint64(0))
						}

					}

					data[property] = uints

				case "uint64":

					tmp, err := strconv.ParseUint(value, 10, 64)

					if err == nil {
						data[property] = uint64(tmp)
					} else {
						data[property] = uint64(0)
					}

				}

			}

		}

	}

	return data

}

func (data *Data) String() (map[string]string, map[string]string) {

	result_values := make(map[string]string)
	result_types  := make(map[string]string)

	for property, _ := range *data {

		typ, val := data.StringProperty(property)

		if typ != "" && val != "" {
			result_types[property]  = typ
			result_values[property] = val
		}

	}

	return result_types, result_values

}

func (data *Data) StringProperty(property string) (string, string) {

	result_type  := ""
	result_value := ""

	if property != "" {

		raw_value, ok := (*data)[property]

		if ok == true {

			switch value := raw_value.(type) {

			case []bool:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatBool(value[v])

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]bool"
				result_value = formatted

			case bool:

				result_type  = "bool"
				result_value = strconv.FormatBool(value)

			case []byte:

				formatted := ""

				for v := 0; v < len(value); v++ {

					hex := strconv.FormatUint(uint64(value[v]), 16)

					if len(hex) == 1 {
						formatted += "0x0" + hex
					} else {
						formatted += "0x" + hex
					}

					if v < len(value) - 1 {
						formatted += " "
					}

				}

				result_type  = "[]byte"
				result_value = formatted

			case []float32:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatFloat(float64(value[v]), 'g', -1, 32)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]float32"
				result_value = formatted

			case float32:

				result_type  = "float32"
				result_value = strconv.FormatFloat(float64(value), 'g', -1, 32)

			case []float64:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatFloat(float64(value[v]), 'g', -1, 64)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]float64"
				result_value = formatted

			case float64:

				result_type  = "float64"
				result_value = strconv.FormatFloat(float64(value), 'g', -1, 64)

			case []int:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatInt(int64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]int"
				result_value = formatted

			case int:

				result_type  = "int"
				result_value = strconv.FormatInt(int64(value), 10)

			case []int8:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatInt(int64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]int8"
				result_value = formatted

			case int8:

				result_type  = "int8"
				result_value = strconv.FormatInt(int64(value), 10)

			case []int16:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatInt(int64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]int16"
				result_value = formatted

			case int16:

				result_type  = "int16"
				result_value = strconv.FormatInt(int64(value), 10)

			case []int32:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatInt(int64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]int32"
				result_value = formatted

			case int32:

				result_type  = "int32"
				result_value = strconv.FormatInt(int64(value), 10)

			case []int64:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatInt(int64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]int64"
				result_value = formatted

			case int64:

				result_type  = "int64"
				result_value = strconv.FormatInt(value, 10)

			case []string:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += value[v]

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]string"
				result_value = formatted

			case string:

				result_type  = "string"
				result_value = value

			case []uint:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatUint(uint64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]uint"
				result_value = formatted

			case uint:

				result_type  = "uint"
				result_value = strconv.FormatUint(uint64(value), 10)

			// case []uint8:
			// XXX: Same as []byte

			case uint8:

				result_type  = "uint8"
				result_value = strconv.FormatUint(uint64(value), 10)

			case []uint16:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatUint(uint64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]uint16"
				result_value = formatted

			case uint16:

				result_type  = "uint16"
				result_value = strconv.FormatUint(uint64(value), 10)

			case []uint32:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatUint(uint64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]uint32"
				result_value = formatted

			case uint32:

				result_type  = "uint32"
				result_value = strconv.FormatUint(uint64(value), 10)

			case []uint64:

				formatted := ""

				for v := 0; v < len(value); v++ {

					formatted += strconv.FormatUint(uint64(value[v]), 10)

					if v < len(value) - 1 {
						formatted += ","
					}

				}

				result_type  = "[]uint64"
				result_value = formatted

			case uint64:

				result_type  = "uint64"
				result_value = strconv.FormatUint(value, 10)

			}

		}

	}

	return result_type, result_value

}

