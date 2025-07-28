package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	result := Iteration("a", 5)
	got := "aaaaa"
	if result != got {
		t.Errorf("result is '%s' but got '%s'", result, got)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a", 5)
	}
}

func ExampleIteration() {
	s := Iteration("b", 5)
	fmt.Println(s)
	// Output: bbbbb
}
