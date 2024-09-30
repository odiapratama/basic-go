package main

import "fmt"

func greeting(greetingType int) interface{} {
	switch greetingType {
	case 1:
		return "Hi"
	case 2:
		return 123
	case 3:
		return true
	default:
		return 0
	}
}

func interfaceEmpty() {
	fmt.Println(greeting(1))
	fmt.Println(greeting(2))
	fmt.Println(greeting(3))
	fmt.Println(greeting(99))
}
