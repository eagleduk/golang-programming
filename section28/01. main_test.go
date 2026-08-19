package main

import (
	"fmt"
	"testing"
)

func TestSection28_sum(t *testing.T) {
	x := section28_sum(1, 2, 3, 4, 5, 6)

	if x != 21 {
		t.Error("incorrect !")
	}
}

func TestSection28_sum2(t *testing.T) {

	type test struct {
		data   []int
		anwser int
	}

	tests := []test{
		test{
			data:   []int{1, 2, 3, 4, 5},
			anwser: 15,
		},
		test{
			data:   []int{1, 0, -1},
			anwser: 0,
		},
	}

	for _, v := range tests {
		x := section28_sum(v.data...)

		if x != v.anwser {
			t.Error("incorrect !")
		}
	}

}

func TestSection28_sum3(t *testing.T) {
	x := section28_sum(1, 2, 3)
	fmt.Println(x)
	// Output:
	// 6
}
