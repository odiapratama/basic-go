package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Haloo")
	data.PushBack(1)
	data.PushFront(true)
	data.PushBack(2.0)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
