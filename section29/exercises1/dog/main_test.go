package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(2)

	if v != 14 {
		t.Error("AAA", 14, "BBB", v)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

func Example_years() {
	v := Years(3)
	fmt.Println(v)
	// Output:
	// 21
}
