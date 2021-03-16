package main

import "fmt"

func main() {
	var months = [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months)
	fmt.Println(len(months))

	var slice1 = months[4:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var array1 = append(slice1, "Test")
	fmt.Println(array1)

	slice2 := make([]string, len(slice1), cap(slice1))
	copy(slice2, slice1)
	fmt.Println(slice2)

	isArray := [...]int{0, 1, 2}
	isSlice := []int{0, 1, 2}
	fmt.Println(isArray)
	fmt.Println(isSlice)
}
