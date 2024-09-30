package main

import "fmt"

func recursive() {
	number := 10
	fmt.Println(factorialLoop(number))
	fmt.Println(factorialRecursive(number))
}

func factorialLoop(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorialRecursive(number-1)
}
