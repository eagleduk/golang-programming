package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	v := CenteredAvg([]int{
		1, 2, 3, 4, 5,
	})

	if v != 3 {
		t.Error("Got ", v, " want ", 3)
	}
}

func ExampleCenteredAvg() {
	v := CenteredAvg([]int{
		1, 1, 1, 1, 1, 1, 1,
	})
	fmt.Println(v)
	// Output:
	// 1
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{
			1, 2, 3, 4, 5,
		})
	}
}
