package main

import "fmt"

func main() {
	const s1 = 8
	const s2 = 9
	var total = s1 * s2

	fmt.Println(total)

	total++
	fmt.Println(total)
}
