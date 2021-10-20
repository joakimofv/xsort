package xsort

import (
	"reflect"
)

// Greatest gets the index of the greatest element, as given by the greater comparison function.
func Greatest(x interface{}, greater func(i, j int) bool) int {
	vv := reflect.ValueOf(x)
	if vv.Len() < 1 {
		panic("empty slice")
	}
	greatestIdx := 0
	for i := 1; i < vv.Len(); i++ {
		if greater(i, greatestIdx) {
			greatestIdx = i
		}
	}
	return greatestIdx
}
