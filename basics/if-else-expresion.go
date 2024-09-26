package main

import "fmt"

func main() {
	id := "218947287809"

	if id == "218947287809" {
		fmt.Println("Good morning!")
	} else if id == "123850123845" {
		fmt.Println("Halo!")
	} else {
		fmt.Println("ID not found")
	}

	if length := len(id); length == 12 {
		fmt.Println("ID Valid!")
	} else {
		fmt.Println("ID Invalid!")
	}
}
