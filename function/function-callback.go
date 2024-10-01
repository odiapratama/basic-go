package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func callback() {
	fmt.Println("=== Function Callback ===")
	a, b := 10, 5
	addition := doMath(a, b, add)
	subtraction := doMath(a, b, subtract)
	multiplication := doMath(a, b, multiply)
	division := doMath(a, b, divide)

	println("Addition:", addition)
	println("Subtraction:", subtraction)
	println("Multiplication:", multiplication)
	println("Division:", division)
}
