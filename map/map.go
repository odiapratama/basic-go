package main

import "fmt"

func main() {
	data := map[string]string{
		"id":   "12312312",
		"name": "Odia",
	}
	fmt.Println(data)
	fmt.Println(len(data))

	books := make(map[string]string)
	books["title"] = "Basic Go"
	books["autor"] = "Odia"
	fmt.Println(books)
	fmt.Println(len(books))

	books["desc"] = "Test"
	fmt.Println(books)
	fmt.Println(len(books))

	delete(books, "desc")
	fmt.Println(books)
	fmt.Println(len(books))
}
