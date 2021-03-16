package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "Put username")
	password := flag.String("password", "root", "Put password")

	flag.Parse()

	fmt.Println(*username)
	fmt.Println(*password)
}
