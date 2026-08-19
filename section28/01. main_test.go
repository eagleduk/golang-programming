package main

import "testing"

func TestSection28_sum(t *testing.T) {
	x := section28_sum(1, 2, 3, 4, 5, 6)

	if x != 21 {
		t.Error("incorrect !")
	}
}
