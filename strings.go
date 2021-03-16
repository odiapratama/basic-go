package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello world"
	fmt.Println(strings.Contains(text, "world"))
	fmt.Println(strings.Count(text, "world"))
	fmt.Println(strings.Replace(text, "world", "you", 1))
	fmt.Println(strings.Split(text, " "))
}
