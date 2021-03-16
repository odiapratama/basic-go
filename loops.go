package main

import "fmt"

func main() {
	index := 1
	for index <= 10 {
		fmt.Println(index)
		index++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	slice := []string{
		"Indonesia",
		"US",
		"Japan",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index ", i, "=", value)
	}

	dataMap := map[string]string{
		"name":  "Odia",
		"title": "Programmer",
	}

	for key, value := range dataMap {
		fmt.Println(key, "=", value)
	}
}
