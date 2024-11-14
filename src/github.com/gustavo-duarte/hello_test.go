package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Django")
	target := "Hello, Django"

	if result != target {
		t.Errorf("Result: '%s', Target: '%s'", result, target)
	}
}
