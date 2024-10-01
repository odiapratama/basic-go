package main

import "fmt"

func closure() {
	fmt.Println("=== Function Closure ===")
	total := 0

	increment := func() {
		total++
		fmt.Println(total)
	}

	increment()
	increment()
	fmt.Println(total)
}
