package xsort

import (
	"reflect"
)

// Insert puts elem into the correct place of the sorted list x. The element that was dropped
// from the end of the list gets returned, which might be elem itself.
// less is a function that returns true if elem should be in a lower index than the value at index i,
// is similar to the less function in sort.Slice.
func Insert(x interface{}, elem interface{}, less func(i int) bool) (drop interface{}) {
	vv := reflect.ValueOf(x)
	v := reflect.ValueOf(elem)
	for i := 0; i < vv.Len(); i++ {
		if !less(i) {
			continue
		}
		drop = vv.Index(vv.Len() - 1).Interface()
		reflect.Copy(vv.Slice(i+1, vv.Len()), vv.Slice(i, vv.Len()-1))
		vv.Index(i).Set(v)
		return
	}
	drop = elem
	return
}
