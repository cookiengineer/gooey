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
