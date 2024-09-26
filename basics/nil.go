package main

import "fmt"

// NewMap function
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("James")
	fmt.Println(person)

	var employee map[string]string = nil
	fmt.Println(employee)
}
