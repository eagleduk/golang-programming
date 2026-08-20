package word

import (
	"fmt"
	"golang-programming/section29/exercises2/quote"
	"testing"
)

func Example_count() {
	fmt.Println(Count("One Two Three"))
	// Output:
	// 3
}

func Benchmark_useCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func Benchmark_count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}

}

func TestCount(t *testing.T) {
	v := Count("One Two Three Tree")

	if v != 4 {
		t.Error("Got", v, "want", 4)
	}
}
