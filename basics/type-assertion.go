package main

import (
	"fmt"
)

func getData() interface{} {
	return "Haloo"
}

func main() {
	switch dataType := getData().(type) {
	case string:
		fmt.Println("String :", dataType)
	case int:
		fmt.Println("Integer :", dataType)
	default:
		fmt.Println("Unknown data")
	}
}
