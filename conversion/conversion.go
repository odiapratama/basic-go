package main

import "fmt"

func main() {
	const value16 int16 = 127
	const value8 int8 = int8(value16)

	fmt.Println(value8)
	fmt.Println(value16)

	const name = "Odia"
	var firstChar = string(name[0])

	fmt.Println(name)
	fmt.Println(firstChar)
}
