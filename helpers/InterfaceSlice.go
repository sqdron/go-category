package helpers;

import "reflect"

func InterfaceSlice(slice interface{}) [] interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return []interface{}{slice}
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}