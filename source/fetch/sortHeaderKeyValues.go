//go:build wasm

package fetch

import "sort"

type headerKeyValues struct {
	key    string
	values []string
}

type headerSorter struct {
	keyvalues []headerKeyValues
}

func (sorter headerSorter) Len() int {
	return len(sorter.keyvalues)
}

func (sorter headerSorter) Swap(i, j int) {
	sorter.keyvalues[i], sorter.keyvalues[j] = sorter.keyvalues[j], sorter.keyvalues[i]
}

func (sorter headerSorter) Less(i, j int) bool {
	return sorter.keyvalues[i].key < sorter.keyvalues[j].key
}

func sortHeaderKeyValues(headers Headers) []headerKeyValues {

	sorter := headerSorter{
		keyvalues: make([]headerKeyValues, 0),
	}

	for key, values := range headers {
		sorter.keyvalues = append(sorter.keyvalues, headerKeyValues{key, values})
	}

	sort.Sort(sorter)

	return sorter.keyvalues

}
