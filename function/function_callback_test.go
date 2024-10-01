package main

import "testing"

func TestAdd(t *testing.T) {
	a, b := 10, 5
	total := add(a, b)
	if total != 15 {
		t.Errorf("Add was incorrect, got: %d, want: %d.", total, 15)
	}
}

func TestSubtract(t *testing.T) {
	a, b := 10, 5
	total := subtract(a, b)
	if total != 5 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestMultiply(t *testing.T) {
	a, b := 10, 5
	total := multiply(a, b)
	if total != 50 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 50)
	}
}

func TestDivide(t *testing.T) {
	a, b := 10, 5
	total := divide(a, b)
	if total != 2 {
		t.Errorf("Divide was incorrect, got: %d, want: %d.", total, 2)
	}
}

func TestDoMath(t *testing.T) {
	a, b := 10, 5
	addition := doMath(a, b, add)
	subtraction := doMath(a, b, subtract)
	multiplication := doMath(a, b, multiply)
	division := doMath(a, b, divide)

	if addition != 15 {
		t.Errorf("Addition was incorrect, got: %d, want: %d.", addition, 15)
	}
	if subtraction != 5 {
		t.Errorf("Subtraction was incorrect, got: %d, want: %d.", subtraction, 5)
	}
	if multiplication != 50 {
		t.Errorf("Multiplication was incorrect, got: %d, want: %d.", multiplication, 50)
	}
	if division != 2 {
		t.Errorf("Division was incorrect, got: %d, want: %d.", division, 2)
	}
}
