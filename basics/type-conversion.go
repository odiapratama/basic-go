package main

import "fmt"

func main() {
	type ID string
	const id ID = "234234234234"
	fmt.Println(id)
}
