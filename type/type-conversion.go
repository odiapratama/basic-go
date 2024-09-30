package main

import "fmt"

func typeConversion() {
	type ID string
	const id ID = "234234234234"
	fmt.Println(id)
}
