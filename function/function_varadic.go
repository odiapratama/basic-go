package main

import "fmt"

func varadic() {
	println("=== Function Varadic ===")
	data := []int{1, 2, 3, 4, 5}
	totalSum := sum(data...)
	totalMultplication := multiplication(data)
	fmt.Println(totalSum)
	fmt.Println(totalMultplication)
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func multiplication(numbers []int) int {
	result := 1
	for _, number := range numbers {
		result *= number
	}
	return result
}
