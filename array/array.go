package main

import "fmt"

func main() {
	var titles = [4]string{
		"Go",
		"Playground",
		"Mantap",
	}

	fmt.Println(titles)
	fmt.Println(len(titles))

	titles[3] = "Asik"
	fmt.Println(len(titles))
}
