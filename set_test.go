package rongqi

import "testing"

func TestIntSet_Add(t *testing.T) {
	set := NewIntSet()
	set.Add(1)

	if set.Len() != 1 {
		t.Errorf("set length should be 1 not %d. \n", set.Len())
	}
	if !set.Contains(1) {
		t.Error("set not contains 1")
	}
}
