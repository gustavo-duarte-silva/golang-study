package loop

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	results := Loops("A", 10)
	target := "AAAAAAAAAA"

	if results != target {
		t.Errorf("Results: '%s' | Target: '%s'", results, target)
	}

}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loops("A", 10)
	}
}

func ExampleLoops() {
	results := Loops("A", 10)
	fmt.Println(results)
	// Output: AAAAAAAAAA
}
