package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :", args)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(name)
	} else {
		fmt.Println("Hostname not found!")
	}

	username := os.Getenv("EXPLORE_GO_USERNAME")
	password := os.Getenv("EXPLORE_GO_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
