package data

import "bytes"
import "sort"

type Dataset []*Data

func NewDataset(length int) Dataset {
	return Dataset(make([]*Data, length))
}

func ToDataset(entries []Data) Dataset {

	var dataset Dataset = Dataset(make([]*Data, len(entries)))

	for e := 0; e < len(entries); e++ {
		entry := entries[e]
		dataset[e] = &entry
	}

	return dataset

}

func (dataset *Dataset) Add(data Data) bool {
	*dataset = append(*dataset, &data)
	return true
}

func (dataset *Dataset) Get(index int) *Data {

	var result *Data

	if dataset != nil && index >= 0 && index < len(*dataset) {
		result = (*dataset)[index]
	}

	return result

}

func (dataset *Dataset) Has(index int) bool {

	if dataset != nil && index >= 0 && index < len(*dataset) {
		return true
	}

	return false

}

func (dataset *Dataset) HasProperty(index int, property string) bool {

	if dataset != nil && index >= 0 && index < len(*dataset) {

		data := (*dataset)[index]
		_, ok2 := (*data)[property]

		if ok2 == true {
			return true
		}

	}

	return false

}

func (dataset *Dataset) Length() int {
	return len(*dataset)
}

func (dataset *Dataset) Set(index int, data Data) bool {

	if dataset != nil && index >= 0 && index < len(*dataset) {
		(*dataset)[index] = &data
		return true
	}

	return false

}

func (dataset *Dataset) SortByProperty(property string) []int {

	result := make([]int, len(*dataset))

	for d, _ := range *dataset {
		result[d] = d
	}

	sort.Slice(result, func(a int, b int) bool {

		value_a, ok_a := (*(*dataset)[result[a]])[property]
		value_b, ok_b := (*(*dataset)[result[b]])[property]

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
