package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(10)

	for i := 0; i < data.Len(); i++ {
		data.Value = i
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
