package reactive

import "reflect"

func is_comparable(item any) bool {

	typ := reflect.TypeOf(item)

	if typ == nil {
		return false
	}

	return typ.Comparable()

}
