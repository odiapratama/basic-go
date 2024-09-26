package main

import (
	"errors"
	"fmt"
)

// Division Number
func Division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider must not 0")
	} else {
		return value / divider, nil
	}
}

func main() {
	result1, err1 := Division(100, 10)
	result2, err2 := Division(100, 0)

	if err1 == nil {
		fmt.Println(result1)
	} else {
		fmt.Println(err1.Error())
	}

	if err2 == nil {
		fmt.Println(result2)
	} else {
		fmt.Println(err2.Error())
	}
}
