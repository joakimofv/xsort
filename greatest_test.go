package xsort

import (
	"testing"
)

func TestGreatest(t *testing.T) {
	ff := []float64{0.1, 0.2, 0.1}
	max := Greatest(ff, func(i, j int) bool { return ff[i] > ff[j] })
	if max != 1 {
		t.Errorf("Expected %v, got %v.", 1, max)
	}
}
