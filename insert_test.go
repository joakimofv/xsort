package xsort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	slice := make([]float64, 5)
	// Less
	for n := 0; n < 10; n++ {
		drop := Insert(slice, float64(n), func(i int) bool { return float64(n) < slice[i] })
		if drop.(float64) != float64(n) {
			t.Error(drop)
		}
		if !sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] < slice[j] }) {
			t.Error(n, slice)
		}
	}
	for i := range slice {
		slice[i] = 10
	}
	for n := 0; n < 10; n++ {
		drop := Insert(slice, float64(n), func(i int) bool { return float64(n) < slice[i] })
		expected := float64(10)
		if n >= 5 {
			expected = float64(n)
		}
		if drop.(float64) != expected {
			t.Errorf("[%v] Expected drop %v, got %v.\n%v", n, expected, drop, slice)
		}
		if !sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] < slice[j] }) {
			t.Error(n, slice)
		}
	}
	// More
	slice = make([]float64, 5)
	for n := 0; n < 10; n++ {
		drop := Insert(slice, float64(n), func(i int) bool { return float64(n) > slice[i] })
		expected := float64(0)
		if n >= 5 {
			expected = float64(n) - 5
		}
		if drop.(float64) != expected {
			t.Errorf("[%v] Expected drop %v, got %v.\n%v", n, expected, drop, slice)
		}
		if !sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] > slice[j] }) {
			t.Error(n, slice)
		}
	}
	// Random
	slice2 := make([]uint64, 10)
	for i := range slice2 {
		slice2[i] = rand.Uint64()
	}
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })
	for n := 0; n < 10000; n++ {
		elem := rand.Uint64()
		drop := Insert(slice2, elem, func(i int) bool { return elem < slice2[i] })
		for i := range slice2 {
			if drop.(uint64) < slice2[i] {
				t.Error(slice2, drop)
			}
		}
		if !sort.SliceIsSorted(slice2, func(i, j int) bool { return slice2[i] < slice2[j] }) {
			t.Error(slice2)
		}
	}
}
