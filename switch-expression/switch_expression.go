package main

import "fmt"

func main() {
	id := "123123468598"

	switch id {
	case "123123468598":
		fmt.Println("Halo!")
	default:
		fmt.Println("Hai")
	}

	switch length := len(id); length {
	case 12:
		fmt.Println("Valid!")
	default:
		fmt.Println("Invalid!")
	}
}
