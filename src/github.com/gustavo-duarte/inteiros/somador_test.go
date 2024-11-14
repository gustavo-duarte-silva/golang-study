package inteiros

import (
	"fmt"
	"testing"
)

func TestSomador(t *testing.T) {
	results := Soma(2, 2)
	target := 4
	if results != target {
		t.Errorf("Results: '%d' | Target: '%d'", results, target)
	}
}

func ExampleSoma() {
	results := Soma(1, 5)
	fmt.Println(results)
	// Output: 6
}
