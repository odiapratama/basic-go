package main

import "testing"

func TestFactorialLoop(t *testing.T) {
	number := 5
	result := factorialLoop(number)
	if result != 120 {
		t.Errorf("FactorialLoop was incorrect, got: %d, want: %d.", result, 120)
	}
}

func TestFactorialRecursive(t *testing.T) {
	number := 5
	result := factorialRecursive(number)
	if result != 120 {
		t.Errorf("FactorialRecursive was incorrect, got: %d, want: %d.", result, 120)
	}
}
