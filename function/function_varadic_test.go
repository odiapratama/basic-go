package main

import "testing"

func TestSum(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	total := sum(data...)
	if total != 15 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 15)
	}
}

func TestMultiplication(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	total := multiplication(data)
	if total != 120 {
		t.Errorf("Multiplication was incorrect, got: %d, want: %d.", total, 120)
	}
}
