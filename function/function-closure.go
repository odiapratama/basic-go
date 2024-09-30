package main

import "fmt"

func closure() {
	total := 0

	increment := func() {
		total++
		fmt.Println(total)
	}

	increment()
	increment()
	fmt.Println(total)
}
