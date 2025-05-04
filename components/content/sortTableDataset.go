package content

import "bytes"
import "sort"

func sortTableDataset(dataset []TableData, property string) []int {

	result := make([]int, len(dataset))

	for d := 0; d < len(dataset); d++ {
		result[d] = d
	}

	sort.Slice(result, func(a int, b int) bool {

		value_a, ok_a := dataset[result[a]][property]
		value_b, ok_b := dataset[result[b]][property]

		if ok_a == true && ok_b == true {

			switch value_a.(type) {

			case []byte:

				tmp_a := value_a.([]byte)
				tmp_b := value_b.([]byte)

				return bytes.Compare(tmp_a, tmp_b) < 0

			case bool:

				tmp_a := value_a.(bool)
				tmp_b := value_b.(bool)

				if tmp_a == true && tmp_b == false {
					return true
				} else {
					return false
				}

			case float32:

				tmp_a := value_a.(float32)
				tmp_b := value_b.(float32)

				return tmp_a < tmp_b

			case float64:

				tmp_a := value_a.(float64)
				tmp_b := value_b.(float64)

				return tmp_a < tmp_b

			case int:

				tmp_a := value_a.(int)
				tmp_b := value_b.(int)

				return tmp_a < tmp_b

			case int8:

				tmp_a := value_a.(int8)
				tmp_b := value_b.(int8)

				return tmp_a < tmp_b

			case int16:

				tmp_a := value_a.(int16)
				tmp_b := value_b.(int16)

				return tmp_a < tmp_b

			case int32:

				tmp_a := value_a.(int32)
				tmp_b := value_b.(int32)

				return tmp_a < tmp_b

			case int64:

				tmp_a := value_a.(int64)
				tmp_b := value_b.(int64)

				return tmp_a < tmp_b

			case string:

				tmp_a := value_a.(string)
				tmp_b := value_b.(string)

				return tmp_a < tmp_b

			case uint:

				tmp_a := value_a.(uint)
				tmp_b := value_b.(uint)

				return tmp_a < tmp_b

			case uint8:

				tmp_a := value_a.(uint8)
				tmp_b := value_b.(uint8)

				return tmp_a < tmp_b

			case uint16:

				tmp_a := value_a.(uint16)
				tmp_b := value_b.(uint16)

				return tmp_a < tmp_b

			case uint32:

				tmp_a := value_a.(uint32)
				tmp_b := value_b.(uint32)

				return tmp_a < tmp_b

			case uint64:

				tmp_a := value_a.(uint64)
				tmp_b := value_b.(uint64)

				return tmp_a < tmp_b

			}

		}

		return false

	})

	return result

}
